// Code generated by MockGen. DO NOT EDIT.
// Source: ./storage_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorageService is a mock of StorageService interface.
type MockStorageService struct {
	ctrl     *gomock.Controller
	recorder *MockStorageServiceMockRecorder
}

// MockStorageServiceMockRecorder is the mock recorder for MockStorageService.
type MockStorageServiceMockRecorder struct {
	mock *MockStorageService
}

// NewMockStorageService creates a new mock instance.
func NewMockStorageService(ctrl *gomock.Controller) *MockStorageService {
	mock := &MockStorageService{ctrl: ctrl}
	mock.recorder = &MockStorageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageService) EXPECT() *MockStorageServiceMockRecorder {
	return m.recorder
}

// GetURL mocks base method.
func (m *MockStorageService) GetURL(ctx context.Context, shortURL string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", ctx, shortURL)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURL indicates an expected call of GetURL.
func (mr *MockStorageServiceMockRecorder) GetURL(ctx, shortURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockStorageService)(nil).GetURL), ctx, shortURL)
}

// SaveURL mocks base method.
func (m *MockStorageService) SaveURL(ctx context.Context, shortURL, originalURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveURL", ctx, shortURL, originalURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveURL indicates an expected call of SaveURL.
func (mr *MockStorageServiceMockRecorder) SaveURL(ctx, shortURL, originalURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveURL", reflect.TypeOf((*MockStorageService)(nil).SaveURL), ctx, shortURL, originalURL)
}