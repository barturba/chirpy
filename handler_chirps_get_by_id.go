package main

import (
	"fmt"
	"internal/database"
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerChirpsGetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("chirpID"))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid chirp id")
		return
	}

	fmt.Printf("id: %v\n", id)

	chirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get chirps")
		return
	}
	var found bool
	var foundChirp database.Chirp
	for _, chirp := range chirps {
		if chirp.ID == id {
			found = true
			foundChirp = chirp
		}
	}
	if found {
		respondWithJSON(w, http.StatusOK, foundChirp)
		return
	} else {
		respondWithError(w, http.StatusNotFound, "Couldn't find chirp")
		return
	}
}
