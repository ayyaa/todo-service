package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ayyaa/todo-services/config"
	"github.com/ayyaa/todo-services/models"
	"github.com/ayyaa/todo-services/repository"
	"github.com/ayyaa/todo-services/services"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

var mockSubList = models.SubList{
	ID:       1,
	ParentID: nil,
	Attachments: []*models.Attachment{
		{ID: 1,
			Filename: "active",
			Filepath: "/upload/1234.pdf"},
		{ID: 2,
			Filename: "active",
			Filepath: "/upload/1234.pdf"},
	},
}

var mockList = models.List{
	SubList: mockSubList,
	SubLists: []models.SubList{
		{
			ID:       1,
			ParentID: nil,
			Attachments: []*models.Attachment{
				{ID: 1,
					Filename: "active",
					Filepath: "/upload/1234.pdf"},
				{ID: 2,
					Filename: "active",
					Filepath: "/upload/1234.pdf"},
			},
		},
		{
			ID:       1,
			ParentID: nil,
			Attachments: []*models.Attachment{
				{ID: 1,
					Filename: "active",
					Filepath: "/upload/1234.pdf"},
				{ID: 2,
					Filename: "active",
					Filepath: "/upload/1234.pdf"},
			},
		},
	},
}

var mockListPagination = models.ListPagination{
	List: []models.List{
		mockList,
	},
	Pagination: models.Pagination{},
}

func TestServer_GetSubListByID(t *testing.T) {
	id := uint(1)
	type fields struct {
		Repository repository.RepositoryInterface
		Service    services.ServiceInterface
		Cfg        config.Config
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		doSetup func() (echo.Context, *httptest.ResponseRecorder)
		doMock  func(mocksrv services.MockServiceInterface)
		wantErr bool
	}{
		{
			name: "success",
			args: args{},
			doSetup: func() (echo.Context, *httptest.ResponseRecorder) {
				e := echo.New()
				req := httptest.NewRequest(http.MethodGet, "/sublists/:id", strings.NewReader(""))
				// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				c := e.NewContext(req, rec)
				c.SetParamNames("id")
				c.SetParamValues("1")
				return c, rec
			},
			doMock: func(mocksrv services.MockServiceInterface) {
				mocksrv.EXPECT().GetSubListByID(ctx, id).Return(&mockSubList, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Server{
				Repository: tt.fields.Repository,
				Service:    tt.fields.Service,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockSrv := services.NewMockServiceInterface(ctrl)
			c.Service = mockSrv

			ctxEcho, rec := tt.doSetup()
			tt.args.ctx = ctxEcho
			tt.doMock(*mockSrv)
			err := c.GetSubListByID(ctxEcho)
			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusOK, rec.Code)

				result := mockSubList
				err := json.Unmarshal([]byte(rec.Body.String()), &result)
				if err != nil {
					fmt.Println("Error unmarshal:", err)
				}
				assert.Equal(t, mockSubList, result)
			}
		})
	}
}

func TestServer_GetSubLists(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
		Service    services.ServiceInterface
		Cfg        config.Config
	}
	type args struct {
		ctx echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		doSetup func() (echo.Context, *httptest.ResponseRecorder)
		doMock  func(mocksrv services.MockServiceInterface)
		wantErr bool
	}{
		// {
		// 	name: "success",
		// 	args: args{},
		// 	doSetup: func() (echo.Context, *httptest.ResponseRecorder) {
		// 		e := echo.New()
		// 		req := httptest.NewRequest(http.MethodGet, "/sublists", strings.NewReader(""))
		// 		fmt.Println("tyest")
		// 		rec := httptest.NewRecorder()
		// 		c := e.NewContext(req, rec)
		// 		// Set the query parameters
		// 		q := req.URL.Query()
		// 		q.Set("page", "1")
		// 		q.Set("size", "10")
		// 		q.Set("orderBy", "title")
		// 		q.Set("keyword", "test")
		// 		req.URL.RawQuery = q.Encode()
		// 		fmt.Println(req.URL.RawQuery)

		// 		c.QueryParam(req.URL.RawQuery)
		// 		return c, rec
		// 	},
		// 	doMock: func(mocksrv services.MockServiceInterface) {
		// 		mocksrv.EXPECT().GetSubLists(ctx, gomock.Any()).Return(&mockListPagination, nil)
		// 	},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Server{
				Repository: tt.fields.Repository,
				Service:    tt.fields.Service,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockSrv := services.NewMockServiceInterface(ctrl)
			c.Service = mockSrv

			ctxEcho, rec := tt.doSetup()
			tt.args.ctx = ctxEcho
			tt.doMock(*mockSrv)
			err := c.GetSubLists(ctxEcho)
			if assert.NoError(t, err) {
				assert.Equal(t, http.StatusOK, rec.Code)

				result := mockListPagination
				err := json.Unmarshal([]byte(rec.Body.String()), &result)
				if err != nil {
					fmt.Println("Error unmarshal:", err)
				}
				assert.Equal(t, mockListPagination, result)
			}
		})
	}
}
