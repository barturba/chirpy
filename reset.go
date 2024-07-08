package main

import "net/http"

func (cfg *apiConfig)handlerReset(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	cfg.fileserverHits = 0
}