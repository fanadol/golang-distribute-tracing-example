package server

import (
	"context"

	"github.com/fanadol/golang-distribute-tracing-example/models"
)

// ServiceInterface represent service contract
type ServiceInterface interface {
	Create(ctx context.Context, title, body string) error
	Get(ctx context.Context) ([]models.Post, error)
}
