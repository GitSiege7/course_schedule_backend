package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiConfig) handlerQuery(w http.ResponseWriter, r *http.Request) {
	// PARSE PARAMETERS
	var req request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, 400, "Failed to decode request")
		return
	}

	// QUERY COLLECTION
	coll := cfg.client.Database(cfg.db_name).Collection(cfg.coll_name)

	//...

	// RETURN JSON DATA

	//...
}
