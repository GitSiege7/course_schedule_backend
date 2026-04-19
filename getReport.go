package main

import (
	"context"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (cfg *apiConfig) handlerGetReport(w http.ResponseWriter, r *http.Request) {
	// QUERY FOR ANY CLASSES WITH INVALID DATA (NO ASSN ROOM "N/A", ASSN PROF "N/A")
	// MAYBE EXPAND TO CLASSES WITH INVALID DAYS (meetingDays not containing only M,T,W,R,F)
	filter := bson.D{{Key: "$or", Value: bson.A{
		bson.D{{Key: "meetingRoom", Value: "N/A"}},
		bson.D{{Key: "instructorName", Value: "N/A"}},
		bson.D{{Key: "instructorEmail", Value: "N/A"}},
	}}}

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
