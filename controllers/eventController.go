package controllers

import (
	"BI.ZONE_test/handlers"
	"BI.ZONE_test/services/event"
	"BI.ZONE_test/utils"
	"fmt"
	"net/http"
	"strconv"
)

var GetAllEvents = func(w http.ResponseWriter, r *http.Request) {
	order := r.FormValue("_order")
	sort := r.FormValue("_sort")
	end, err1 := strconv.Atoi(r.FormValue("_end"))
	start, err2 := strconv.Atoi(r.FormValue("_start"))
	filterParam := r.FormValue("_param")
	filterBy := r.FormValue("_filter")

	if err1 != nil || err2 != nil {
		handlers.HandleBadRequest(w, fmt.Errorf("bad _start or _end parameter value"))
		return
	}
	utils.CheckOrderAndSortParams(&order, &sort)
	eventsList, err := event.GetEventsList(&sort, &order, start, end, &filterParam, &filterBy)
	if err != nil {
		handlers.HandleInternalServerError(w, err)
		return
	}
	if err != nil {
		handlers.HandleBadRequest(w, fmt.Errorf("Invalid ID "))
		return
	}
	count, err := event.GetEventsAmount(&filterParam, &filterBy)
	if err != nil {
		handlers.HandleInternalServerError(w, err)
		return
	}

	utils.SetTotalCountHeader(w, strconv.FormatInt(count, 10))
	handlers.Respond(w, eventsList)
}
