package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/fanadol/golang-distribute-tracing-example/tracing"
	stan "github.com/nats-io/stan.go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const (
	URL       = stan.DefaultNatsURL
	clientID  = "client-id"
	clusterID = "test-cluster"
)

func main() {
	tracer, closer := tracing.Init("Post-Client")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	sc, err := stan.Connect(clusterID, clientID)
	if err != nil {
		panic("Error when connecting to stan: " + err.Error())
	}

	sub, err := sc.Subscribe("foo", func(m *stan.Msg) {
		t := tracing.NewTraceMsg(m)

		spanCtx, err := tracer.Extract(opentracing.Binary, t)
		if err != nil {
			log.Fatalf("Some error has occur: %v", err.Error())
		}

		// Setup a span referring to the span context of the incoming NATS message.
		span := tracer.StartSpan("Received-Message", ext.SpanKindConsumer, opentracing.FollowsFrom(spanCtx))
		defer span.Finish()

		fmt.Printf("Receive a message: %s\n", string(m.Data))

		// Set sleep for better observation in opentracing
		time.Sleep(50 * time.Millisecond)
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
