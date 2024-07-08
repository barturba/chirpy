package main

import (
	"errors"
	"strings"
)

func validateChirp(s string) (string, error) {
	bodyLen := len(s)

	if bodyLen > 140 {
		return "", errors.New("Chirp is too long")
	}
	words := strings.Split(s, " ")
	var newWords []string
	for _, word := range words {
		if strings.ToLower(word) == "kerfuffle" {
			word = "****"
		} else if strings.ToLower(word) == "sharbert" {
			word = "****"
		} else if strings.ToLower(word) == "fornax" {
			word = "****"
		}
		newWords = append(newWords, word)
	}
	newBody := strings.Join(newWords, " ")
	return newBody, nil
}
