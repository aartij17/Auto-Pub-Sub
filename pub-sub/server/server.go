package server

import (
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

var (
	natsServer = "nats://localhost:4222"
	// uncomment this when Dockerized
	// natsServer = "nats://cloud-nats-svc:4222"
)

func RunServer() {
	natsOptions := nats.Options{
		Servers: []string{natsServer},
	}

	nc, err := natsOptions.Connect()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("error connecting to the NATS server")
		return
	}

	_, err = nc.Subscribe("natsInbox", func(m *nats.Msg) {
		log.WithFields(log.Fields{
			"message": string(m.Data),
		}).Info("[SERVER] Received a message from client")
	})

}
