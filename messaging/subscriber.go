package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	stan "github.com/nats-io/stan.go"
)

const (
	URL       = stan.DefaultNatsURL
	clientID  = "client-id"
	clusterID = "test-cluster"
)

func main() {
	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		panic("Error when connecting to stan: " + err.Error())
	}

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Receive a message: %s\n", string(m.Data))
	})

	if err != nil {
		panic("Error when trying to subscribe: " + err.Error())
	}

	log.Printf("Connected to %s clusterID: [%s] clientID: [%s]\n", URL, clusterID, clientID)

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			sub.Unsubscribe()
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}
