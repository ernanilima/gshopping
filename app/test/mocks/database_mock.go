// Code generated by MockGen. DO NOT EDIT.
// Source: app/repository/database/connection.go

// Package mock_database is a generated GoMock package.
package mocks

import (
        sql "database/sql"
        reflect "reflect"

        "go.uber.org/mock/gomock"
)

// MockDatabaseConnector is a mock of DatabaseConnector interface.
type MockDatabaseConnector struct {
        ctrl     *gomock.Controller
        recorder *MockDatabaseConnectorMockRecorder
}

// MockDatabaseConnectorMockRecorder is the mock recorder for MockDatabaseConnector.
type MockDatabaseConnectorMockRecorder struct {
        mock *MockDatabaseConnector
}

// NewMockDatabaseConnector creates a new mock instance.
func NewMockDatabaseConnector(ctrl *gomock.Controller) *MockDatabaseConnector {
        mock := &MockDatabaseConnector{ctrl: ctrl}
        mock.recorder = &MockDatabaseConnectorMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseConnector) EXPECT() *MockDatabaseConnectorMockRecorder {
        return m.recorder
}

// OpenConnection mocks base method.
func (m *MockDatabaseConnector) OpenConnection() *sql.DB {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "OpenConnection")
        ret0, _ := ret[0].(*sql.DB)
        return ret0
}

// OpenConnection indicates an expected call of OpenConnection.
func (mr *MockDatabaseConnectorMockRecorder) OpenConnection() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenConnection", reflect.TypeOf((*MockDatabaseConnector)(nil).OpenConnection))
}

// UPMigrations mocks base method.
func (m *MockDatabaseConnector) UPMigrations() {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "UPMigrations")
}

// UPMigrations indicates an expected call of UPMigrations.
func (mr *MockDatabaseConnectorMockRecorder) UPMigrations() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UPMigrations", reflect.TypeOf((*MockDatabaseConnector)(nil).UPMigrations))
}