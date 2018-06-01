// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/agent/keymanager (interfaces: KeyManager)

// Package mock_keymanager is a generated GoMock package.
package mock_keymanager

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	keymanager "github.com/spiffe/spire/proto/agent/keymanager"
	plugin "github.com/spiffe/spire/proto/common/plugin"
	reflect "reflect"
)

// MockKeyManager is a mock of KeyManager interface
type MockKeyManager struct {
	ctrl     *gomock.Controller
	recorder *MockKeyManagerMockRecorder
}

// MockKeyManagerMockRecorder is the mock recorder for MockKeyManager
type MockKeyManagerMockRecorder struct {
	mock *MockKeyManager
}

// NewMockKeyManager creates a new mock instance
func NewMockKeyManager(ctrl *gomock.Controller) *MockKeyManager {
	mock := &MockKeyManager{ctrl: ctrl}
	mock.recorder = &MockKeyManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeyManager) EXPECT() *MockKeyManagerMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockKeyManager) Configure(arg0 context.Context, arg1 *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockKeyManagerMockRecorder) Configure(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockKeyManager)(nil).Configure), arg0, arg1)
}

// FetchPrivateKey mocks base method
func (m *MockKeyManager) FetchPrivateKey(arg0 context.Context, arg1 *keymanager.FetchPrivateKeyRequest) (*keymanager.FetchPrivateKeyResponse, error) {
	ret := m.ctrl.Call(m, "FetchPrivateKey", arg0, arg1)
	ret0, _ := ret[0].(*keymanager.FetchPrivateKeyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchPrivateKey indicates an expected call of FetchPrivateKey
func (mr *MockKeyManagerMockRecorder) FetchPrivateKey(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchPrivateKey", reflect.TypeOf((*MockKeyManager)(nil).FetchPrivateKey), arg0, arg1)
}

// GenerateKeyPair mocks base method
func (m *MockKeyManager) GenerateKeyPair(arg0 context.Context, arg1 *keymanager.GenerateKeyPairRequest) (*keymanager.GenerateKeyPairResponse, error) {
	ret := m.ctrl.Call(m, "GenerateKeyPair", arg0, arg1)
	ret0, _ := ret[0].(*keymanager.GenerateKeyPairResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateKeyPair indicates an expected call of GenerateKeyPair
func (mr *MockKeyManagerMockRecorder) GenerateKeyPair(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateKeyPair", reflect.TypeOf((*MockKeyManager)(nil).GenerateKeyPair), arg0, arg1)
}

// GetPluginInfo mocks base method
func (m *MockKeyManager) GetPluginInfo(arg0 context.Context, arg1 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0, arg1)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockKeyManagerMockRecorder) GetPluginInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockKeyManager)(nil).GetPluginInfo), arg0, arg1)
}
