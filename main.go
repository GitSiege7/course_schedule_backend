package main

import (
	"net/http"
)

// PORT: 95887

func main() {

	// CREATE CFG TO HOLD API STATE
	cfg := &apiConfig{
		uri:       "mongodb+srv://root:WvgrxVMB8aRcOMUe@course-schedule.gdpssbf.mongodb.net/?appName=course-schedule",
		db_name:   "course_scheduler_db",
		coll_name: "bellini_classes",
	}

	// DB CONNECT
	cfg.dbConnect()

	// DEFER DB DISCONNECT
	defer cfg.dbDisconnect()

	// CREATE HTTP SERVICE
	mux := http.NewServeMux()
	server := http.Server{
		Handler: mux,
		Addr:    ":95887",
	}

	// SET HANDLER FUNCTIONS
	mux.HandleFunc("GET /api/courses", cfg.handlerQuery)

	server.ListenAndServe()
}
