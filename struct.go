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
	Term              string   `bson:"term" json:"term"`
	Level             string   `bson:"level" json:"level"`
	Section           string   `bson:"section" json:"section"`
	CRN               string   `bson:"crn" json:"crn"`
	ShortName         string   `bson:"shortName" json:"shortName"`
	LongName          string   `bson:"longName" json:"longName"`
	Enrollment        int      `bson:"enrollment" json:"enrollment"`
	TotalTAs          int      `bson:"totalTAs" json:"totalTAs"`
	MeetingRoom       string   `bson:"meetingRoom" json:"meetingRoom"`
	MeetingDays       []string `bson:"meetingDays" json:"meetingDays"`
	MeetingTimes      []string `bson:"meetingTimes" json:"meetingTimes"`
	InstructorName    string   `bson:"instructorName" json:"instructorName"`
	InstructorEmail   string   `bson:"instructorEmail" json:"instructorEmail"`
	CourseDescription string   `bson:"courseDescription" json:"courseDescription"`
}
