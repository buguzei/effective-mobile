// Code generated by MockGen. DO NOT EDIT.
// Source: auth/internal/repo/repo.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	reflect "reflect"

	models "github.com/buguzei/effective-mobile/auth/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// FindUserByEmail mocks base method.
func (m *MockUserRepo) FindUserByEmail(ctx context.Context, email string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", ctx, email)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockUserRepoMockRecorder) FindUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockUserRepo)(nil).FindUserByEmail), ctx, email)
}

// IsUserExist mocks base method.
func (m *MockUserRepo) IsUserExist(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserExist", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserExist indicates an expected call of IsUserExist.
func (mr *MockUserRepoMockRecorder) IsUserExist(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserExist", reflect.TypeOf((*MockUserRepo)(nil).IsUserExist), ctx, email)
}

// NewUser mocks base method.
func (m *MockUserRepo) NewUser(ctx context.Context, user models.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewUser indicates an expected call of NewUser.
func (mr *MockUserRepoMockRecorder) NewUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewUser", reflect.TypeOf((*MockUserRepo)(nil).NewUser), ctx, user)
}

// VerifyEmailAndPass mocks base method.
func (m *MockUserRepo) VerifyEmailAndPass(ctx context.Context, email, password string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyEmailAndPass", ctx, email, password)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyEmailAndPass indicates an expected call of VerifyEmailAndPass.
func (mr *MockUserRepoMockRecorder) VerifyEmailAndPass(ctx, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyEmailAndPass", reflect.TypeOf((*MockUserRepo)(nil).VerifyEmailAndPass), ctx, email, password)
}

// MockRefreshRepo is a mock of RefreshRepo interface.
type MockRefreshRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRefreshRepoMockRecorder
}

// MockRefreshRepoMockRecorder is the mock recorder for MockRefreshRepo.
type MockRefreshRepoMockRecorder struct {
	mock *MockRefreshRepo
}

// NewMockRefreshRepo creates a new mock instance.
func NewMockRefreshRepo(ctrl *gomock.Controller) *MockRefreshRepo {
	mock := &MockRefreshRepo{ctrl: ctrl}
	mock.recorder = &MockRefreshRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRefreshRepo) EXPECT() *MockRefreshRepoMockRecorder {
	return m.recorder
}

// GetRefresh mocks base method.
func (m *MockRefreshRepo) GetRefresh(ctx context.Context, id int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefresh", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRefresh indicates an expected call of GetRefresh.
func (mr *MockRefreshRepoMockRecorder) GetRefresh(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefresh", reflect.TypeOf((*MockRefreshRepo)(nil).GetRefresh), ctx, id)
}

// SetRefresh mocks base method.
func (m *MockRefreshRepo) SetRefresh(ctx context.Context, refreshToken string, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetRefresh", ctx, refreshToken, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRefresh indicates an expected call of SetRefresh.
func (mr *MockRefreshRepoMockRecorder) SetRefresh(ctx, refreshToken, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRefresh", reflect.TypeOf((*MockRefreshRepo)(nil).SetRefresh), ctx, refreshToken, id)
}
