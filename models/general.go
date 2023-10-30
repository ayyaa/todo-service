package models

import (
	"mime/multipart"
)

type ErrorResponse struct {
	Data    *map[string]string `json:"data,omitempty"`
	Message string             `json:"message"`
}

type SuccessResponse struct {
	Message    string      `json:"message"`
	Pagination interface{} `json:"pagination,omitempty"`
	Data       interface{} `json:"data"`
}

type Pagination struct {
	Size       int `json:"size"`
	Page       int `json:"page"`
	TotalPages int `json:"totalPages"`
	TotalItems int `json:"totalItems"`
}

type ListRequest struct {
	Title       string                  `form:"title" validate:"required,max=100,alphanumeric"`
	Description string                  `form:"description" validate:"required,max=1000"`
	Priority    int                     `form:"priority" `
	File        []*multipart.FileHeader `form:"file" validate:"extention"`
}

type AdditionalRequest struct {
	ListRequest
	ParentID uint `form:"parentID" validate:"required"`
}

type ListRequestEdit struct {
	Title       string                  `form:"title" validate:"max=100,alphanumeric"`
	Description string                  `form:"description" validate:"max=1000"`
	Priority    int                     `form:"priority" `
	File        []*multipart.FileHeader `form:"file,omitempty" validate:"extention"`
}

type AdditionalRequestEdit struct {
	ListRequestEdit
	ParentID uint `form:"parentID" `
}

type ParamRequest struct {
	OrderBy string `json:"orderBy" query:"orderBy"`
	Keyword string `json:"keyword" query:"keyword"`
	Size    int    `json:"size" query:"size"`
	Page    int    `json:"page" query:"page"`
	Preload bool   `json:"preload" query:"preload,omitempty"`
}

type ListPagination struct {
	List       []List     `json:"data"`
	Pagination Pagination `json:"pagination"`
}
