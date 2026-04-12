package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (cfg *apiConfig) dbConnect() {
	// CONNECT TO DB
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.uri))
	if err != nil {
		log.Fatal(err)
	}

	cfg.client = client
}

func (cfg *apiConfig) dbDisconnect() {
	if err := cfg.client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
