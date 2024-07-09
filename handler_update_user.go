package main

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	params, err := decodeParameters(r)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	foundUser, err := cfg.getUserWithEmail(params.Email)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find user")
		return
	}

	jwt, err := cfg.createJWT(params.ExpiresInSeconds, foundUser.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
	}

	type responseParams struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
		Token string `json:"token"`
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(params.Password))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	} else {
		respondWithJSON(w, http.StatusOK, responseParams{
			ID:    foundUser.ID,
			Email: foundUser.Email,
			Token: jwt})
		return
	}
}
