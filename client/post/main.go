package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fanadol/golang-distribute-tracing-example/models"
	"github.com/nats-io/stan.go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main() {
	span, _ := opentracing.StartSpanFromContext(context.Background(), "Post-Client")
	defer span.Finish()
	// Create new post to server
	client := &http.Client{}
	post := models.Post{
		Title: "Cat and Dogs",
		Body:  "Once upon a time in hohaw land",
	}
	payload, _ := json.Marshal(post)
	url := "http://localhost:8080/post"

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		panic("Error when trying to create new request : " + err.Error())
	}

	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, url)
	ext.HTTPMethod.Set(span, "GET")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header),
	)

	_, err = client.Do(request)
	if err != nil {
		panic("Error when trying to do request: " + err.Error())
	}

	fmt.Println("Success HTTP POST!")

	// Publish message
	sc, err := stan.Connect("test-cluster", "client-clientID")
	if err != nil {
		panic("Error when connecting to stan: " + err.Error())
	}

	msg := "Hellow Sir"

	// Setup a span for the operation to publish a message.
	ext.MessageBusDestination.Set(span, "foo")

	// Inject span context into our traceMsg.
	if err := span.Tracer().Inject(span.Context(), opentracing.Binary, &msg); err != nil {
		log.Fatalf("%v for Inject.", err)
	}

	err = sc.Publish("foo", []byte(msg))
	if err != nil {
		panic("Error when trying to publish: " + err.Error())
	}

	fmt.Println("Success Publish to Broker")
}
