package controllers

import (
	"BI.ZONE_test/handlers"
	"BI.ZONE_test/services/event"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var IncidentHandler = func(w http.ResponseWriter, r *http.Request) {
	stringId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(stringId)
	if err != nil {
		handlers.HandleBadRequest(w, fmt.Errorf("Invalid ID "))
		return
	}
	err = event.MarkAsIncident(uint(id))
	if err != nil {
		handlers.HandleBadRequest(w, err)
		return
	}
	handlers.Respond(w, handlers.Message(true, "Event has been marked as incident"))
}

var CheckMessage = func(w http.ResponseWriter, r *http.Request) {
	stringId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(stringId)
	if err != nil {
		handlers.HandleBadRequest(w, fmt.Errorf("Invalid ID "))
		return
	}
	event, err := event.GetDecodedData(uint(id))
	if err != nil {
		handlers.HandleInternalServerError(w, err)
		return
	}
	handlers.Respond(w, event)
}
