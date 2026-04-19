package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (cfg *apiConfig) handlerDeleteCourse(w http.ResponseWriter, r *http.Request) {
	type request struct {
		CRN string `json:"crnToDelete"`
	}
	var req request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, 400, "Failed to decode request")
		return
	}

	// DELETE QUERY
	_, err = cfg.coll.DeleteOne(context.TODO(), bson.D{{Key: "crn", Value: req.CRN}})
	if err != nil {
		respondWithError(w, 500, "Failed to delete document")
		return
	}

	// RESPOND WITH SUCCESS
	err = respondWithJSON(w, 200, map[string]string{"msg": "Document Successfully Deleted"})
	if err != nil {
		fmt.Println("Failed to respond")
	}
}
