package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (cfg *apiConfig) handlerUpdateCourseDescription(w http.ResponseWriter, r *http.Request) {
	// PARSE PARAMETERS
	type request struct {
		CRN  string `json:"crnToModify"`
		Desc string `json:"newDescription"`
	}
	var req request

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		respondWithError(w, 400, "Failed to decode request")
		return
	}

	// UPDATE COLLECTION
	_, err = cfg.coll.UpdateOne(context.TODO(),
		bson.D{{Key: "crn", Value: req.CRN}},
		bson.D{{Key: "$set", Value: bson.D{{
			Key: "courseDescription", Value: req.Desc,
		}}}})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed update query: %s", err))
		return
	}

	// RETURN DATA
	err = respondWithJSON(w, 200, map[string]string{"msg": "Update Successful"})
	if err != nil {
		fmt.Println("Failed to respond")
	}
}
