// This file contains the repository implementation layer.
package repository

import (
	"github.com/ayyaa/todo-services/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB  *gorm.DB
	Cfg config.Config
}

type NewRepositoryOptions struct {
	Cfg config.Config
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := gorm.Open(postgres.Open(opts.Cfg.DB.ConnectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Repository{
		DB:  db,
		Cfg: opts.Cfg,
	}
}
