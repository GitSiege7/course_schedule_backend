package main

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (cfg *apiConfig) handlerGetProfessorSchedule(w http.ResponseWriter, r *http.Request) {
	// PARSE PARAMETERS
	type request struct {
		Email string
	}
	var req request

	req.Email = r.URL.Query().Get("professorEmail")

	// QUERY COLLECTION

	cursor, err := cfg.coll.Find(context.TODO(), bson.D{{Key: "instructorEmail", Value: req.Email}})
	if err != nil {
		respondWithError(w, 500, "Failed to query collection")
		return
	}
	defer cursor.Close(context.TODO())

	var result []course
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		respondWithError(w, 500, "Failed to retrieve all documents")
		return
	}

	// RETURN DATA

	err = respondWithJSON(w, 200, result)
	if err != nil {
		fmt.Println("Failed to respond")
	}

}
