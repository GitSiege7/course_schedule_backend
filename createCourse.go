package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerCreateCourse(w http.ResponseWriter, r *http.Request) {
	// PARSE PARAMETERS
	type request struct {
		NewCourse course `json:"newCourse"`
	}
	var req request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, 400, "Failed to decode request")
		return
	}

	// INSERT QUERY
	new, err := cfg.coll.InsertOne(context.TODO(), req.NewCourse)
	if err != nil {
		respondWithError(w, 500, "Failed to insert document")
		return
	}

	fmt.Println("createdID: ", new.InsertedID)

	// RESPOND WITH SUCCESS
	err = respondWithJSON(w, 200, map[string]string{"msg": "Document Successfully Created"})
	if err != nil {
		fmt.Println("Failed to respond")
	}
}
