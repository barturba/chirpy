package main

import "net/http"

func (cfg *apiConfig) handleUsers(w http.ResponseWriter, r *http.Request) {

	respondWithJSON(w, http.StatusOK, "OK")
}
