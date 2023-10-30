package services

import (
	"context"
	"fmt"

	errors "github.com/ayyaa/todo-services/lib/customerrors"
	"github.com/ayyaa/todo-services/lib/helper"
	"github.com/ayyaa/todo-services/models"
)

func (s *Service) GetListByID(ctx context.Context, id uint) (*models.List, error) {
	return s.Repository.GetListPreloadByID(ctx, id)
}

func (s *Service) GetLists(ctx context.Context, queryParam models.ParamRequest) (*models.ListPagination, error) {
	lists, paginate, err := s.Repository.GetLists(ctx, queryParam)
	if err != nil {
		return nil, err
	}

	dataPagination := models.ListPagination{
		List:       lists,
		Pagination: *paginate,
	}
	return &dataPagination, nil
}

func (s *Service) GetSubListByID(ctx context.Context, id uint) (*models.SubList, error) {
	return s.Repository.GetSubListPreloadByID(ctx, id)
}

func (s *Service) GetSubLists(ctx context.Context, queryParam models.ParamRequest) (*models.ListPagination, error) {
	lists, paginate, err := s.Repository.GetSubLists(ctx, queryParam)
	if err != nil {
		return nil, err
	}

	dataPagination := models.ListPagination{
		List:       lists,
		Pagination: *paginate,
	}
	return &dataPagination, nil
}

func (s *Service) CreateList(ctx context.Context, list *models.AdditionalRequest) (*models.List, error) {
	var err error

	listData := models.List{
		SubList: models.SubList{
			Title:       list.Title,
			Description: list.Description,
			Priority:    list.Priority,
		},
	}

	if list.ParentID != 0 {
		listData.ParentID = &list.ParentID

		// check if add sublist with the correct list
		_, err := s.Repository.GetListByID(ctx, list.ParentID)
		if err != nil {
			return nil, errors.NewBadRequestErrorf("List id %d not found, please make sure corrent list id", list.ParentID)
		}
	}

	createdAttachment := []*models.Attachment{}
	fmt.Println(list.ListRequest.File)
	if len(list.ListRequest.File) != 0 {

		// upload file to s3
		createdAttachment, err = helper.Upload(ctx, &list.ListRequest, s.Cfg.AWS)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}
	}

	// create list and attachment with transaction
	createdList, err := s.Repository.CreateTx(ctx, &listData, createdAttachment)
	if err != nil {
		return nil, err
	}

	createdList.Attachments = createdAttachment

	return createdList, nil
}

func (s *Service) DeleteList(ctx context.Context, id uint) error {
	// check if add sublist with the correct list
	_, err := s.Repository.GetListByID(ctx, id)
	if err != nil {
		return errors.NewBadRequestErrorf("List id %d not found, please make sure corrent list id", id)

	}

	// create list and attachment with transaction
	err = s.Repository.MarkAsDeletedTx(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteSubList(ctx context.Context, id uint) error {
	// check if add sublist with the correct list
	_, err := s.Repository.GetSubListByID(ctx, id)
	if err != nil {
		return errors.NewBadRequestErrorf("Sublist id %d not found, please make sure corrent sub list id", id)
	}

	// create list and attachment with transaction
	err = s.Repository.MarkAsDeletedTx(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) EditList(ctx context.Context, list *models.AdditionalRequestEdit, id uint) (*models.List, error) {
	// check correct list
	listByID, err := s.Repository.GetListByID(ctx, id)
	if err != nil {
		return nil, errors.NewBadRequestErrorf("List id %d not found, please make sure corrent list id", id)
	}

	listData := models.List{
		SubList: models.SubList{
			Title:       listByID.Title,
			Description: listByID.Description,
			Priority:    listByID.Priority,
			ID:          id,
		},
	}

	if list.Title != "" {
		listData.Title = list.Title
	}

	if list.Description != "" {
		listData.Description = list.Description
	}

	if list.Priority != 0 {
		listData.Priority = list.Priority
	}

	createdAttachment := []*models.Attachment{}
	if len(list.ListRequestEdit.File) != 0 {
		attach := models.ListRequest{
			Title:       list.ListRequestEdit.Title,
			Description: list.ListRequestEdit.Description,
			File:        list.ListRequestEdit.File,
			Priority:    list.ListRequestEdit.Priority,
		}
		// upload file to s3
		createdAttachment, err = helper.Upload(ctx, &attach, s.Cfg.AWS)
		if err != nil {
			return nil, err
		}
	}

	// create list and attachment with transaction
	createdList, err := s.Repository.UpdateTx(ctx, &listData, createdAttachment, id)
	if err != nil {
		return nil, err
	}

	if len(list.ListRequestEdit.File) != 0 {
		createdList.Attachments = createdAttachment
	} else {
		createdList.Attachments = listByID.Attachments
	}

	return createdList, nil
}

func (s *Service) EditSubList(ctx context.Context, list *models.AdditionalRequestEdit, id uint) (*models.List, error) {
	// check correct list
	listByID, err := s.Repository.GetSubListByID(ctx, id)
	if err != nil {
		return nil, errors.NewBadRequestErrorf("Sublist id %d not found, please make sure corrent sublist id", id)
	}

	listData := models.List{
		SubList: models.SubList{
			Title:       listByID.Title,
			Description: listByID.Description,
			Priority:    listByID.Priority,
			ID:          id,
			ParentID:    listByID.ParentID,
		},
	}

	if list.Title != "" {
		listData.Title = list.Title
	}

	if list.Description != "" {
		listData.Description = list.Description
	}

	if list.Priority != 0 {
		listData.Priority = list.Priority
	}

	if list.ParentID != 0 {
		listData.ParentID = &list.ParentID
	}

	createdAttachment := []*models.Attachment{}
	if len(list.ListRequestEdit.File) != 0 {
		attach := models.ListRequest{
			Title:       list.ListRequestEdit.Title,
			Description: list.ListRequestEdit.Description,
			File:        list.ListRequestEdit.File,
			Priority:    list.ListRequestEdit.Priority,
		}
		// upload file to s3
		createdAttachment, err = helper.Upload(ctx, &attach, s.Cfg.AWS)
		if err != nil {
			return nil, err
		}
	}

	// create list and attachment with transaction
	createdList, err := s.Repository.UpdateTx(ctx, &listData, createdAttachment, id)
	if err != nil {
		return nil, err
	}

	if len(list.ListRequestEdit.File) != 0 {
		createdList.Attachments = createdAttachment
	} else {
		createdList.Attachments = listByID.Attachments
	}

	return createdList, nil
}
