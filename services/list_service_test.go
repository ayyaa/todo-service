package services

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/ayyaa/todo-services/config"
	"github.com/ayyaa/todo-services/models"
	"github.com/ayyaa/todo-services/repository"
	gomock "github.com/golang/mock/gomock"
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

func TestService_GetListByID(t *testing.T) {

	type fields struct {
		Repository repository.RepositoryInterface
		Cfg        config.Config
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(repo repository.MockRepositoryInterface)
		want    *models.List
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListPreloadByID(ctx, gomock.Any()).Return(&mockList, nil)
			},
			want:    &mockList,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListPreloadByID(ctx, gomock.Any()).Return(nil, sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repository: tt.fields.Repository,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockRepo := repository.NewMockRepositoryInterface(ctrl)

			s.Repository = mockRepo
			tt.doMock(*mockRepo)
			got, err := s.GetListByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetListByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetListByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetLists(t *testing.T) {
	response := models.ListPagination{
		List: []models.List{
			mockList,
		},
		Pagination: models.Pagination{},
	}
	type fields struct {
		Repository repository.RepositoryInterface
		Cfg        config.Config
	}
	type args struct {
		ctx        context.Context
		queryParam models.ParamRequest
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(repo repository.MockRepositoryInterface)
		want    *models.ListPagination
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				queryParam: models.ParamRequest{
					Page:    1,
					Size:    10,
					OrderBy: "title",
					Preload: false,
					Keyword: "ini",
				},
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetLists(ctx, gomock.Any()).Return([]models.List{
					mockList,
				}, &models.Pagination{}, nil)
			},
			want:    &response,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				ctx: ctx,
				queryParam: models.ParamRequest{
					Page:    1,
					Size:    10,
					OrderBy: "title",
					Preload: false,
					Keyword: "ini",
				},
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetLists(ctx, gomock.Any()).Return(nil, nil, sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repository: tt.fields.Repository,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockRepo := repository.NewMockRepositoryInterface(ctrl)

			s.Repository = mockRepo
			tt.doMock(*mockRepo)
			got, err := s.GetLists(tt.args.ctx, tt.args.queryParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetSubListByID(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
		Cfg        config.Config
	}
	ctx := context.Background()
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(repo repository.MockRepositoryInterface)
		want    *models.SubList
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetSubListPreloadByID(ctx, gomock.Any()).Return(&mockSubList, nil)
			},
			want:    &mockSubList,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				ctx: ctx,
				id:  1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetSubListPreloadByID(ctx, gomock.Any()).Return(nil, sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repository: tt.fields.Repository,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockRepo := repository.NewMockRepositoryInterface(ctrl)

			s.Repository = mockRepo
			tt.doMock(*mockRepo)
			got, err := s.GetSubListByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetSubListByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetSubListByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetSubLists(t *testing.T) {
	response := models.ListPagination{
		List: []models.List{
			mockList,
		},
		Pagination: models.Pagination{},
	}
	type fields struct {
		Repository repository.RepositoryInterface
		Cfg        config.Config
	}
	type args struct {
		ctx        context.Context
		queryParam models.ParamRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(repo repository.MockRepositoryInterface)
		want    *models.ListPagination
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				queryParam: models.ParamRequest{
					Page:    1,
					Size:    10,
					OrderBy: "title",
					Preload: false,
					Keyword: "ini",
				},
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetSubLists(ctx, gomock.Any()).Return([]models.List{
					mockList,
				}, &models.Pagination{}, nil)
			},
			want:    &response,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				ctx: ctx,
				queryParam: models.ParamRequest{
					Page:    1,
					Size:    10,
					OrderBy: "title",
					Preload: false,
					Keyword: "ini",
				},
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetSubLists(ctx, gomock.Any()).Return(nil, nil, sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repository: tt.fields.Repository,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockRepo := repository.NewMockRepositoryInterface(ctrl)

			s.Repository = mockRepo
			tt.doMock(*mockRepo)
			got, err := s.GetSubLists(tt.args.ctx, tt.args.queryParam)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetSubLists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetSubLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_DeleteList(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
		Cfg        config.Config
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		fields  fields
		doMock  func(repo repository.MockRepositoryInterface)
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListByID(ctx, gomock.Any()).Return(&mockList, nil)
				mockrepo.EXPECT().MarkAsDeletedTx(ctx, gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "error when deleted",
			args: args{
				ctx: ctx,
				id:  1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListByID(ctx, gomock.Any()).Return(&mockList, nil)
				mockrepo.EXPECT().MarkAsDeletedTx(ctx, gomock.Any()).Return(sql.ErrNoRows)
			},
			wantErr: true,
		},
		{
			name: "error id not found",
			args: args{
				ctx: ctx,
				id:  1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListByID(ctx, gomock.Any()).Return(nil, sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repository: tt.fields.Repository,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockRepo := repository.NewMockRepositoryInterface(ctrl)

			s.Repository = mockRepo
			tt.doMock(*mockRepo)
			if err := s.DeleteList(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_DeleteSubList(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
		Cfg        config.Config
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name    string
		fields  fields
		doMock  func(repo repository.MockRepositoryInterface)
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				id:  1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetSubListByID(ctx, gomock.Any()).Return(&mockSubList, nil)
				mockrepo.EXPECT().MarkAsDeletedTx(ctx, gomock.Any()).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "error when deleted",
			args: args{
				ctx: ctx,
				id:  1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetSubListByID(ctx, gomock.Any()).Return(&mockSubList, nil)
				mockrepo.EXPECT().MarkAsDeletedTx(ctx, gomock.Any()).Return(sql.ErrNoRows)
			},
			wantErr: true,
		},
		{
			name: "error id not found",
			args: args{
				ctx: ctx,
				id:  1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetSubListByID(ctx, gomock.Any()).Return(nil, sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repository: tt.fields.Repository,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockRepo := repository.NewMockRepositoryInterface(ctrl)

			s.Repository = mockRepo
			tt.doMock(*mockRepo)
			if err := s.DeleteSubList(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteSubList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_EditList(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
		Cfg        config.Config
	}
	type args struct {
		ctx  context.Context
		list *models.AdditionalRequestEdit
		id   uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(repo repository.MockRepositoryInterface)
		want    *models.List
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequestEdit{
					ListRequestEdit: models.ListRequestEdit{
						Title:       "judul",
						Description: "description",
					},
					ParentID: 1,
				},
				id: 1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListByID(ctx, gomock.Any()).Return(&mockList, nil)
				mockrepo.EXPECT().UpdateTx(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(&mockList, nil)

			},
			want:    &mockList,
			wantErr: false,
		},
		{
			name: "error when get list",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequestEdit{
					ListRequestEdit: models.ListRequestEdit{
						Title:       "judul",
						Description: "description",
					},
					ParentID: 1,
				},
				id: 1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListByID(ctx, gomock.Any()).Return(nil, sql.ErrNoRows)

			},
			wantErr: true,
		},
		{
			name: "error when deleted",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequestEdit{
					ListRequestEdit: models.ListRequestEdit{
						Title:       "judul",
						Description: "description",
					},
					ParentID: 1,
				},
				id: 1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListByID(ctx, gomock.Any()).Return(&mockList, nil)
				mockrepo.EXPECT().UpdateTx(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(&mockList, sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repository: tt.fields.Repository,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockRepo := repository.NewMockRepositoryInterface(ctrl)

			s.Repository = mockRepo
			tt.doMock(*mockRepo)
			got, err := s.EditList(tt.args.ctx, tt.args.list, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.EditList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.EditList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_EditSubList(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
		Cfg        config.Config
	}
	type args struct {
		ctx  context.Context
		list *models.AdditionalRequestEdit
		id   uint
	}
	tests := []struct {
		name    string
		fields  fields
		doMock  func(repo repository.MockRepositoryInterface)
		args    args
		want    *models.List
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequestEdit{
					ListRequestEdit: models.ListRequestEdit{
						Title:       "judul",
						Description: "description",
					},
					ParentID: 1,
				},
				id: 1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetSubListByID(ctx, gomock.Any()).Return(&mockSubList, nil)
				mockrepo.EXPECT().UpdateTx(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(&mockList, nil)

			},
			want:    &mockList,
			wantErr: false,
		},
		{
			name: "error when get list",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequestEdit{
					ListRequestEdit: models.ListRequestEdit{
						Title:       "judul",
						Description: "description",
					},
					ParentID: 1,
				},
				id: 1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetSubListByID(ctx, gomock.Any()).Return(nil, sql.ErrNoRows)

			},
			wantErr: true,
		},
		{
			name: "error when deleted",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequestEdit{
					ListRequestEdit: models.ListRequestEdit{
						Title:       "judul",
						Description: "description",
					},
					ParentID: 1,
				},
				id: 1,
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetSubListByID(ctx, gomock.Any()).Return(&mockSubList, nil)
				mockrepo.EXPECT().UpdateTx(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(&mockList, sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repository: tt.fields.Repository,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockRepo := repository.NewMockRepositoryInterface(ctrl)

			s.Repository = mockRepo
			tt.doMock(*mockRepo)
			got, err := s.EditSubList(tt.args.ctx, tt.args.list, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.EditSubList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.EditSubList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateList(t *testing.T) {
	type fields struct {
		Repository repository.RepositoryInterface
		Cfg        config.Config
	}
	type args struct {
		ctx  context.Context
		list *models.AdditionalRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(repo repository.MockRepositoryInterface)
		want    *models.List
		wantErr bool
	}{
		{
			name: "success create list",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequest{
					ListRequest: models.ListRequest{
						Title:       "judul",
						Description: "description",
					},
				},
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().CreateTx(ctx, gomock.Any(), gomock.Any()).Return(&mockList, nil)

			},
			want:    &mockList,
			wantErr: false,
		},
		{
			name: "error when created list",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequest{
					ListRequest: models.ListRequest{
						Title:       "judul",
						Description: "description",
					},
				},
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().CreateTx(ctx, gomock.Any(), gomock.Any()).Return(&mockList, sql.ErrNoRows)
			},
			wantErr: true,
		},

		{
			name: "success create sub list",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequest{
					ListRequest: models.ListRequest{
						Title:       "judul",
						Description: "description",
					},
					ParentID: 1,
				},
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListByID(ctx, gomock.Any()).Return(&mockList, nil)
				mockrepo.EXPECT().CreateTx(ctx, gomock.Any(), gomock.Any()).Return(&mockList, nil)

			},
			want:    &mockList,
			wantErr: false,
		},
		{
			name: "success when get list",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequest{
					ListRequest: models.ListRequest{
						Title:       "judul",
						Description: "description",
					},
					ParentID: 1,
				},
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListByID(ctx, gomock.Any()).Return(nil, sql.ErrNoRows)

			},
			wantErr: true,
		},
		{
			name: "error when created sublist",
			args: args{
				ctx: ctx,
				list: &models.AdditionalRequest{
					ListRequest: models.ListRequest{
						Title:       "judul",
						Description: "description",
					},
					ParentID: 1,
				},
			},
			doMock: func(mockrepo repository.MockRepositoryInterface) {
				mockrepo.EXPECT().GetListByID(ctx, gomock.Any()).Return(&mockList, nil)
				mockrepo.EXPECT().CreateTx(ctx, gomock.Any(), gomock.Any()).Return(&mockList, sql.ErrNoRows)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repository: tt.fields.Repository,
				Cfg:        tt.fields.Cfg,
			}
			ctrl := gomock.NewController(t)
			mockRepo := repository.NewMockRepositoryInterface(ctrl)

			s.Repository = mockRepo
			tt.doMock(*mockRepo)
			got, err := s.CreateList(tt.args.ctx, tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.CreateList() = %v, want %v", got, tt.want)
			}
		})
	}
}
