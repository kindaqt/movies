// Code generated by MockGen. DO NOT EDIT.
// Source: ./models/movies.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/kindaqt/movies/api/models"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetMovies mocks base method
func (m *MockStore) GetMovies() ([]models.Movie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovies")
	ret0, _ := ret[0].([]models.Movie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovies indicates an expected call of GetMovies
func (mr *MockStoreMockRecorder) GetMovies() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovies", reflect.TypeOf((*MockStore)(nil).GetMovies))
}

// UpdateWatched mocks base method
func (m *MockStore) UpdateWatched(id string, value bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWatched", id, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateWatched indicates an expected call of UpdateWatched
func (mr *MockStoreMockRecorder) UpdateWatched(id, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWatched", reflect.TypeOf((*MockStore)(nil).UpdateWatched), id, value)
}

// DeleteMovie mocks base method
func (m *MockStore) DeleteMovie(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMovie", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMovie indicates an expected call of DeleteMovie
func (mr *MockStoreMockRecorder) DeleteMovie(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockStore)(nil).DeleteMovie), id)
}

// CreateMovie mocks base method
func (m *MockStore) CreateMovie(movie *models.Movie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMovie", movie)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMovie indicates an expected call of CreateMovie
func (mr *MockStoreMockRecorder) CreateMovie(movie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMovie", reflect.TypeOf((*MockStore)(nil).CreateMovie), movie)
}
