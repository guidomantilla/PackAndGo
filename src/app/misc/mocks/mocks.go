package mocks

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockDBDataSource struct {
	Connection *sql.DB
	Mock       sqlmock.Sqlmock
}

func NewMockDBDataSource() *MockDBDataSource {
	connection, mock, _ := sqlmock.New()
	return &MockDBDataSource{
		Connection: connection,
		Mock:       mock,
	}
}

func (mockDBDataSource *MockDBDataSource) GetDatabase() *sql.DB {
	return mockDBDataSource.Connection
}
