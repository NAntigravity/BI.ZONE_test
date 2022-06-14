package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func HandleBadRequest(w http.ResponseWriter, err error) {
	log.Error(err)
	w.WriteHeader(http.StatusBadRequest)
	Respond(w, Message(false, err.Error()))
}

func HandleInternalServerError(w http.ResponseWriter, err error) {
	log.Error(err)
	w.WriteHeader(http.StatusInternalServerError)
	Respond(w, Message(false, err.Error()))
}
