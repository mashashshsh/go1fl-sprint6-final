package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Error for reading file from 'index.html'", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(data))
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Error for ParseMultipartForm", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error for getting file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error for reading ulpoaded file", http.StatusInternalServerError)
		return
	}

	inputString := string(data)

	result, err := service.Convert(inputString)
	if err != nil {
		http.Error(w, "Error for during convert file text", http.StatusInternalServerError)
		return
	}

	filename := "result_" + time.Now().UTC().Format("20060102_150405") + filepath.Ext(header.Filename)

	if err := os.WriteFile(filename, []byte(result), 0755); err != nil {
		http.Error(w, "Error for writing file result", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(result))
}
