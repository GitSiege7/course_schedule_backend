package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (cfg *apiConfig) handlerGenericSearch(w http.ResponseWriter, r *http.Request) {
	// PARSE PARAMETERS
	type request struct {
		QueryParam string `json:"queryParam"`
		ParamValue string `json:"paramValue"`
	}
	var req request

	req.QueryParam = r.URL.Query().Get("queryParam")
	req.ParamValue = r.URL.Query().Get("paramValue")

	// QUERY COLLECTION
	var filter bson.D

	switch req.QueryParam {
	case "Instructor":
		filter = bson.D{{Key: "instructorName", Value: req.ParamValue}}
	case "Building":
		filter = bson.D{{Key: "meetingRoom", Value: req.ParamValue}}
	case "Meeting Days":
		days := strings.Split(req.ParamValue, "/")
		filter = bson.D{{Key: "meetingDays", Value: bson.D{{
			Key: "$all", Value: bson.A{days},
		}}}}
	case "Faculty Ratio":
		pValue, err := strconv.Atoi(req.ParamValue)
		if err != nil {
			respondWithError(w, 400, "Invalid parameter value input")
			return
		}
		filter = bson.D{{Key: "$expr", Value: bson.D{{
			Key: "$gte", Value: bson.A{
				bson.D{{
					Key: "$divide", Value: bson.A{
						"$enrollment", bson.D{{
							Key: "$add", Value: bson.A{
								"$totalTAs", 1,
							}}},
					},
				}},
				pValue,
			},
		}}}}
	}

	cursor, err := cfg.coll.Find(context.TODO(), filter)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Failed to query collection: %s", err))
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
