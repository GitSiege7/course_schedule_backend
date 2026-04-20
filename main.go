package main

import (
	"net/http"

	"github.com/rs/cors"
)

// PORT: 9517

func main() {

	// CREATE CFG TO HOLD API STATE
	cfg := &apiConfig{
		uri:       "mongodb+srv://root:root@cen4020.qtbkpgn.mongodb.net/?appName=CEN4020",
		db_name:   "bellini_classes_db",
		coll_name: "bellini_classes",
	}

	// DB CONNECT
	cfg.dbConnect()

	// DEFER DB DISCONNECT
	defer cfg.dbDisconnect()

	// CREATE HTTP SERVICE
	mux := http.NewServeMux()
	handler := cors.Default().Handler(mux)

	server := http.Server{
		Handler: handler,
		Addr:    ":9517",
	}

	// SET HANDLER FUNCTIONS
	mux.HandleFunc("GET /api/courses", cfg.handlerGenericSearch)
	mux.HandleFunc("GET /api/schedule", cfg.handlerGetProfessorSchedule)
	mux.HandleFunc("GET /api/report", cfg.handlerGetReport)
	mux.HandleFunc("GET /api/rooms", cfg.handlerGetRooms)
	mux.HandleFunc("GET /api/all", cfg.handlerGetAllCourses)
	mux.HandleFunc("POST /api/descriptions", cfg.handlerUpdateCourseDescription)
	mux.HandleFunc("POST /api/courses", cfg.handlerCreateCourse)
	mux.HandleFunc("POST /api/archive", cfg.handlerDeleteCourse)

	// START SERVICE
	server.ListenAndServe()
}
