package server

import (
	"context"

	"github.com/fanadol/golang-distribute-tracing-example/models"
	"github.com/opentracing/opentracing-go"
)

type Server struct {
	repo RepositoryInterface
}

func NewServerService(repo RepositoryInterface) *Server {
	return &Server{repo}
}

func (s *Server) Create(ctx context.Context, title, body string) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Server-Create-Service")
	defer span.Finish()

	return s.repo.Create(ctx, title, body)
}

func (s *Server) Get(ctx context.Context) ([]models.Post, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "Server-Get-Service")
	defer span.Finish()

	return s.repo.Get(ctx)
}
