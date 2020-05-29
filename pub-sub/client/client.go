package client

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

var (
	natsServer = "nats://localhost:4222"
	// uncomment this when Dockerized
	// natsServer = "nats://cloud-nats-svc:4222"
)

func RunClient() {
	natsOptions := nats.Options{
		Servers:        []string{natsServer},
		AllowReconnect: true,
	}
	nc, err := natsOptions.Connect()

	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("error connecting to nats server")
		return
	}

	go func() {
		numRequest := 1
		ticker := time.NewTicker(5 * time.Second)
		log.Info("[CLIENT]: publishing request to the server every 5 seconds")
		for {
			select {
			case <-ticker.C:
				request := fmt.Sprintf("request #%d", numRequest)
				err = nc.PublishRequest("natsInbox", "", []byte(request))
				log.WithFields(log.Fields{
					"message": request,
				}).Info("[CLIENT] published request to server")
				numRequest += 1
			}
		}
	}()

}
