// Code generated by MockGen. DO NOT EDIT.
// Source: domain/menu_domain.go
//
// Generated by this command:
//
//	mockgen -source domain/menu_domain.go -destination domain/mocks/menu_domain.go -package mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	domain "github.com/ogurilab/school-lunch-api/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockMenuRepository is a mock of MenuRepository interface.
type MockMenuRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMenuRepositoryMockRecorder
}

// MockMenuRepositoryMockRecorder is the mock recorder for MockMenuRepository.
type MockMenuRepositoryMockRecorder struct {
	mock *MockMenuRepository
}

// NewMockMenuRepository creates a new mock instance.
func NewMockMenuRepository(ctrl *gomock.Controller) *MockMenuRepository {
	mock := &MockMenuRepository{ctrl: ctrl}
	mock.recorder = &MockMenuRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMenuRepository) EXPECT() *MockMenuRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMenuRepository) Create(ctx context.Context, menu *domain.Menu) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, menu)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMenuRepositoryMockRecorder) Create(ctx, menu any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMenuRepository)(nil).Create), ctx, menu)
}

// FetchByCity mocks base method.
func (m *MockMenuRepository) FetchByCity(ctx context.Context, limit, offset int32, offered time.Time, city int32) ([]*domain.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByCity", ctx, limit, offset, offered, city)
	ret0, _ := ret[0].([]*domain.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchByCity indicates an expected call of FetchByCity.
func (mr *MockMenuRepositoryMockRecorder) FetchByCity(ctx, limit, offset, offered, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByCity", reflect.TypeOf((*MockMenuRepository)(nil).FetchByCity), ctx, limit, offset, offered, city)
}

// GetByID mocks base method.
func (m *MockMenuRepository) GetByID(ctx context.Context, id string, city int32) (*domain.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id, city)
	ret0, _ := ret[0].(*domain.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMenuRepositoryMockRecorder) GetByID(ctx, id, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMenuRepository)(nil).GetByID), ctx, id, city)
}

// MockMenuUsecase is a mock of MenuUsecase interface.
type MockMenuUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockMenuUsecaseMockRecorder
}

// MockMenuUsecaseMockRecorder is the mock recorder for MockMenuUsecase.
type MockMenuUsecaseMockRecorder struct {
	mock *MockMenuUsecase
}

// NewMockMenuUsecase creates a new mock instance.
func NewMockMenuUsecase(ctrl *gomock.Controller) *MockMenuUsecase {
	mock := &MockMenuUsecase{ctrl: ctrl}
	mock.recorder = &MockMenuUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMenuUsecase) EXPECT() *MockMenuUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMenuUsecase) Create(ctx context.Context, menu *domain.Menu) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, menu)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMenuUsecaseMockRecorder) Create(ctx, menu any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMenuUsecase)(nil).Create), ctx, menu)
}

// FetchByCity mocks base method.
func (m *MockMenuUsecase) FetchByCity(ctx context.Context, limit, offset int32, offered time.Time, city int32) ([]*domain.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByCity", ctx, limit, offset, offered, city)
	ret0, _ := ret[0].([]*domain.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchByCity indicates an expected call of FetchByCity.
func (mr *MockMenuUsecaseMockRecorder) FetchByCity(ctx, limit, offset, offered, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByCity", reflect.TypeOf((*MockMenuUsecase)(nil).FetchByCity), ctx, limit, offset, offered, city)
}

// GetByID mocks base method.
func (m *MockMenuUsecase) GetByID(ctx context.Context, id string, city int32) (*domain.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id, city)
	ret0, _ := ret[0].(*domain.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMenuUsecaseMockRecorder) GetByID(ctx, id, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMenuUsecase)(nil).GetByID), ctx, id, city)
}

// MockMenuWithDishesRepository is a mock of MenuWithDishesRepository interface.
type MockMenuWithDishesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMenuWithDishesRepositoryMockRecorder
}

// MockMenuWithDishesRepositoryMockRecorder is the mock recorder for MockMenuWithDishesRepository.
type MockMenuWithDishesRepositoryMockRecorder struct {
	mock *MockMenuWithDishesRepository
}

// NewMockMenuWithDishesRepository creates a new mock instance.
func NewMockMenuWithDishesRepository(ctrl *gomock.Controller) *MockMenuWithDishesRepository {
	mock := &MockMenuWithDishesRepository{ctrl: ctrl}
	mock.recorder = &MockMenuWithDishesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMenuWithDishesRepository) EXPECT() *MockMenuWithDishesRepositoryMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockMenuWithDishesRepository) Fetch(ctx context.Context, limit, offset int32, offered time.Time) ([]*domain.MenuWithDishes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, limit, offset, offered)
	ret0, _ := ret[0].([]*domain.MenuWithDishes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockMenuWithDishesRepositoryMockRecorder) Fetch(ctx, limit, offset, offered any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockMenuWithDishesRepository)(nil).Fetch), ctx, limit, offset, offered)
}

// FetchByCity mocks base method.
func (m *MockMenuWithDishesRepository) FetchByCity(ctx context.Context, limit, offset int32, offered time.Time, city int32) ([]*domain.MenuWithDishes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByCity", ctx, limit, offset, offered, city)
	ret0, _ := ret[0].([]*domain.MenuWithDishes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchByCity indicates an expected call of FetchByCity.
func (mr *MockMenuWithDishesRepositoryMockRecorder) FetchByCity(ctx, limit, offset, offered, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByCity", reflect.TypeOf((*MockMenuWithDishesRepository)(nil).FetchByCity), ctx, limit, offset, offered, city)
}

// GetByID mocks base method.
func (m *MockMenuWithDishesRepository) GetByID(ctx context.Context, id string, city int32) (*domain.MenuWithDishes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id, city)
	ret0, _ := ret[0].(*domain.MenuWithDishes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMenuWithDishesRepositoryMockRecorder) GetByID(ctx, id, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMenuWithDishesRepository)(nil).GetByID), ctx, id, city)
}

// MockMenuWithDishesUsecase is a mock of MenuWithDishesUsecase interface.
type MockMenuWithDishesUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockMenuWithDishesUsecaseMockRecorder
}

// MockMenuWithDishesUsecaseMockRecorder is the mock recorder for MockMenuWithDishesUsecase.
type MockMenuWithDishesUsecaseMockRecorder struct {
	mock *MockMenuWithDishesUsecase
}

// NewMockMenuWithDishesUsecase creates a new mock instance.
func NewMockMenuWithDishesUsecase(ctrl *gomock.Controller) *MockMenuWithDishesUsecase {
	mock := &MockMenuWithDishesUsecase{ctrl: ctrl}
	mock.recorder = &MockMenuWithDishesUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMenuWithDishesUsecase) EXPECT() *MockMenuWithDishesUsecaseMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockMenuWithDishesUsecase) Fetch(ctx context.Context, limit, offset int32, offered time.Time) ([]*domain.MenuWithDishes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, limit, offset, offered)
	ret0, _ := ret[0].([]*domain.MenuWithDishes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockMenuWithDishesUsecaseMockRecorder) Fetch(ctx, limit, offset, offered any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockMenuWithDishesUsecase)(nil).Fetch), ctx, limit, offset, offered)
}

// FetchByCity mocks base method.
func (m *MockMenuWithDishesUsecase) FetchByCity(ctx context.Context, limit, offset int32, offered time.Time, city int32) ([]*domain.MenuWithDishes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByCity", ctx, limit, offset, offered, city)
	ret0, _ := ret[0].([]*domain.MenuWithDishes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchByCity indicates an expected call of FetchByCity.
func (mr *MockMenuWithDishesUsecaseMockRecorder) FetchByCity(ctx, limit, offset, offered, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByCity", reflect.TypeOf((*MockMenuWithDishesUsecase)(nil).FetchByCity), ctx, limit, offset, offered, city)
}

// GetByID mocks base method.
func (m *MockMenuWithDishesUsecase) GetByID(ctx context.Context, id string, city int32) (*domain.MenuWithDishes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id, city)
	ret0, _ := ret[0].(*domain.MenuWithDishes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMenuWithDishesUsecaseMockRecorder) GetByID(ctx, id, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMenuWithDishesUsecase)(nil).GetByID), ctx, id, city)
}
