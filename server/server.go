package server

import "github.com/bwlee13/gopherdb/store/base"

type Service struct {
	addr  string
	store *base.Store
}

func NewService(addr string, store *base.Store) *Service {
	return &Service{
		addr:  addr,
		store: store,
	}
}

func (service *Service) Start() {
	// put server start and server configs in here
	// aka most of what was main.go
}
