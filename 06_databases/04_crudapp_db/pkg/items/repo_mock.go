// Code generated by MockGen. DO NOT EDIT.
// Source: item.go

// Package items is a generated GoMock package.
package items

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockItemRepo is a mock of ItemRepo interface.
type MockItemRepo struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepoMockRecorder
}

// MockItemRepoMockRecorder is the mock recorder for MockItemRepo.
type MockItemRepoMockRecorder struct {
	mock *MockItemRepo
}

// NewMockItemRepo creates a new mock instance.
func NewMockItemRepo(ctrl *gomock.Controller) *MockItemRepo {
	mock := &MockItemRepo{ctrl: ctrl}
	mock.recorder = &MockItemRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepo) EXPECT() *MockItemRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockItemRepo) Add(arg0 *Item) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockItemRepoMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockItemRepo)(nil).Add), arg0)
}

// Delete mocks base method.
func (m *MockItemRepo) Delete(arg0 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockItemRepoMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockItemRepo)(nil).Delete), arg0)
}

// GetAll mocks base method.
func (m *MockItemRepo) GetAll() ([]*Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockItemRepoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockItemRepo)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockItemRepo) GetByID(arg0 int64) (*Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(*Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockItemRepoMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockItemRepo)(nil).GetByID), arg0)
}

// Update mocks base method.
func (m *MockItemRepo) Update(arg0 *Item) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockItemRepoMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockItemRepo)(nil).Update), arg0)
}
