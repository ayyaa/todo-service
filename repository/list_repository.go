package repository

import (
	"context"

	"github.com/ayyaa/todo-services/lang"
	"github.com/ayyaa/todo-services/models"
	"gorm.io/gorm"
)

func (r *Repository) GetListPreloadByID(ctx context.Context, id uint) (*models.List, error) {
	var list models.List
	// get list by spesific id only list where parent is null
	if err := r.DB.WithContext(ctx).Preload("Attachments", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ? ", "active")
	}).Preload("SubLists", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ? ", "active").Order("priority ASC").Order("id DESC").Preload("Attachments", func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ? ", "active")
		})
	}).Where("parent_id IS NULL").Where("status = ? ", "active").First(&list, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, lang.ErrListNotFound
		}
		return nil, err
	}
	return &list, nil
}

func (r *Repository) GetLists(ctx context.Context, filter models.ParamRequest) ([]models.List, *models.Pagination, error) {
	var (
		list       []models.List
		pagination models.Pagination
		totalItems int64
		// totalPages int64
	)

	offset := (filter.Page - 1) * filter.Size

	// Query the database with pagination, filtering, sorting, and optional joins
	query := r.DB.Order("priority ASC").Order("id DESC").Where("parent_id IS NULL").Where("status = ?", "active")

	if filter.Preload {
		query = query.Preload("Attachments", func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ? ", "active")
		}).Preload("SubLists", func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ? ", "active").Order("priority ASC").Order("id DESC").Preload("Attachments", func(db *gorm.DB) *gorm.DB {
				return db.Where("status = ? ", "active")
			})
		})
	} else {
		query = query.Preload("Attachments", func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ? ", "active")
		})
	}

	if filter.Keyword != "" && filter.OrderBy != "" {
		query = query.Where(filter.OrderBy+" LIKE ?", "%"+filter.Keyword+"%")
	}

	result := query.Find(&list)

	// count totalItems by filter
	result = result.Count(&totalItems)
	pagination.TotalItems = int(totalItems)

	result = result.Offset(offset).Limit(filter.Size).Find(&list)

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

	return list, &pagination, nil
}

func (r *Repository) GetListByID(ctx context.Context, id uint) (*models.List, error) {
	var list models.List
	// get list by spesific id only list where parent is null
	if err := r.DB.WithContext(ctx).Preload("Attachments", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ? ", "active")
	}).Where("parent_id IS NULL").Where("status = ? ", "active").First(&list, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, lang.ErrListNotFound
		}
		return nil, err
	}
	return &list, nil
}

func (r *Repository) CreateTx(ctx context.Context, list *models.List, attachments []*models.Attachment) (*models.List, error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		// insert list to db
		if err := tx.WithContext(ctx).Create(list).Error; err != nil {
			return err
		}

		newAttach := []*models.Attachment{}
		if len(attachments) != 0 {
			// mapping new list id for attachment
			for _, attachment := range attachments {
				attachment.ListID = list.ID

				newAttach = append(newAttach, attachment)
			}

			// insert attachment to db
			if err := tx.WithContext(ctx).Create(newAttach).Error; err != nil {
				return err
			}
		}

		list.Attachments = newAttach

		// return nil will commit the whole transaction
		return nil
	})

	return list, nil
}

func (r *Repository) MarkAsDeletedTx(ctx context.Context, id uint) error {
	list := models.List{}
	attachments := []*models.Attachment{}
	ids := []uint{}
	r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Model(&list).Where("id = ? OR parent_id = ?", id, id).Update("status", "deleted").Error
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Model(&list).Select("id").Where("id = ? OR parent_id = ?", id, id).Find(&ids).Error
		if err != nil {
			return err
		}

		err = tx.WithContext(ctx).Model(attachments).Where("list_id IN ?", ids).Update("status", "deleted").Error
		if err != nil {
			return err
		}

		return nil
	})

	list.Attachments = attachments

	return nil
}

func (r *Repository) UpdateTx(ctx context.Context, list *models.List, attachments []*models.Attachment, id uint) (*models.List, error) {
	r.DB.Transaction(func(tx *gorm.DB) error {
		// insert list to db
		if err := tx.WithContext(ctx).Model(&list).Updates(list).Error; err != nil {
			return err
		}

		if len(attachments) != 0 {
			// doft deleted old attachted file
			if err := tx.WithContext(ctx).Model(attachments).Where("list_id = ?", id).Update("status", "deleted").Error; err != nil {
				return err
			}

			// mapping new list id for attachment
			newAttach := []*models.Attachment{}
			for _, attachment := range attachments {
				attachment.ListID = id

				newAttach = append(newAttach, attachment)
			}

			// insert attachment to db
			if err := tx.WithContext(ctx).Create(newAttach).Error; err != nil {
				return err
			}

			list.Attachments = newAttach
		}

		return nil
	})

	return list, nil
}
