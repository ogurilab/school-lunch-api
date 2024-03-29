// Code generated by MockGen. DO NOT EDIT.
// Source: domain/menu_with_dishes_domain.go
//
// Generated by this command:
//
//	mockgen -source domain/menu_with_dishes_domain.go -destination domain/mocks/menu_with_dishes_domain.go -package mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	echo "github.com/labstack/echo/v4"
	domain "github.com/ogurilab/school-lunch-api/domain"
	gomock "go.uber.org/mock/gomock"
)

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

// MockMenuWithDishesController is a mock of MenuWithDishesController interface.
type MockMenuWithDishesController struct {
	ctrl     *gomock.Controller
	recorder *MockMenuWithDishesControllerMockRecorder
}

// MockMenuWithDishesControllerMockRecorder is the mock recorder for MockMenuWithDishesController.
type MockMenuWithDishesControllerMockRecorder struct {
	mock *MockMenuWithDishesController
}

// NewMockMenuWithDishesController creates a new mock instance.
func NewMockMenuWithDishesController(ctrl *gomock.Controller) *MockMenuWithDishesController {
	mock := &MockMenuWithDishesController{ctrl: ctrl}
	mock.recorder = &MockMenuWithDishesControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMenuWithDishesController) EXPECT() *MockMenuWithDishesControllerMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockMenuWithDishesController) Fetch(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fetch indicates an expected call of Fetch.
func (mr *MockMenuWithDishesControllerMockRecorder) Fetch(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockMenuWithDishesController)(nil).Fetch), c)
}

// FetchByCity mocks base method.
func (m *MockMenuWithDishesController) FetchByCity(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByCity", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchByCity indicates an expected call of FetchByCity.
func (mr *MockMenuWithDishesControllerMockRecorder) FetchByCity(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByCity", reflect.TypeOf((*MockMenuWithDishesController)(nil).FetchByCity), c)
}

// GetByID mocks base method.
func (m *MockMenuWithDishesController) GetByID(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMenuWithDishesControllerMockRecorder) GetByID(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMenuWithDishesController)(nil).GetByID), c)
}
