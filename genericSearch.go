package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
	var cursor *mongo.Cursor
	var err error
	var sort int = 0

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
		if req.ParamValue == "Most Support" {
			sort = 1
		} else {
			sort = -1
		}
		filter = bson.D{{Key: "enrollment", Value: bson.D{{Key: "$ne", Value: 0}}}}
	}

	if sort == 0 {
		cursor, err = cfg.coll.Find(context.TODO(), filter)
		if err != nil {
			respondWithError(w, 500, fmt.Sprintf("Failed to query collection: %s", err))
			return
		}
	} else {
		cursor, err = cfg.coll.Aggregate(context.TODO(), mongo.Pipeline{
			// filter
			{{
				Key: "$match", Value: filter,
			}},
			// expression
			{{
				Key: "$addFields", Value: bson.D{{
					Key: "ratio", Value: bson.D{{
						Key: "$divide", Value: bson.A{
							"$enrollment", bson.D{{
								Key: "$add", Value: bson.A{
									"$totalTAs", 1,
								}}},
						},
					}},
				}},
			}},
			// sort
			{{
				Key: "$sort", Value: bson.D{{
					Key: "ratio", Value: sort,
				}},
			}},
		})
		if err != nil {
			respondWithError(w, 500, fmt.Sprintf("Failed to query collection: %s", err))
			return
		}
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
