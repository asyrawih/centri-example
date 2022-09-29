package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/centrifugal/centrifuge-go"
	"github.com/google/uuid"
)

type MessageData struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

func main() {

	client := centrifuge.NewJsonClient(
		"ws://localhost:8000/connection/websocket",
		centrifuge.Config{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJpbmRvZGF4IiwiZXhwIjoxNjYzOTA4NTY4LCJpYXQiOjE2NjMzMDM3Njh9.clYbRihfMRxATwYYSVw6ImRMbhi7L3mRAlw_Wcxj2gY",
		})

	defer client.Close()
	// Listening event client try to connecting the server
	client.OnConnecting(func(ce centrifuge.ConnectingEvent) {})
	// Listening event client try to connected the server
	client.OnConnected(func(ce centrifuge.ConnectedEvent) {})
	// Listening event client try to disconnected the server
	client.OnDisconnected(func(de centrifuge.DisconnectedEvent) {})

	if err := client.Connect(); err != nil {
		log.Printf("Error Happend %s", err)
	}
	sub, err := client.NewSubscription(
		"some",
		centrifuge.SubscriptionConfig{
			Recoverable: true,
			JoinLeave:   true,
		},
	)

	if err != nil {
		log.Printf("error %s", err)
	}

	pubMessage := func(text string) error {
		message := &MessageData{Id: uuid.New().String(), Message: text}
		data, _ := json.Marshal(message)
		_, err := sub.Publish(context.Background(), data)
		return err
	}

	sub.OnPublication(func(pe centrifuge.PublicationEvent) {})

	sub.OnSubscribing(func(se centrifuge.SubscribingEvent) {})

	sub.OnSubscribed(func(se centrifuge.SubscribedEvent) {})

	if err := sub.Subscribe(); err != nil {
		log.Printf("Cannot Subscribe the channel , %s", err)
	}

	for {
		time.Sleep(time.Duration(1) * time.Second)
		if err := pubMessage("Hello"); err != nil {
			log.Printf("Error Pubs the message %s", err)
		}
	}
}
