package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func (apiCfg *apiConfig) handlerPost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		// any missing fields will simply have their values in the struct set to their zero value
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	bodyLen := len(params.Body)

	if bodyLen > 140 {
		respondWithError(w, 400, "Chirp is too long")
		return
	}
	words := strings.Split(params.Body, " ")
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

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}
	respBody := returnVals{
		CleanedBody: newBody,
	}
	respondWithJSON(w, http.StatusOK, respBody)
	return
}
