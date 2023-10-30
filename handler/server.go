package handler

import (
	"github.com/ayyaa/todo-services/config"
	"github.com/ayyaa/todo-services/repository"
	"github.com/ayyaa/todo-services/services"
)

type Server struct {
	Repository repository.RepositoryInterface
	Service    services.ServiceInterface
	Cfg        config.Config
}

type NewServerOptions struct {
	Repository repository.RepositoryInterface
	Service    services.ServiceInterface
	Cfg        config.Config
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Repository: opts.Repository,
		Service:    opts.Service,
		Cfg:        opts.Cfg,
	}
}
