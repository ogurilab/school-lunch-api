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

	domain "github.com/ogurilab/school-lunch-api/domain"
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

// CreateCity mocks base method.
func (m *MockQuery) CreateCity(ctx context.Context, arg db.CreateCityParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCity", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCity indicates an expected call of CreateCity.
func (mr *MockQueryMockRecorder) CreateCity(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCity", reflect.TypeOf((*MockQuery)(nil).CreateCity), ctx, arg)
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

// CreateDishTx mocks base method.
func (m *MockQuery) CreateDishTx(ctx context.Context, dish *domain.Dish, menuID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDishTx", ctx, dish, menuID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDishTx indicates an expected call of CreateDishTx.
func (mr *MockQueryMockRecorder) CreateDishTx(ctx, dish, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDishTx", reflect.TypeOf((*MockQuery)(nil).CreateDishTx), ctx, dish, menuID)
}

// CreateDishesTx mocks base method.
func (m *MockQuery) CreateDishesTx(ctx context.Context, dishes []*domain.Dish, menuID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDishesTx", ctx, dishes, menuID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDishesTx indicates an expected call of CreateDishesTx.
func (mr *MockQueryMockRecorder) CreateDishesTx(ctx, dishes, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDishesTx", reflect.TypeOf((*MockQuery)(nil).CreateDishesTx), ctx, dishes, menuID)
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

// CreateMenuDish mocks base method.
func (m *MockQuery) CreateMenuDish(ctx context.Context, arg db.CreateMenuDishParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMenuDish", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMenuDish indicates an expected call of CreateMenuDish.
func (mr *MockQueryMockRecorder) CreateMenuDish(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMenuDish", reflect.TypeOf((*MockQuery)(nil).CreateMenuDish), ctx, arg)
}

// GetCity mocks base method.
func (m *MockQuery) GetCity(ctx context.Context, cityCode int32) (db.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCity", ctx, cityCode)
	ret0, _ := ret[0].(db.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCity indicates an expected call of GetCity.
func (mr *MockQueryMockRecorder) GetCity(ctx, cityCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCity", reflect.TypeOf((*MockQuery)(nil).GetCity), ctx, cityCode)
}

// GetDish mocks base method.
func (m *MockQuery) GetDish(ctx context.Context, arg db.GetDishParams) ([]db.GetDishRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDish", ctx, arg)
	ret0, _ := ret[0].([]db.GetDishRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDish indicates an expected call of GetDish.
func (mr *MockQueryMockRecorder) GetDish(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDish", reflect.TypeOf((*MockQuery)(nil).GetDish), ctx, arg)
}

// GetDishInCity mocks base method.
func (m *MockQuery) GetDishInCity(ctx context.Context, arg db.GetDishInCityParams) ([]db.GetDishInCityRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDishInCity", ctx, arg)
	ret0, _ := ret[0].([]db.GetDishInCityRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDishInCity indicates an expected call of GetDishInCity.
func (mr *MockQueryMockRecorder) GetDishInCity(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDishInCity", reflect.TypeOf((*MockQuery)(nil).GetDishInCity), ctx, arg)
}

// GetMenu mocks base method.
func (m *MockQuery) GetMenu(ctx context.Context, arg db.GetMenuParams) (db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenu", ctx, arg)
	ret0, _ := ret[0].(db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenu indicates an expected call of GetMenu.
func (mr *MockQueryMockRecorder) GetMenu(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenu", reflect.TypeOf((*MockQuery)(nil).GetMenu), ctx, arg)
}

// GetMenuWithDishes mocks base method.
func (m *MockQuery) GetMenuWithDishes(ctx context.Context, arg db.GetMenuWithDishesParams) ([]db.GetMenuWithDishesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMenuWithDishes", ctx, arg)
	ret0, _ := ret[0].([]db.GetMenuWithDishesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMenuWithDishes indicates an expected call of GetMenuWithDishes.
func (mr *MockQueryMockRecorder) GetMenuWithDishes(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMenuWithDishes", reflect.TypeOf((*MockQuery)(nil).GetMenuWithDishes), ctx, arg)
}

// ListCities mocks base method.
func (m *MockQuery) ListCities(ctx context.Context, arg db.ListCitiesParams) ([]db.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCities", ctx, arg)
	ret0, _ := ret[0].([]db.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCities indicates an expected call of ListCities.
func (mr *MockQueryMockRecorder) ListCities(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCities", reflect.TypeOf((*MockQuery)(nil).ListCities), ctx, arg)
}

// ListCitiesByName mocks base method.
func (m *MockQuery) ListCitiesByName(ctx context.Context, arg db.ListCitiesByNameParams) ([]db.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCitiesByName", ctx, arg)
	ret0, _ := ret[0].([]db.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCitiesByName indicates an expected call of ListCitiesByName.
func (mr *MockQueryMockRecorder) ListCitiesByName(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCitiesByName", reflect.TypeOf((*MockQuery)(nil).ListCitiesByName), ctx, arg)
}

// ListCitiesByPrefecture mocks base method.
func (m *MockQuery) ListCitiesByPrefecture(ctx context.Context, arg db.ListCitiesByPrefectureParams) ([]db.City, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCitiesByPrefecture", ctx, arg)
	ret0, _ := ret[0].([]db.City)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCitiesByPrefecture indicates an expected call of ListCitiesByPrefecture.
func (mr *MockQueryMockRecorder) ListCitiesByPrefecture(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCitiesByPrefecture", reflect.TypeOf((*MockQuery)(nil).ListCitiesByPrefecture), ctx, arg)
}

// ListDish mocks base method.
func (m *MockQuery) ListDish(ctx context.Context, arg db.ListDishParams) ([]db.ListDishRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDish", ctx, arg)
	ret0, _ := ret[0].([]db.ListDishRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDish indicates an expected call of ListDish.
func (mr *MockQueryMockRecorder) ListDish(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDish", reflect.TypeOf((*MockQuery)(nil).ListDish), ctx, arg)
}

// ListDishByMenuID mocks base method.
func (m *MockQuery) ListDishByMenuID(ctx context.Context, menuID string) ([]db.ListDishByMenuIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDishByMenuID", ctx, menuID)
	ret0, _ := ret[0].([]db.ListDishByMenuIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDishByMenuID indicates an expected call of ListDishByMenuID.
func (mr *MockQueryMockRecorder) ListDishByMenuID(ctx, menuID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDishByMenuID", reflect.TypeOf((*MockQuery)(nil).ListDishByMenuID), ctx, menuID)
}

// ListDishByName mocks base method.
func (m *MockQuery) ListDishByName(ctx context.Context, arg db.ListDishByNameParams) ([]db.ListDishByNameRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDishByName", ctx, arg)
	ret0, _ := ret[0].([]db.ListDishByNameRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDishByName indicates an expected call of ListDishByName.
func (mr *MockQueryMockRecorder) ListDishByName(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDishByName", reflect.TypeOf((*MockQuery)(nil).ListDishByName), ctx, arg)
}

// ListMenu mocks base method.
func (m *MockQuery) ListMenu(ctx context.Context, arg db.ListMenuParams) ([]db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMenu", ctx, arg)
	ret0, _ := ret[0].([]db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMenu indicates an expected call of ListMenu.
func (mr *MockQueryMockRecorder) ListMenu(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMenu", reflect.TypeOf((*MockQuery)(nil).ListMenu), ctx, arg)
}

// ListMenuByCity mocks base method.
func (m *MockQuery) ListMenuByCity(ctx context.Context, arg db.ListMenuByCityParams) ([]db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMenuByCity", ctx, arg)
	ret0, _ := ret[0].([]db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMenuByCity indicates an expected call of ListMenuByCity.
func (mr *MockQueryMockRecorder) ListMenuByCity(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMenuByCity", reflect.TypeOf((*MockQuery)(nil).ListMenuByCity), ctx, arg)
}

// ListMenuInIds mocks base method.
func (m *MockQuery) ListMenuInIds(ctx context.Context, arg db.ListMenuInIdsParams) ([]db.Menu, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMenuInIds", ctx, arg)
	ret0, _ := ret[0].([]db.Menu)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMenuInIds indicates an expected call of ListMenuInIds.
func (mr *MockQueryMockRecorder) ListMenuInIds(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMenuInIds", reflect.TypeOf((*MockQuery)(nil).ListMenuInIds), ctx, arg)
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

// ListMenuWithDishesByCity mocks base method.
func (m *MockQuery) ListMenuWithDishesByCity(ctx context.Context, arg db.ListMenuWithDishesByCityParams) ([]db.ListMenuWithDishesByCityRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMenuWithDishesByCity", ctx, arg)
	ret0, _ := ret[0].([]db.ListMenuWithDishesByCityRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMenuWithDishesByCity indicates an expected call of ListMenuWithDishesByCity.
func (mr *MockQueryMockRecorder) ListMenuWithDishesByCity(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMenuWithDishesByCity", reflect.TypeOf((*MockQuery)(nil).ListMenuWithDishesByCity), ctx, arg)
}

// UpdateAvailable mocks base method.
func (m *MockQuery) UpdateAvailable(ctx context.Context, cityCode int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvailable", ctx, cityCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvailable indicates an expected call of UpdateAvailable.
func (mr *MockQueryMockRecorder) UpdateAvailable(ctx, cityCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvailable", reflect.TypeOf((*MockQuery)(nil).UpdateAvailable), ctx, cityCode)
}
