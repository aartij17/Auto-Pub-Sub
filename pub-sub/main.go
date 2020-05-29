package main

import (
	"Auto-Pub-Sub/pub-sub/client"
	"Auto-Pub-Sub/pub-sub/server"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
)

func main() {

	client.RunClient()
	server.RunServer()

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			log.Info("Received an interrupt, stopping all connections...")
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
