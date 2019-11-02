package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fanadol/golang-distribute-tracing-example/models"
	"github.com/fanadol/golang-distribute-tracing-example/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func main() {
	tracer, closer := tracing.Init("Get-Client")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	span := tracer.StartSpan("Get-Client-Root")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	Get(ctx)
}

func Get(ctx context.Context) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Get-Request")
	defer span.Finish()

	var data []models.Post
	client := &http.Client{Timeout: 10 * time.Second}
	url := "http://localhost:8080/post"

	request, err := http.NewRequest("GET", url, nil)
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

	response, err := client.Do(request)
	if err != nil {
		panic("Error when trying to do request: " + err.Error())
	}

	err = json.NewDecoder(response.Body).Decode(&data)
	fmt.Println(fmt.Sprintf("%+v", data))
}
