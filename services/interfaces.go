package services

import (
	"context"

	"github.com/ayyaa/todo-services/models"
)

//go:generate mockgen --destination=mocks/mock_service.go --package=mock_interfaces --source=interfaces.go

type ServiceInterface interface {
	GetListByID(ctx context.Context, id uint) (*models.List, error)
	CreateList(ctx context.Context, todo *models.AdditionalRequest) (*models.List, error)
	DeleteList(ctx context.Context, id uint) error
	EditList(ctx context.Context, list *models.AdditionalRequestEdit, id uint) (*models.List, error)
	GetLists(ctx context.Context, queryParam models.ParamRequest) (*models.ListPagination, error)

	GetSubListByID(ctx context.Context, id uint) (*models.SubList, error)
	EditSubList(ctx context.Context, list *models.AdditionalRequestEdit, id uint) (*models.List, error)
	DeleteSubList(ctx context.Context, id uint) error
	GetSubLists(ctx context.Context, queryParam models.ParamRequest) (*models.ListPagination, error)
}
