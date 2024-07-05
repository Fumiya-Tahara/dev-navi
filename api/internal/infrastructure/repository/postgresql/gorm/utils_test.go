package gorm

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func SetUpDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create sqlmock: %w", err)
	}

	return gdb, mock, fmt.Errorf("failed to create sqlmock: %w", err)
}
