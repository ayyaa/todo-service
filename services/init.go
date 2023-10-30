// This file contains the repository implementation layer.
package services

import (
	"github.com/ayyaa/todo-services/config"
	"github.com/ayyaa/todo-services/repository"
)

type Service struct {
	Repository repository.RepositoryInterface
	Cfg        config.Config
}

type NewServiceOptions struct {
	Repository repository.RepositoryInterface
	Cfg        config.Config
}

func NewService(opts NewServiceOptions) *Service {

	return &Service{
		Repository: opts.Repository,
		Cfg:        opts.Cfg,
	}
}
