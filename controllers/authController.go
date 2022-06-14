package controllers

import (
	"BI.ZONE_test/handlers"
	"BI.ZONE_test/models"
	"BI.ZONE_test/services/auth"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var Register = func(w http.ResponseWriter, r *http.Request) {
	newAccount := models.RegisterCredentials{}
	err := json.NewDecoder(r.Body).Decode(&newAccount)
	if err != nil {
		handlers.HandleBadRequest(w, err)
		return
	}

	if newAccount.Role != models.ADMIN_ROLE && newAccount.Role != models.ANALYST_ROLE {
		handlers.HandleBadRequest(w, fmt.Errorf("Invalid role "))
		return
	}

	err = auth.RegisterUser(newAccount)
	if err != nil {
		handlers.HandleInternalServerError(w, err)
		return
	}

	handlers.Respond(w, handlers.Message(true, "User has been registered!"))
}

var Login = func(w http.ResponseWriter, r *http.Request) {
	account := &models.LoginCredentials{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		handlers.HandleBadRequest(w, err)
		return
	}

	resp, err := auth.LoginUser(account.Username, account.Password)
	if err != nil {
		handlers.HandleBadRequest(w, err)
		return
	}

	handlers.Respond(w, resp)
}

var Verification = func(w http.ResponseWriter, r *http.Request) {
	stringId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(stringId)
	if err != nil {
		handlers.HandleBadRequest(w, fmt.Errorf("Invalid ID "))
		return
	}
	err = auth.Verify(id)
	if err != nil {
		handlers.HandleBadRequest(w, err)
		return
	}
	handlers.Respond(w, handlers.Message(true, "User has been verified"))
}
