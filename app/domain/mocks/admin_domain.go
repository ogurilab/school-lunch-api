// Code generated by MockGen. DO NOT EDIT.
// Source: domain/admin_domain.go
//
// Generated by this command:
//
//	mockgen -source domain/admin_domain.go -destination domain/mocks/admin_domain.go -package mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	echo "github.com/labstack/echo/v4"
	gomock "go.uber.org/mock/gomock"
)

// MockAdminController is a mock of AdminController interface.
type MockAdminController struct {
	ctrl     *gomock.Controller
	recorder *MockAdminControllerMockRecorder
}

// MockAdminControllerMockRecorder is the mock recorder for MockAdminController.
type MockAdminControllerMockRecorder struct {
	mock *MockAdminController
}

// NewMockAdminController creates a new mock instance.
func NewMockAdminController(ctrl *gomock.Controller) *MockAdminController {
	mock := &MockAdminController{ctrl: ctrl}
	mock.recorder = &MockAdminControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminController) EXPECT() *MockAdminControllerMockRecorder {
	return m.recorder
}

// CreateDish mocks base method.
func (m *MockAdminController) CreateDish(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDish", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDish indicates an expected call of CreateDish.
func (mr *MockAdminControllerMockRecorder) CreateDish(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDish", reflect.TypeOf((*MockAdminController)(nil).CreateDish), c)
}

// CreateDishes mocks base method.
func (m *MockAdminController) CreateDishes(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDishes", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDishes indicates an expected call of CreateDishes.
func (mr *MockAdminControllerMockRecorder) CreateDishes(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDishes", reflect.TypeOf((*MockAdminController)(nil).CreateDishes), c)
}

// CreateMenu mocks base method.
func (m *MockAdminController) CreateMenu(c echo.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMenu", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMenu indicates an expected call of CreateMenu.
func (mr *MockAdminControllerMockRecorder) CreateMenu(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMenu", reflect.TypeOf((*MockAdminController)(nil).CreateMenu), c)
}