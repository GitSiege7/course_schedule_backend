package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerDeleteCourse(w http.ResponseWriter, r *http.Request) {
	type request struct {
		crnToDelete string
	}
	var req request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, 400, "Failed to decode request")
		return
	}

}
