package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fanadol/golang-distribute-tracing-example/models"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type ServerHandler struct {
	service ServiceInterface
}

func NewServerHandler(service ServiceInterface) *ServerHandler {
	return &ServerHandler{service}
}

func (s *ServerHandler) Create(w http.ResponseWriter, r *http.Request) {
	spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := opentracing.GlobalTracer().StartSpan("Server-Create", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	body := models.Post{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if body.Title == "" || body.Body == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	err = s.service.Create(ctx, body.Title, body.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (s *ServerHandler) Get(w http.ResponseWriter, r *http.Request) {
	spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := opentracing.GlobalTracer().StartSpan("Server-Get", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	posts, err := s.service.Get(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
