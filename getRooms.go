package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (cfg *apiConfig) handlerGetRooms(w http.ResponseWriter, r *http.Request) {
	// QUERY FOR ALL ROOMS AND UNAVAILABLE ROOMS ==> COMPUTE AVAILABLE ROOMS
	type request struct {
		time string
		days string
	}
	var req request

	req.time = r.URL.Query().Get("timeToCheck")
	req.days = r.URL.Query().Get("daysToCheck")

	// DISTINCT QUERY ALL ROOMS

	var allRooms []string
	err := cfg.coll.Distinct(context.TODO(), "meetingRoom", bson.D{{}}).Decode(&allRooms)
	if err != nil {
		respondWithError(w, 400, "Failed to query collection")
		return
	}

	// DISTINCT QUERY UNAVAILABLE ROOMS

	days := strings.Split(req.days, "/")

	// ASSUME TIME VALUE IS meetingTimes.0
	filter := bson.D{{Key: "$and", Value: bson.A{
		bson.D{{Key: "meetingTimes.0", Value: req.time}},
		bson.D{{Key: "meetingDays", Value: bson.D{{
			Key: "$all", Value: days,
		}}}},
	}}}

	var unavailableRooms []string
	err = cfg.coll.Distinct(context.TODO(), "meetingRoom", filter).Decode(&unavailableRooms)
	if err != nil {
		respondWithError(w, 400, "Failed to query collection")
		return
	}

	// COMPUTE AVAILABLE ROOMS

	var availableRooms []string
	for _, all := range allRooms {
		available := true
		for _, un := range unavailableRooms {
			if all == un {
				available = false
				break
			}
		}
		if available {
			availableRooms = append(availableRooms, all)
		}
	}

	// RESPOND WITH AVAILABLE ROOMS

	err = respondWithJSON(w, 200, availableRooms)
	if err != nil {
		fmt.Println("Failed to respond")
	}
}
