package main

import (
	"encoding/json"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
)

type Order struct {
	OrderId string
	Amount  int
}

type EventStore struct {
	EventId       string
	EventType     string
	EventData     string
	AggregateId   string
	AggregateType string
}

const (
	queue     = "Order.OrdersCreatedQueue"
	subject   = "Order.OrderCreated"
	aggregate = "Order"
	event     = "OrderCreated"
)

func main() {
	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)

	// Subscribe to subject
	natsConnection.Subscribe(subject, func(msg *nats.Msg) {
		eventStore := EventStore{}
		err := json.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			// Handle the message
			log.Printf("Received message in EventStore service 1: %+v\n", eventStore)
		}
	})

	natsConnection.Subscribe(subject, func(msg *nats.Msg) {
		eventStore := EventStore{}
		err := json.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			// Handle the message
			log.Printf("Received message in EventStore service 2: %+v\n", eventStore)
		}
	})

	// Subscribe to subject
	natsConnection.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		eventStore := EventStore{}
		err := json.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			// Handle the message
			log.Printf("Subscribed message in Worker 1: %+v\n", eventStore)
		}
	})

	// Subscribe to subject
	natsConnection.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		eventStore := EventStore{}
		err := json.Unmarshal(msg.Data, &eventStore)
		if err == nil {
			// Handle the message
			log.Printf("Subscribed message in Worker 2: %+v\n", eventStore)
		}
	})

	// Keep the connection alive
	runtime.Goexit()
}
