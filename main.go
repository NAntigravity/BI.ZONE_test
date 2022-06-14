package main

import (
	"BI.ZONE_test/routers"
	"BI.ZONE_test/tcp"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	go tcp.StartTCPServer()
	router := routers.InitRouter()
	log.Infof("HTTP server has been started")
	log.Fatal(http.ListenAndServe(":8080", router))
}
