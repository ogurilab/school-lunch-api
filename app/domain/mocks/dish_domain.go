// Code generated by MockGen. DO NOT EDIT.
// Source: domain/dish_domain.go
//
// Generated by this command:
//
//	mockgen -source domain/dish_domain.go -destination domain/mocks/dish_domain.go -package mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	domain "github.com/ogurilab/school-lunch-api/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockDishRepository is a mock of DishRepository interface.
type MockDishRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDishRepositoryMockRecorder
}

// MockDishRepositoryMockRecorder is the mock recorder for MockDishRepository.
type MockDishRepositoryMockRecorder struct {
	mock *MockDishRepository
}

// NewMockDishRepository creates a new mock instance.
func NewMockDishRepository(ctrl *gomock.Controller) *MockDishRepository {
	mock := &MockDishRepository{ctrl: ctrl}
	mock.recorder = &MockDishRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDishRepository) EXPECT() *MockDishRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDishRepository) Create(ctx context.Context, dish *domain.Dish, menuID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, dish, menuID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDishRepositoryMockRecorder) Create(ctx, dish, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDishRepository)(nil).Create), ctx, dish, menuID)
}

// Fetch mocks base method.
func (m *MockDishRepository) Fetch(ctx context.Context, limit, offset int32) ([]*domain.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, limit, offset)
	ret0, _ := ret[0].([]*domain.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockDishRepositoryMockRecorder) Fetch(ctx, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockDishRepository)(nil).Fetch), ctx, limit, offset)
}

// FetchByMenuID mocks base method.
func (m *MockDishRepository) FetchByMenuID(ctx context.Context, menuID string) ([]*domain.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByMenuID", ctx, menuID)
	ret0, _ := ret[0].([]*domain.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchByMenuID indicates an expected call of FetchByMenuID.
func (mr *MockDishRepositoryMockRecorder) FetchByMenuID(ctx, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByMenuID", reflect.TypeOf((*MockDishRepository)(nil).FetchByMenuID), ctx, menuID)
}

// FetchByName mocks base method.
func (m *MockDishRepository) FetchByName(ctx context.Context, search string, limit, offset int32) ([]*domain.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByName", ctx, search, limit, offset)
	ret0, _ := ret[0].([]*domain.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchByName indicates an expected call of FetchByName.
func (mr *MockDishRepositoryMockRecorder) FetchByName(ctx, search, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByName", reflect.TypeOf((*MockDishRepository)(nil).FetchByName), ctx, search, limit, offset)
}

// GetByID mocks base method.
func (m *MockDishRepository) GetByID(ctx context.Context, id string, limit, offset int32) (*domain.DishWithMenuIDs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id, limit, offset)
	ret0, _ := ret[0].(*domain.DishWithMenuIDs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockDishRepositoryMockRecorder) GetByID(ctx, id, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockDishRepository)(nil).GetByID), ctx, id, limit, offset)
}

// GetByIdInCity mocks base method.
func (m *MockDishRepository) GetByIdInCity(ctx context.Context, id string, limit, offset, city int32) (*domain.DishWithMenuIDs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIdInCity", ctx, id, limit, offset, city)
	ret0, _ := ret[0].(*domain.DishWithMenuIDs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIdInCity indicates an expected call of GetByIdInCity.
func (mr *MockDishRepositoryMockRecorder) GetByIdInCity(ctx, id, limit, offset, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdInCity", reflect.TypeOf((*MockDishRepository)(nil).GetByIdInCity), ctx, id, limit, offset, city)
}

// MockDishUsecase is a mock of DishUsecase interface.
type MockDishUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockDishUsecaseMockRecorder
}

// MockDishUsecaseMockRecorder is the mock recorder for MockDishUsecase.
type MockDishUsecaseMockRecorder struct {
	mock *MockDishUsecase
}

// NewMockDishUsecase creates a new mock instance.
func NewMockDishUsecase(ctrl *gomock.Controller) *MockDishUsecase {
	mock := &MockDishUsecase{ctrl: ctrl}
	mock.recorder = &MockDishUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDishUsecase) EXPECT() *MockDishUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockDishUsecase) Create(ctx context.Context, dish *domain.Dish, menuID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, dish, menuID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockDishUsecaseMockRecorder) Create(ctx, dish, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDishUsecase)(nil).Create), ctx, dish, menuID)
}

// Fetch mocks base method.
func (m *MockDishUsecase) Fetch(ctx context.Context, search string, limit, offset int32) ([]*domain.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, search, limit, offset)
	ret0, _ := ret[0].([]*domain.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockDishUsecaseMockRecorder) Fetch(ctx, search, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockDishUsecase)(nil).Fetch), ctx, search, limit, offset)
}

// FetchByMenuID mocks base method.
func (m *MockDishUsecase) FetchByMenuID(ctx context.Context, menuID string) ([]*domain.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByMenuID", ctx, menuID)
	ret0, _ := ret[0].([]*domain.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchByMenuID indicates an expected call of FetchByMenuID.
func (mr *MockDishUsecaseMockRecorder) FetchByMenuID(ctx, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByMenuID", reflect.TypeOf((*MockDishUsecase)(nil).FetchByMenuID), ctx, menuID)
}

// GetByID mocks base method.
func (m *MockDishUsecase) GetByID(ctx context.Context, id string, limit, offset int32) (*domain.DishWithMenuIDs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id, limit, offset)
	ret0, _ := ret[0].(*domain.DishWithMenuIDs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockDishUsecaseMockRecorder) GetByID(ctx, id, limit, offset any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockDishUsecase)(nil).GetByID), ctx, id, limit, offset)
}

// GetByIdInCity mocks base method.
func (m *MockDishUsecase) GetByIdInCity(ctx context.Context, id string, limit, offset, city int32) (*domain.DishWithMenuIDs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIdInCity", ctx, id, limit, offset, city)
	ret0, _ := ret[0].(*domain.DishWithMenuIDs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIdInCity indicates an expected call of GetByIdInCity.
func (mr *MockDishUsecaseMockRecorder) GetByIdInCity(ctx, id, limit, offset, city any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdInCity", reflect.TypeOf((*MockDishUsecase)(nil).GetByIdInCity), ctx, id, limit, offset, city)
}

// MockDishController is a mock of DishController interface.
type MockDishController struct {
	ctrl     *gomock.Controller
	recorder *MockDishControllerMockRecorder
}

// MockDishControllerMockRecorder is the mock recorder for MockDishController.
type MockDishControllerMockRecorder struct {
	mock *MockDishController
}

// NewMockDishController creates a new mock instance.
func NewMockDishController(ctrl *gomock.Controller) *MockDishController {
	mock := &MockDishController{ctrl: ctrl}
	mock.recorder = &MockDishControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDishController) EXPECT() *MockDishControllerMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockDishController) Fetch(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Fetch indicates an expected call of Fetch.
func (mr *MockDishControllerMockRecorder) Fetch(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockDishController)(nil).Fetch), c)
}

// FetchByMenuID mocks base method.
func (m *MockDishController) FetchByMenuID(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByMenuID", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchByMenuID indicates an expected call of FetchByMenuID.
func (mr *MockDishControllerMockRecorder) FetchByMenuID(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByMenuID", reflect.TypeOf((*MockDishController)(nil).FetchByMenuID), c)
}

// GetByID mocks base method.
func (m *MockDishController) GetByID(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetByID indicates an expected call of GetByID.
func (mr *MockDishControllerMockRecorder) GetByID(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockDishController)(nil).GetByID), c)
}

// GetByIdInCity mocks base method.
func (m *MockDishController) GetByIdInCity(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIdInCity", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetByIdInCity indicates an expected call of GetByIdInCity.
func (mr *MockDishControllerMockRecorder) GetByIdInCity(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIdInCity", reflect.TypeOf((*MockDishController)(nil).GetByIdInCity), c)
}
