package server

// RepositoryInterface represent repository contract
type RepositoryInterface interface {
	Create(title, body string) error
}
