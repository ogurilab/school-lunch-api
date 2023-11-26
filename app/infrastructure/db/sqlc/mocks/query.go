// Code generated by MockGen. DO NOT EDIT.
// Source: infrastructure/db/sqlc/query.go
//
// Generated by this command:
//
//	mockgen -source infrastructure/db/sqlc/query.go -destination infrastructure/db/sqlc/mocks/query.go -package mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	db "github.com/ogurilab/school-lunch-api/infrastructure/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockQuery is a mock of Query interface.
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery.
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance.
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// CreateDish mocks base method.
func (m *MockQuery) CreateDish(ctx context.Context, arg db.CreateDishParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDish", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDish indicates an expected call of CreateDish.
func (mr *MockQueryMockRecorder) CreateDish(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDish", reflect.TypeOf((*MockQuery)(nil).CreateDish), ctx, arg)
}

// CreateMenu mocks base method.
func (m *MockQuery) CreateMenu(ctx context.Context, arg db.CreateMenuParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMenu", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMenu indicates an expected call of CreateMenu.
func (mr *MockQueryMockRecorder) CreateMenu(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMenu", reflect.TypeOf((*MockQuery)(nil).CreateMenu), ctx, arg)
}

// GetDish mocks base method.
func (m *MockQuery) GetDish(ctx context.Context, id string) (db.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDish", ctx, id)
	ret0, _ := ret[0].(db.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDish indicates an expected call of GetDish.
func (mr *MockQueryMockRecorder) GetDish(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDish", reflect.TypeOf((*MockQuery)(nil).GetDish), ctx, id)
}

// GetDishByNames mocks base method.
func (m *MockQuery) GetDishByNames(ctx context.Context, arg db.GetDishByNamesParams) ([]db.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDishByNames", ctx, arg)
	ret0, _ := ret[0].([]db.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDishByNames indicates an expected call of GetDishByNames.
func (mr *MockQueryMockRecorder) GetDishByNames(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDishByNames", reflect.TypeOf((*MockQuery)(nil).GetDishByNames), ctx, arg)
}

// GetMenu mocks base method.
func (m *MockQuery) GetMenu(ctx context.Context, id string) (db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenu", ctx, id)
	ret0, _ := ret[0].(db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenu indicates an expected call of GetMenu.
func (mr *MockQueryMockRecorder) GetMenu(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenu", reflect.TypeOf((*MockQuery)(nil).GetMenu), ctx, id)
}

// GetMenuByOfferedAt mocks base method.
func (m *MockQuery) GetMenuByOfferedAt(ctx context.Context, offeredAt time.Time) (db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuByOfferedAt", ctx, offeredAt)
	ret0, _ := ret[0].(db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenuByOfferedAt indicates an expected call of GetMenuByOfferedAt.
func (mr *MockQueryMockRecorder) GetMenuByOfferedAt(ctx, offeredAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuByOfferedAt", reflect.TypeOf((*MockQuery)(nil).GetMenuByOfferedAt), ctx, offeredAt)
}

// GetMenuWithDishes mocks base method.
func (m *MockQuery) GetMenuWithDishes(ctx context.Context, id string) (db.GetMenuWithDishesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuWithDishes", ctx, id)
	ret0, _ := ret[0].(db.GetMenuWithDishesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenuWithDishes indicates an expected call of GetMenuWithDishes.
func (mr *MockQueryMockRecorder) GetMenuWithDishes(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuWithDishes", reflect.TypeOf((*MockQuery)(nil).GetMenuWithDishes), ctx, id)
}

// GetMenuWithDishesByOfferedAt mocks base method.
func (m *MockQuery) GetMenuWithDishesByOfferedAt(ctx context.Context, offeredAt time.Time) (db.GetMenuWithDishesByOfferedAtRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuWithDishesByOfferedAt", ctx, offeredAt)
	ret0, _ := ret[0].(db.GetMenuWithDishesByOfferedAtRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenuWithDishesByOfferedAt indicates an expected call of GetMenuWithDishesByOfferedAt.
func (mr *MockQueryMockRecorder) GetMenuWithDishesByOfferedAt(ctx, offeredAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuWithDishesByOfferedAt", reflect.TypeOf((*MockQuery)(nil).GetMenuWithDishesByOfferedAt), ctx, offeredAt)
}

// ListDishes mocks base method.
func (m *MockQuery) ListDishes(ctx context.Context, menuID string) ([]db.Dish, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDishes", ctx, menuID)
	ret0, _ := ret[0].([]db.Dish)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDishes indicates an expected call of ListDishes.
func (mr *MockQueryMockRecorder) ListDishes(ctx, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDishes", reflect.TypeOf((*MockQuery)(nil).ListDishes), ctx, menuID)
}

// ListMenuWithDishes mocks base method.
func (m *MockQuery) ListMenuWithDishes(ctx context.Context, arg db.ListMenuWithDishesParams) ([]db.ListMenuWithDishesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMenuWithDishes", ctx, arg)
	ret0, _ := ret[0].([]db.ListMenuWithDishesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMenuWithDishes indicates an expected call of ListMenuWithDishes.
func (mr *MockQueryMockRecorder) ListMenuWithDishes(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMenuWithDishes", reflect.TypeOf((*MockQuery)(nil).ListMenuWithDishes), ctx, arg)
}

// ListMenuWithDishesByOfferedAt mocks base method.
func (m *MockQuery) ListMenuWithDishesByOfferedAt(ctx context.Context, arg db.ListMenuWithDishesByOfferedAtParams) ([]db.ListMenuWithDishesByOfferedAtRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMenuWithDishesByOfferedAt", ctx, arg)
	ret0, _ := ret[0].([]db.ListMenuWithDishesByOfferedAtRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMenuWithDishesByOfferedAt indicates an expected call of ListMenuWithDishesByOfferedAt.
func (mr *MockQueryMockRecorder) ListMenuWithDishesByOfferedAt(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMenuWithDishesByOfferedAt", reflect.TypeOf((*MockQuery)(nil).ListMenuWithDishesByOfferedAt), ctx, arg)
}

// ListMenus mocks base method.
func (m *MockQuery) ListMenus(ctx context.Context, arg db.ListMenusParams) ([]db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMenus", ctx, arg)
	ret0, _ := ret[0].([]db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMenus indicates an expected call of ListMenus.
func (mr *MockQueryMockRecorder) ListMenus(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMenus", reflect.TypeOf((*MockQuery)(nil).ListMenus), ctx, arg)
}

// ListMenusByOfferedAt mocks base method.
func (m *MockQuery) ListMenusByOfferedAt(ctx context.Context, arg db.ListMenusByOfferedAtParams) ([]db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMenusByOfferedAt", ctx, arg)
	ret0, _ := ret[0].([]db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMenusByOfferedAt indicates an expected call of ListMenusByOfferedAt.
func (mr *MockQueryMockRecorder) ListMenusByOfferedAt(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMenusByOfferedAt", reflect.TypeOf((*MockQuery)(nil).ListMenusByOfferedAt), ctx, arg)
}
