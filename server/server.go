package server

type Server struct {
	repo RepositoryInterface
}

func NewServerService(repo RepositoryInterface) *Server {
	return &Server{repo}
}

func (s *Server) Create(title, body string) error {
	return s.repo.Create(title, body)
}
