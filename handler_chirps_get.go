package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {

	chirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps")
	}

	respondWithJSON(w, http.StatusCreated, chirps)
}
