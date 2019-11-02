package server

import (
	"encoding/json"
	"net/http"

	"github.com/fanadol/golang-distribute-tracing-example/models"
	"github.com/opentracing/opentracing-go"
)

type ServerHandler struct {
	service ServiceInterface
}

func NewServerHandler(service ServiceInterface) *ServerHandler {
	return &ServerHandler{service}
}

func (s *ServerHandler) Create(w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "Server-Create")
	defer span.Finish()
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
	span, ctx := opentracing.StartSpanFromContext(r.Context(), "Server-Get")
	defer span.Finish()

	posts, err := s.service.Get(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
