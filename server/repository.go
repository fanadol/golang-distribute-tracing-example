package server

import "github.com/fanadol/golang-distribute-tracing-example/models"

// RepositoryInterface represent repository contract
type RepositoryInterface interface {
	Create(title, body string) error
	Get() ([]models.Post, error)
}
