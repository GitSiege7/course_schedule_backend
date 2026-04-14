package main

import "go.mongodb.org/mongo-driver/v2/mongo"

type apiConfig struct {
	uri       string
	db_name   string
	coll_name string
	client    *mongo.Client
	coll      *mongo.Collection
}

type course struct {
	term              string
	level             string
	section           string
	crn               string
	shortName         string // ex. "COP 4930"
	longName          string // ex. "Programming Concepts"
	enrollment        int    // # of students in class
	totalTAs          int    // # of UG and GR TAs
	meetingRoom       string
	meetingDays       []string // slice of single letter day indicators (M, T, W, R, F)
	meetingTimes      []string // two element slice: start/end time
	instructorName    string
	instructorEmail   string
	courseDescription string
}
