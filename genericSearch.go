package main

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (cfg *apiConfig) handlerGenericSearch(w http.ResponseWriter, r *http.Request) {
	// PARSE PARAMETERS
	type request struct {
		QueryParam string // "Instructor" | "Building" | "Meeting Days" | "Faculty Ratio"
		ParamValue string
	}
	var req request

	req.QueryParam = r.URL.Query().Get("queryParam")
	req.ParamValue = r.URL.Query().Get("paramValue")

	// QUERY COLLECTION

	var filter bson.D

	switch req.QueryParam {
	case "Instructor":
		filter = bson.D{{Key: "Instructor", Value: req.ParamValue}}
	case "Building":
		filter = bson.D{{Key: "Building", Value: req.ParamValue}}
	case "Meeting Days":
		filter = bson.D{{Key: "Meeting Days", Value: req.ParamValue}}
	case "Faculty Ratio":
		filter = bson.D{{Key: "Faculty Ratio", Value: req.ParamValue}}
	}

	cursor, err := cfg.coll.Find(context.TODO(), filter)
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

	// RETURN COURSE DATA

	err = respondWithJSON(w, 200, result)
	if err != nil {
		fmt.Println("Failed to respond")
	}
}
