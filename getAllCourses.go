package main

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (cfg *apiConfig) handlerGetAllCourses(w http.ResponseWriter, r *http.Request) {
	// QUERY COLLECTION
	cursor, err := cfg.coll.Find(context.TODO(), bson.D{})
	if err != nil {
		respondWithError(w, 500, "Failed to query collection")
		return
	}
	defer cursor.Close(context.TODO())

	// GET DATA W/ CURSOR
	var result []course
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		respondWithError(w, 500, "Failed to retrieve all documents")
		return
	}

	// RETURN COURSE DATA
	err = respondWithJSON(w, 200, result)
	if err != nil {
		fmt.Println("Failed to respond")
	}
}
