package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ayyaa/todo-services/config"
	"github.com/ayyaa/todo-services/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepository_GetListPreloadByID(t *testing.T) {
	type fields struct {
		DB  *gorm.DB
		Cfg config.Config
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
		doMock  func(input int) (*gorm.DB, *sql.DB)
		isValue bool
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				ctx: ctx,
				id:  uint(1),
			},
			doMock: func(input int) (*gorm.DB, *sql.DB) {

				db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

				dialector := postgres.New(postgres.Config{
					DSN:                  "sqlmock_db_0",
					DriverName:           "postgres",
					Conn:                 db,
					PreferSimpleProtocol: true,
				})

				dbGorm, _ := gorm.Open(dialector, &gorm.Config{})
				query := `SELECT * FROM "list" WHERE parent_id IS NULL AND status = $1 AND "list"."id" = $2 ORDER BY "list"."id" LIMIT 1`
				mock.ExpectQuery(query).WillReturnError(errors.New("unexpected error"))
				return dbGorm, db

			},
			isValue: false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				DB:  tt.fields.DB,
				Cfg: tt.fields.Cfg,
			}
			DBPostgre, db := tt.doMock(int(tt.args.id))
			defer db.Close()

			r.DB = DBPostgre
			got, err := r.GetListPreloadByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetListPreloadByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) != tt.isValue {
				t.Errorf("Repository.GetListPreloadByID() = %v, want %v", got, tt.isValue)
			}
		})
	}
}

func TestRepository_GetLists(t *testing.T) {
	type fields struct {
		DB  *gorm.DB
		Cfg config.Config
	}
	type args struct {
		ctx    context.Context
		filter models.ParamRequest
	}
	ctx := context.Background()
	// mockList := models.List{
	// 	SubList: models.SubList{
	// 		ID:       1,
	// 		ParentID: nil,
	// 		Attachments: []*models.Attachment{
	// 			{ID: 1,
	// 				Filename: "active",
	// 				Filepath: "/upload/1234.pdf"},
	// 			{ID: 2,
	// 				Filename: "active",
	// 				Filepath: "/upload/1234.pdf"},
	// 		},
	// 	},
	// 	SubLists: []models.SubList{
	// 		{
	// 			ID:       1,
	// 			ParentID: nil,
	// 			Attachments: []*models.Attachment{
	// 				{ID: 1,
	// 					Filename: "active",
	// 					Filepath: "/upload/1234.pdf"},
	// 				{ID: 2,
	// 					Filename: "active",
	// 					Filepath: "/upload/1234.pdf"},
	// 			},
	// 		},
	// 		{
	// 			ID:       1,
	// 			ParentID: nil,
	// 			Attachments: []*models.Attachment{
	// 				{ID: 1,
	// 					Filename: "active",
	// 					Filepath: "/upload/1234.pdf"},
	// 				{ID: 2,
	// 					Filename: "active",
	// 					Filepath: "/upload/1234.pdf"},
	// 			},
	// 		},
	// 	},
	// }
	tests := []struct {
		name    string
		fields  fields
		args    args
		doMock  func(filter models.ParamRequest) (*gorm.DB, *sql.DB)
		want    []models.List
		want1   *models.Pagination
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				ctx: ctx,
				filter: models.ParamRequest{
					Page:    1,
					Size:    10,
					OrderBy: "title",
					Preload: false,
					Keyword: "ini",
				},
			},
			doMock: func(filter models.ParamRequest) (*gorm.DB, *sql.DB) {

				db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

				dialector := postgres.New(postgres.Config{
					DSN:                  "sqlmock_db_0",
					DriverName:           "postgres",
					Conn:                 db,
					PreferSimpleProtocol: true,
				})

				dbGorm, _ := gorm.Open(dialector, &gorm.Config{})
				query := `SELECT * FROM "list" WHERE parent_id IS NULL AND status = $1 AND "list"."id" = $2 ORDER BY "list"."id" LIMIT 1`
				mock.ExpectQuery(query).WillReturnError(errors.New("unexpected error"))
				return dbGorm, db

			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				DB:  tt.fields.DB,
				Cfg: tt.fields.Cfg,
			}
			DBPostgre, db := tt.doMock(tt.args.filter)
			defer db.Close()

			r.DB = DBPostgre
			got, got1, err := r.GetLists(tt.args.ctx, tt.args.filter)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)

		})
	}
}

func TestRepository_GetListByID(t *testing.T) {
	type fields struct {
		DB  *gorm.DB
		Cfg config.Config
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
		doMock  func(id int) (*gorm.DB, *sql.DB)
		want    *models.List
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				ctx: ctx,
				id:  uint(1),
			},
			doMock: func(input int) (*gorm.DB, *sql.DB) {

				db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

				dialector := postgres.New(postgres.Config{
					DSN:                  "sqlmock_db_0",
					DriverName:           "postgres",
					Conn:                 db,
					PreferSimpleProtocol: true,
				})

				dbGorm, _ := gorm.Open(dialector, &gorm.Config{})
				query := `SELECT * FROM "list" WHERE parent_id IS NULL AND status = $1 AND "list"."id" = $2 ORDER BY "list"."id" LIMIT 1`
				mock.ExpectQuery(query).WillReturnError(errors.New("unexpected error"))
				return dbGorm, db

			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				DB:  tt.fields.DB,
				Cfg: tt.fields.Cfg,
			}
			DBPostgre, db := tt.doMock(int(tt.args.id))
			defer db.Close()

			r.DB = DBPostgre
			got, err := r.GetListByID(tt.args.ctx, tt.args.id)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
