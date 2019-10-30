package server

// ServiceInterface represent service contract
type ServiceInterface interface {
	Create(title, body string) error
}
