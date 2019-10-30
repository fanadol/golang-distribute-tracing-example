package server

import (
	"encoding/json"
	"net/http"

	"github.com/fanadol/golang-distribute-tracing-example/models"
)

type ServerHandler struct {
	service ServiceInterface
}

func NewServerHandler(service ServiceInterface) *ServerHandler {
	return &ServerHandler{service}
}

func (s *ServerHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	err = s.service.Create(body.Title, body.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(body)
	w.Write(data)
	return
}
