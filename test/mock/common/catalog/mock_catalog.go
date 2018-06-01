// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/pkg/common/catalog (interfaces: Catalog)

// Package mock_catalog is a generated GoMock package.
package mock_catalog

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	catalog "github.com/spiffe/spire/pkg/common/catalog"
	reflect "reflect"
)

// MockCatalog is a mock of Catalog interface
type MockCatalog struct {
	ctrl     *gomock.Controller
	recorder *MockCatalogMockRecorder
}

// MockCatalogMockRecorder is the mock recorder for MockCatalog
type MockCatalogMockRecorder struct {
	mock *MockCatalog
}

// NewMockCatalog creates a new mock instance
func NewMockCatalog(ctrl *gomock.Controller) *MockCatalog {
	mock := &MockCatalog{ctrl: ctrl}
	mock.recorder = &MockCatalogMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCatalog) EXPECT() *MockCatalogMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockCatalog) Find(arg0 catalog.Plugin) *catalog.ManagedPlugin {
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(*catalog.ManagedPlugin)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockCatalogMockRecorder) Find(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCatalog)(nil).Find), arg0)
}

// Plugins mocks base method
func (m *MockCatalog) Plugins() []*catalog.ManagedPlugin {
	ret := m.ctrl.Call(m, "Plugins")
	ret0, _ := ret[0].([]*catalog.ManagedPlugin)
	return ret0
}

// Plugins indicates an expected call of Plugins
func (mr *MockCatalogMockRecorder) Plugins() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Plugins", reflect.TypeOf((*MockCatalog)(nil).Plugins))
}

// Reload mocks base method
func (m *MockCatalog) Reload(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Reload", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reload indicates an expected call of Reload
func (mr *MockCatalogMockRecorder) Reload(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reload", reflect.TypeOf((*MockCatalog)(nil).Reload), arg0)
}

// Run mocks base method
func (m *MockCatalog) Run(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockCatalogMockRecorder) Run(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCatalog)(nil).Run), arg0)
}

// Stop mocks base method
func (m *MockCatalog) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockCatalogMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockCatalog)(nil).Stop))
}
