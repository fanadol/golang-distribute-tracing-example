package server

import "github.com/fanadol/golang-distribute-tracing-example/models"

// ServiceInterface represent service contract
type ServiceInterface interface {
	Create(title, body string) error
	Get() ([]models.Post, error)
}
