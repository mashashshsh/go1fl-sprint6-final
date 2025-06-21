package service

import (
	"errors"
	"strings"

	morse "go1fl-sprint6-final/pkg/morse"
)

func Convert(someText string) (string, error) {
	trimmedText := strings.TrimSpace(someText)

	if isMorse(trimmedText) {
		result := morse.ToText(trimmedText)

		if result == "" {
			return "", errors.New("Error durig converting to text")
		}

		return result, nil
	} else {
		result := morse.ToMorse(trimmedText)

		if result == "" {
			return "", errors.New("Error durig converting to Morse")
		}

		return result, nil
	}
}

func isMorse(someText string) bool {
	for _, ch := range someText {
		if ch != '.' && ch != '-' && ch != ' ' {
			return false
		}
	}

	return true
}
