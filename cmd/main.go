package main

import (
	"log"
	"os"

	"go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "logs: ", log.LstdFlags|log.Lshortfile)

	srv := server.CreateNewServer(logger)

	logger.Println("Start server on port 8080")
	if err := srv.Start(); err != nil {
		logger.Fatal("Some error during starting server:", err)
	}
}
