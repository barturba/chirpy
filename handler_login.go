package main

import (
	"encoding/json"
	"errors"
	"internal/database"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type parameters struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	ExpiresInSeconds int    `json:"expires_in_seconds"`
}

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	params, err := decodeParameters(r)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}
	if params.ExpiresInSeconds == 0 {
		params.ExpiresInSeconds = 86400
	} else if params.ExpiresInSeconds > 86400 {
		params.ExpiresInSeconds = 86400
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

func decodeParameters(r *http.Request) (parameters, error) {
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		return parameters{}, err
	}
	return params, nil
}

func (cfg *apiConfig) getUserWithEmail(email string) (database.User, error) {
	users, err := cfg.DB.GetUsers()
	if err != nil {
		return database.User{}, err
	}
	var found bool
	var foundUser database.User
	for _, user := range users {
		if user.Email == email {
			found = true
			foundUser = user
		}
	}
	if !found {
		return database.User{}, errors.New("User not found")
	}
	return foundUser, nil
}
