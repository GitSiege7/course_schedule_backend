package main

import "go.mongodb.org/mongo-driver/v2/mongo"

type apiConfig struct {
	uri       string
	db_name   string
	coll_name string
	client    *mongo.Client
}

type request struct {
	QueryParam string `json:"queryParam"`
	ParamValue string `json:"paramValue"`
}
