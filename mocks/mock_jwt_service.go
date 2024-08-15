// Code generated by MockGen. DO NOT EDIT.
// Source: Domain/jwt.service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/solo21-12/A2SV_back_end_track/tree/main/task_seven/Domain"
)

// MockJwtService is a mock of JwtService interface.
type MockJwtService struct {
	ctrl     *gomock.Controller
	recorder *MockJwtServiceMockRecorder
}

// MockJwtServiceMockRecorder is the mock recorder for MockJwtService.
type MockJwtServiceMockRecorder struct {
	mock *MockJwtService
}

// NewMockJwtService creates a new mock instance.
func NewMockJwtService(ctrl *gomock.Controller) *MockJwtService {
	mock := &MockJwtService{ctrl: ctrl}
	mock.recorder = &MockJwtServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJwtService) EXPECT() *MockJwtServiceMockRecorder {
	return m.recorder
}

// CreateAccessToken mocks base method.
func (m *MockJwtService) CreateAccessToken(user domain.UserDTO) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccessToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccessToken indicates an expected call of CreateAccessToken.
func (mr *MockJwtServiceMockRecorder) CreateAccessToken(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccessToken", reflect.TypeOf((*MockJwtService)(nil).CreateAccessToken), user)
}

// GetClaims mocks base method.
func (m *MockJwtService) GetClaims(authHeader string) (*domain.JWTCustome, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClaims", authHeader)
	ret0, _ := ret[0].(*domain.JWTCustome)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClaims indicates an expected call of GetClaims.
func (mr *MockJwtServiceMockRecorder) GetClaims(authHeader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClaims", reflect.TypeOf((*MockJwtService)(nil).GetClaims), authHeader)
}

// ValidateAuthHeader mocks base method.
func (m *MockJwtService) ValidateAuthHeader(authHeader string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAuthHeader", authHeader)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateAuthHeader indicates an expected call of ValidateAuthHeader.
func (mr *MockJwtServiceMockRecorder) ValidateAuthHeader(authHeader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAuthHeader", reflect.TypeOf((*MockJwtService)(nil).ValidateAuthHeader), authHeader)
}

// ValidateToken mocks base method.
func (m *MockJwtService) ValidateToken(tokenStr string) (*domain.JWTCustome, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", tokenStr)
	ret0, _ := ret[0].(*domain.JWTCustome)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken.
func (mr *MockJwtServiceMockRecorder) ValidateToken(tokenStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockJwtService)(nil).ValidateToken), tokenStr)
}
