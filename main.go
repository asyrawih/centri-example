package main

import (
	"log"

	"github.com/centrifugal/centrifuge-go"
)

type MessageData struct {
	Message string
}

// main function  
func main() {

	go func() {
		log.Printf("Running in here")
	}()

	client := centrifuge.NewJsonClient(
		"ws://localhost:8000/connection/websocket",
		centrifuge.Config{
			Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJpbmRvZGF4IiwiZXhwIjoxNjYzOTA4NTY4LCJpYXQiOjE2NjMzMDM3Njh9.clYbRihfMRxATwYYSVw6ImRMbhi7L3mRAlw_Wcxj2gY",
		})

	defer client.Close()
	// Listening event client try to connecting the server
	client.OnConnecting(OnConnecting)
	// Listening event client try to connected the server
	client.OnConnected(OnConnected)
	// Listening event client try to disconnected the server
	client.OnDisconnected(OnDisconneted)

	if err := client.Connect(); err != nil {
		log.Printf("Error Happend %s", err)
	}

	// Subs Some Channel In Here
	sub, err := client.NewSubscription(
		"some",
		centrifuge.SubscriptionConfig{
			Recoverable: true,
			JoinLeave:   true,
		},
	)

	sub.OnJoin(OnJoin)

	sub.OnLeave(OnLeave)

	sub.OnPublication(OnPublication)

	sub.OnSubscribing(OnSubscribing)

	sub.OnSubscribed(OnSubscribed)

	if err != nil {
		log.Printf("Some Eror %s", err)
	}

	if err := sub.Subscribe(); err != nil {
		log.Printf("Error Subs , %s", err)
	}

	select {}

}

// OnSubscribed function  
func OnSubscribed(e centrifuge.SubscribedEvent) {
	log.Printf("Client Has Subscribe, %s", e.Data)
}

// OnSubscribing function  
func OnSubscribing(e centrifuge.SubscribingEvent) {
	log.Printf("Client Has Subscribe, %s", e.Reason)
}

// OnJoin function  
func OnJoin(je centrifuge.JoinEvent) {
	log.Printf("Some Join %s", je.User)
}

// OnLeave function  
func OnLeave(je centrifuge.LeaveEvent) {
	log.Printf("Client Leave %s", je.ChanInfo)
}

// OnPublication function  
func OnPublication(je centrifuge.PublicationEvent) {
	log.Printf("%s | %s", string(je.Data), je.Info.User)
}

// OnConnecting function  
func OnConnecting(ce centrifuge.ConnectingEvent) {
	log.Printf("Connecting - %d (%s)", ce.Code, ce.Reason)
}

// OnConnected function  
func OnConnected(ce centrifuge.ConnectedEvent) {
	log.Printf("Connected - %s ", ce.ClientID)
}

// OnDisconneted function  
func OnDisconneted(ce centrifuge.DisconnectedEvent) {
	log.Printf("Disconnected - %d (%s)", ce.Code, ce.Reason)
}
