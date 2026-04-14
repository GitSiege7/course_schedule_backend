package main

import "net/http"

func (cfg *apiConfig) handlerGetRooms(w http.ResponseWriter, r *http.Request) {
	type request struct {
		time string
		days string
	}
	var req request

	req.time = r.URL.Query().Get("timeToCheck")
	req.days = r.URL.Query().Get("daysToCheck")
}
