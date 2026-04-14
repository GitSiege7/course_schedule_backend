package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerCreateCourse(w http.ResponseWriter, r *http.Request) {
	// PARSE PARAMETERS
	type request struct {
		newCourse course
	}
	var req request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, 400, "Failed to decode request")
		return
	}

}
