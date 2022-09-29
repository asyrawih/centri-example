package main

import (
	"encoding/json"
	"log"

	"github.com/centrifugal/centrifuge-go"
)

type MessageData struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

// main function  
func main() {

	go func() {
		log.Printf("Running in here")
	}()

	client := centrifuge.NewJsonClient(
		"ws://localhost:8000/connection/websocket",
		centrifuge.Config{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJoYW5hbiIsImV4cCI6MTY2NDQzNjcyNiwiaWF0IjoxNjYzODMxOTI2fQ.3vslSFCCQy_2XObPjE2-dtqICV9h_PV6qhNCiSgBOgU",
		})

	defer client.Close()
	// Listening event client try to connecting the server
	client.OnConnecting(func(ce centrifuge.ConnectingEvent) {})
	// Listening event client try to connected the server
	client.OnConnected(func(ce centrifuge.ConnectedEvent) {
	})
	// Listening event client try to disconnected the server
	client.OnDisconnected(func(de centrifuge.DisconnectedEvent) {})

	if err := client.Connect(); err != nil {
		log.Printf("Error Happend %s", err)
	}

	// Subs Some Channel In Here
	sub, err := client.NewSubscription(
		"some",
		centrifuge.SubscriptionConfig{
			JoinLeave: true,
		},
	)

	sub.OnJoin(func(je centrifuge.JoinEvent) {
		log.Printf("%s", string(je.ConnInfo))
	})

	sub.OnLeave(func(le centrifuge.LeaveEvent) {
		log.Printf("%s", string(le.ChanInfo))
	})

	sub.OnPublication(func(pe centrifuge.PublicationEvent) {
		var message MessageData
		if err := json.Unmarshal(pe.Data, &message); err != nil {
			log.Printf("Error Happen On Publication : %s", err)
		}
		log.Printf("message:%s  id: %s", message.Message, message.Id)
	})

	sub.OnSubscribing(func(se centrifuge.SubscribingEvent) {
		log.Printf("%s", se.Reason)
	})

	sub.OnSubscribed(func(se centrifuge.SubscribedEvent) {
		log.Printf("%s", string(se.Data))
	})

	if err != nil {
		log.Printf("Some Eror %s", err)
	}

	if err := sub.Subscribe(); err != nil {
		log.Printf("Error Subs , %s", err)
	}

	select {}

}
