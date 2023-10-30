// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"

	"github.com/ayyaa/todo-services/models"
)

type RepositoryInterface interface {
	GetListByID(ctx context.Context, id uint) (*models.List, error)
	GetListPreloadByID(ctx context.Context, id uint) (*models.List, error)
	GetLists(ctx context.Context, filter models.ParamRequest) ([]models.List, *models.Pagination, error)

	GetSubListByID(ctx context.Context, id uint) (*models.SubList, error)
	GetSubListPreloadByID(ctx context.Context, id uint) (*models.SubList, error)
	GetSubLists(ctx context.Context, filter models.ParamRequest) ([]models.List, *models.Pagination, error)

	CreateTx(ctx context.Context, list *models.List, attachments []*models.Attachment) (*models.List, error)
	MarkAsDeletedTx(ctx context.Context, id uint) error
	UpdateTx(ctx context.Context, list *models.List, attachments []*models.Attachment, id uint) (*models.List, error)
}
