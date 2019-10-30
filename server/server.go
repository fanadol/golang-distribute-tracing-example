package server

import "github.com/fanadol/golang-distribute-tracing-example/models"

type Server struct {
	repo RepositoryInterface
}

func NewServerService(repo RepositoryInterface) *Server {
	return &Server{repo}
}

func (s *Server) Create(title, body string) error {
	return s.repo.Create(title, body)
}

func (s *Server) Get() ([]models.Post, error) {
	return s.repo.Get()
}
