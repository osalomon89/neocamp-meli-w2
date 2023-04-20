// Code generated by MockGen. DO NOT EDIT.
// Source: ./item_usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockItemUsecase is a mock of ItemUsecase interface
type MockItemUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockItemUsecaseMockRecorder
}

// MockItemUsecaseMockRecorder is the mock recorder for MockItemUsecase
type MockItemUsecaseMockRecorder struct {
	mock *MockItemUsecase
}

// NewMockItemUsecase creates a new mock instance
func NewMockItemUsecase(ctrl *gomock.Controller) *MockItemUsecase {
	mock := &MockItemUsecase{ctrl: ctrl}
	mock.recorder = &MockItemUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockItemUsecase) EXPECT() *MockItemUsecaseMockRecorder {
	return m.recorder
}

// CreateItem mocks base method
func (m *MockItemUsecase) CreateItem(name string, stock int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", name, stock)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateItem indicates an expected call of CreateItem
func (mr *MockItemUsecaseMockRecorder) CreateItem(name, stock interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockItemUsecase)(nil).CreateItem), name, stock)
}

// GetItemByID mocks base method
func (m *MockItemUsecase) GetItemByID(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItemByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetItemByID indicates an expected call of GetItemByID
func (mr *MockItemUsecaseMockRecorder) GetItemByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItemByID", reflect.TypeOf((*MockItemUsecase)(nil).GetItemByID), id)
}
