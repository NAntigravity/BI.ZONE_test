package tcp

import (
	"BI.ZONE_test/models"
	"BI.ZONE_test/services/event"
	"bufio"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net"
)

func StartTCPServer() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Infof("TCP server has been started")
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Infof("TCP client connected")
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	for {
		message, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			log.Infof("TCP client disconnected")
			return
		}
		var msg models.Event
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Error(err)
			continue
		}
		err = event.SaveEvent(msg)
		if err != nil {
			log.Error(err)
			continue
		}
	}
}
