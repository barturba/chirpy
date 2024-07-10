package main

import (
	"errors"
	"internal/database"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (cfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	params, err := decodeParameters(r)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	authorization := r.Header.Get("Authorization")
	authorization = strings.TrimPrefix(authorization, "Bearer ")
	type MyCustomClaims struct {
		Foo string `json:"foo"`
		jwt.RegisteredClaims
	}

	token, err := jwt.ParseWithClaims(authorization, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	userID := 0
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Bad token")
		return
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		userID, err = strconv.Atoi(claims.RegisteredClaims.Subject)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Internal server error")
			return

		}
	} else {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	_, err = cfg.getUserWithID(userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find user")
		return
	}

	updatedUser, err := cfg.DB.UpdateUser(userID, params.Email, params.Password)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
	}

	type responseParams struct {
		ID    int    `json:"id"`
		Email string `json:"email"`
	}

	respondWithJSON(w, http.StatusOK, responseParams{
		ID:    updatedUser.ID,
		Email: updatedUser.Email,
	})

}

func (cfg *apiConfig) getUserWithID(id int) (database.User, error) {
	users, err := cfg.DB.GetUsers()
	if err != nil {
		return database.User{}, err
	}
	var found bool
	var foundUser database.User
	for _, user := range users {
		if user.ID == id {
			found = true
			foundUser = user
		}
	}
	if !found {
		return database.User{}, errors.New("User not found")
	}
	return foundUser, nil
}
