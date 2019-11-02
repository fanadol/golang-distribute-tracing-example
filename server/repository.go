package server

import (
	"context"

	"github.com/fanadol/golang-distribute-tracing-example/models"
)

// RepositoryInterface represent repository contract
type RepositoryInterface interface {
	Create(ctx context.Context, title, body string) error
	Get(ctx context.Context) ([]models.Post, error)
}
