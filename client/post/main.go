package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fanadol/golang-distribute-tracing-example/models"
	"github.com/fanadol/golang-distribute-tracing-example/tracing"
	"github.com/nats-io/stan.go"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main() {
	tracer, closer := tracing.Init("Post-Client")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("Post-Client-Root")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	Post(ctx)
	Publish(ctx)
}

func Post(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Post-Request")
	defer span.Finish()

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
	ext.HTTPMethod.Set(span, "POST")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header),
	)

	_, err = http.DefaultClient.Do(request)
	if err != nil {
		panic("Error when trying to do request: " + err.Error())
	}

	fmt.Println("Success HTTP POST!")
}

func Publish(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Publish-Message")
	defer span.Finish()

	sc, err := stan.Connect("test-cluster", "client-clientID")
	if err != nil {
		panic("Error when connecting to stan: " + err.Error())
	}

	var t tracing.TraceMsg

	msg := []byte("Hellow Sir")

	// Inject span context into our traceMsg.
	if err := span.Tracer().Inject(span.Context(), opentracing.Binary, &t); err != nil {
		log.Fatalf("%v for Inject.", err)
	}

	t.Write(msg)

	err = sc.Publish("foo", t.Bytes())
	if err != nil {
		panic("Error when trying to publish: " + err.Error())
	}

	fmt.Println("Success Publish to Broker")
}
