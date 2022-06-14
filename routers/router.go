package routers

import (
	"BI.ZONE_test/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api").Subrouter()
	api.Use(loggingHandler)
	api.Use(authCheck)
	initUserRouter(api)
	initAnalystRouter(api)
	initEventsRouter(api)
	return router
}

func initUserRouter(router *mux.Router) {
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/register", controllers.Register).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc("/login", controllers.Login).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc("/{id:[0-9]*}/verify", adminCheck(controllers.Verification)).Methods(http.MethodPost, http.MethodOptions)
	userRouter.HandleFunc("/{id:[0-9]*}/verify", adminCheck(controllers.Verification)).Methods(http.MethodPost, http.MethodOptions)
}

func initAnalystRouter(router *mux.Router) {
	userRouter := router.PathPrefix("/analyze").Subrouter()
	userRouter.Use(analystCheck)
	userRouter.HandleFunc("/{id:[0-9]*}/check", controllers.CheckMessage).Methods(http.MethodGet, http.MethodOptions)
	userRouter.HandleFunc("/{id:[0-9]*}/incident", controllers.IncidentHandler).Methods(http.MethodPost, http.MethodOptions)
}

func initEventsRouter(router *mux.Router) {
	eventsRouter := router.PathPrefix("/events").Subrouter()
	eventsRouter.HandleFunc("/list", controllers.GetAllEvents).Methods(http.MethodGet, http.MethodOptions)
}
