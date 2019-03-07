package service

import "reference-implementation/schemod/service/query"

// Server represents a service server instance
type Server interface {
	Serve(Transport) error
	Shutdown() error
	Query(query.Query) (query.Result, error)
}

type server struct {
	engine Engine
}

// NewServer creates a new service server instance
func NewServer(engine Engine) Server {
	return &server{engine: engine}
}

func (srv *server) Serve(transport Transport) error {
	return nil
}

func (srv *server) Shutdown() error {
	return nil
}

func (srv *server) Query(q query.Query) (query.Result, error) {
	return query.Result{}, nil
}
