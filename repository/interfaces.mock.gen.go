// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	models "github.com/ayyaa/todo-services/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateTx mocks base method.
func (m *MockRepositoryInterface) CreateTx(ctx context.Context, list *models.List, attachments []*models.Attachment) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTx", ctx, list, attachments)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockRepositoryInterfaceMockRecorder) CreateTx(ctx, list, attachments interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateTx), ctx, list, attachments)
}

// GetListByID mocks base method.
func (m *MockRepositoryInterface) GetListByID(ctx context.Context, id uint) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListByID", ctx, id)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListByID indicates an expected call of GetListByID.
func (mr *MockRepositoryInterfaceMockRecorder) GetListByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListByID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetListByID), ctx, id)
}

// GetListPreloadByID mocks base method.
func (m *MockRepositoryInterface) GetListPreloadByID(ctx context.Context, id uint) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListPreloadByID", ctx, id)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListPreloadByID indicates an expected call of GetListPreloadByID.
func (mr *MockRepositoryInterfaceMockRecorder) GetListPreloadByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListPreloadByID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetListPreloadByID), ctx, id)
}

// GetLists mocks base method.
func (m *MockRepositoryInterface) GetLists(ctx context.Context, filter models.ParamRequest) ([]models.List, *models.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLists", ctx, filter)
	ret0, _ := ret[0].([]models.List)
	ret1, _ := ret[1].(*models.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetLists indicates an expected call of GetLists.
func (mr *MockRepositoryInterfaceMockRecorder) GetLists(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLists", reflect.TypeOf((*MockRepositoryInterface)(nil).GetLists), ctx, filter)
}

// GetSubListByID mocks base method.
func (m *MockRepositoryInterface) GetSubListByID(ctx context.Context, id uint) (*models.SubList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubListByID", ctx, id)
	ret0, _ := ret[0].(*models.SubList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubListByID indicates an expected call of GetSubListByID.
func (mr *MockRepositoryInterfaceMockRecorder) GetSubListByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubListByID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetSubListByID), ctx, id)
}

// GetSubListPreloadByID mocks base method.
func (m *MockRepositoryInterface) GetSubListPreloadByID(ctx context.Context, id uint) (*models.SubList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubListPreloadByID", ctx, id)
	ret0, _ := ret[0].(*models.SubList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubListPreloadByID indicates an expected call of GetSubListPreloadByID.
func (mr *MockRepositoryInterfaceMockRecorder) GetSubListPreloadByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubListPreloadByID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetSubListPreloadByID), ctx, id)
}

// GetSubLists mocks base method.
func (m *MockRepositoryInterface) GetSubLists(ctx context.Context, filter models.ParamRequest) ([]models.List, *models.Pagination, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubLists", ctx, filter)
	ret0, _ := ret[0].([]models.List)
	ret1, _ := ret[1].(*models.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSubLists indicates an expected call of GetSubLists.
func (mr *MockRepositoryInterfaceMockRecorder) GetSubLists(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubLists", reflect.TypeOf((*MockRepositoryInterface)(nil).GetSubLists), ctx, filter)
}

// MarkAsDeletedTx mocks base method.
func (m *MockRepositoryInterface) MarkAsDeletedTx(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkAsDeletedTx", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkAsDeletedTx indicates an expected call of MarkAsDeletedTx.
func (mr *MockRepositoryInterfaceMockRecorder) MarkAsDeletedTx(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAsDeletedTx", reflect.TypeOf((*MockRepositoryInterface)(nil).MarkAsDeletedTx), ctx, id)
}

// UpdateTx mocks base method.
func (m *MockRepositoryInterface) UpdateTx(ctx context.Context, list *models.List, attachments []*models.Attachment, id uint) (*models.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTx", ctx, list, attachments, id)
	ret0, _ := ret[0].(*models.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTx indicates an expected call of UpdateTx.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateTx(ctx, list, attachments, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTx", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateTx), ctx, list, attachments, id)
}
