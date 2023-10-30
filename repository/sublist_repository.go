package repository

import (
	"context"

	"github.com/ayyaa/todo-services/lang"
	"github.com/ayyaa/todo-services/models"
	"gorm.io/gorm"
)

func (r *Repository) GetSubListPreloadByID(ctx context.Context, id uint) (*models.SubList, error) {
	var sublist models.SubList
	// get sublist by spesific id only sublist where parent is not null
	if err := r.DB.WithContext(ctx).Preload("Attachments", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ? ", "active")
	}).Where("id = ? AND parent_id IS NOT NULL AND status = ?", id, "active").First(&sublist).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, lang.ErrSubListNotFound
		}
		return nil, err
	}

	return &sublist, nil
}

func (r *Repository) GetSubLists(ctx context.Context, filter models.ParamRequest) ([]models.List, *models.Pagination, error) {
	var (
		sublist    []models.List
		pagination models.Pagination
		totalItems int64
	)

	offset := (filter.Page - 1) * filter.Size

	// Query the database with pagination, filtering, sorting, and optional joins
	query := r.DB.Order("priority ASC").Order("id DESC").Where("parent_id IS NOT NULL").Where("status = ?", "active")

	query = query.Preload("Attachments", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ? ", "active")
	})

	if filter.Keyword != "" && filter.OrderBy != "" {
		query = query.Where(filter.OrderBy+" LIKE ?", "%"+filter.Keyword+"%")
	}

	result := query.Find(&sublist)

	// count totalItems by filter
	result = result.Count(&totalItems)
	pagination.TotalItems = int(totalItems)

	result = result.Offset(offset).Limit(filter.Size).Find(&sublist)

	if result.Error != nil {
		return nil, nil, result.Error
	}
	// Calculate the total number of pages
	if totalItems%int64(filter.Size) == 0 {
		pagination.TotalPages = int(totalItems / int64(filter.Size))
	} else {
		pagination.TotalPages = int(totalItems/int64(filter.Size)) + 1
	}

	pagination.Page = filter.Page
	pagination.Size = filter.Size

	return sublist, &pagination, nil
}

func (r *Repository) GetSubListByID(ctx context.Context, id uint) (*models.SubList, error) {
	var sublist models.SubList
	// get sublist by spesific id only sublist where parent is not null
	if err := r.DB.WithContext(ctx).Where("id = ? AND parent_id IS NOT NULL", id).Where("status = ? ", "active").First(&sublist).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, lang.ErrSubListNotFound
		}
		return nil, err
	}

	return &sublist, nil
}
