package main

import (
	"net/http"
)

// PORT: 9517

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
		Addr:    ":9517",
	}

	// SET HANDLER FUNCTIONS
	mux.HandleFunc("GET /api/courses", cfg.handlerGenericSearch)                 // UNTESTED
	mux.HandleFunc("GET /api/schedule", cfg.handlerGetProfessorSchedule)         // UNTESTED
	mux.HandleFunc("GET /api/report", cfg.handlerGetReport)                      // TODO
	mux.HandleFunc("GET /api/rooms", cfg.handlerGetRooms)                        // TODO
	mux.HandleFunc("POST /api/descriptions", cfg.handlerUpdateCourseDescription) // TODO
	mux.HandleFunc("POST /api/courses", cfg.handlerCreateCourse)                 // TODO
	mux.HandleFunc("POST /api/archive", cfg.handlerDeleteCourse)                 // TODO

	// START SERVICE
	server.ListenAndServe()
}
