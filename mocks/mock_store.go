// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jonada182/cover-letter-ai-api/internal/store (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	uuid "github.com/google/uuid"
	types "github.com/jonada182/cover-letter-ai-api/types"
	mongo "go.mongodb.org/mongo-driver/mongo"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockStore) Connect() (*mongo.Client, context.Context, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(*mongo.Client)
	ret1, _ := ret[1].(context.Context)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Connect indicates an expected call of Connect.
func (mr *MockStoreMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockStore)(nil).Connect))
}

// DeleteJobApplication mocks base method.
func (m *MockStore) DeleteJobApplication(arg0 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteJobApplication", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteJobApplication indicates an expected call of DeleteJobApplication.
func (mr *MockStoreMockRecorder) DeleteJobApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJobApplication", reflect.TypeOf((*MockStore)(nil).DeleteJobApplication), arg0)
}

// Disconnect mocks base method.
func (m *MockStore) Disconnect(arg0 context.Context, arg1 *mongo.Client) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disconnect", arg0, arg1)
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockStoreMockRecorder) Disconnect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockStore)(nil).Disconnect), arg0, arg1)
}

// GetCareerProfileByEmail mocks base method.
func (m *MockStore) GetCareerProfileByEmail(arg0 string) (*types.CareerProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCareerProfileByEmail", arg0)
	ret0, _ := ret[0].(*types.CareerProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCareerProfileByEmail indicates an expected call of GetCareerProfileByEmail.
func (mr *MockStoreMockRecorder) GetCareerProfileByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCareerProfileByEmail", reflect.TypeOf((*MockStore)(nil).GetCareerProfileByEmail), arg0)
}

// GetCareerProfileByID mocks base method.
func (m *MockStore) GetCareerProfileByID(arg0 uuid.UUID) (*types.CareerProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCareerProfileByID", arg0)
	ret0, _ := ret[0].(*types.CareerProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCareerProfileByID indicates an expected call of GetCareerProfileByID.
func (mr *MockStoreMockRecorder) GetCareerProfileByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCareerProfileByID", reflect.TypeOf((*MockStore)(nil).GetCareerProfileByID), arg0)
}

// GetJobApplicationByID mocks base method.
func (m *MockStore) GetJobApplicationByID(arg0 uuid.UUID) (*types.JobApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobApplicationByID", arg0)
	ret0, _ := ret[0].(*types.JobApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobApplicationByID indicates an expected call of GetJobApplicationByID.
func (mr *MockStoreMockRecorder) GetJobApplicationByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobApplicationByID", reflect.TypeOf((*MockStore)(nil).GetJobApplicationByID), arg0)
}

// GetJobApplications mocks base method.
func (m *MockStore) GetJobApplications(arg0 uuid.UUID) (*[]types.JobApplication, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobApplications", arg0)
	ret0, _ := ret[0].(*[]types.JobApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobApplications indicates an expected call of GetJobApplications.
func (mr *MockStoreMockRecorder) GetJobApplications(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobApplications", reflect.TypeOf((*MockStore)(nil).GetJobApplications), arg0)
}

// StoreAccessToken mocks base method.
func (m *MockStore) StoreAccessToken(arg0 uuid.UUID, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreAccessToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreAccessToken indicates an expected call of StoreAccessToken.
func (mr *MockStoreMockRecorder) StoreAccessToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreAccessToken", reflect.TypeOf((*MockStore)(nil).StoreAccessToken), arg0, arg1)
}

// StoreCareerProfile mocks base method.
func (m *MockStore) StoreCareerProfile(arg0 *types.CareerProfile) (*types.CareerProfile, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreCareerProfile", arg0)
	ret0, _ := ret[0].(*types.CareerProfile)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StoreCareerProfile indicates an expected call of StoreCareerProfile.
func (mr *MockStoreMockRecorder) StoreCareerProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreCareerProfile", reflect.TypeOf((*MockStore)(nil).StoreCareerProfile), arg0)
}

// StoreJobApplication mocks base method.
func (m *MockStore) StoreJobApplication(arg0 *types.JobApplication) (*types.JobApplication, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreJobApplication", arg0)
	ret0, _ := ret[0].(*types.JobApplication)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// StoreJobApplication indicates an expected call of StoreJobApplication.
func (mr *MockStoreMockRecorder) StoreJobApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreJobApplication", reflect.TypeOf((*MockStore)(nil).StoreJobApplication), arg0)
}

// ValidateAccessToken mocks base method.
func (m *MockStore) ValidateAccessToken(arg0 uuid.UUID, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAccessToken", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateAccessToken indicates an expected call of ValidateAccessToken.
func (mr *MockStoreMockRecorder) ValidateAccessToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAccessToken", reflect.TypeOf((*MockStore)(nil).ValidateAccessToken), arg0, arg1)
}
