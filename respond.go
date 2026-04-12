package main

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func respondWithJSON(w http.ResponseWriter, code int, dat interface{}) error {
	resp, err := json.Marshal(dat)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
	return nil
}

func respondWithBSON(w http.ResponseWriter, code int, dat interface{}) error {
	resp, err := bson.Marshal(dat)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/bson")
	w.WriteHeader(code)
	w.Write(resp)
	return nil
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}
