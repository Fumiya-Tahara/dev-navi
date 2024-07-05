package gorm

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/Fumiya-Tahara/dev-navi/internal/entity"

	"github.com/stretchr/testify/assert"
)

var (
	testName     = "testUser"
	testPassword = "testpass"
)

func TestUserCreate(t *testing.T) {
	db, mock, err := SetUpDBMock()
	assert.NoError(t, err, "Expected no error, got %v", err)

	u := &entity.User{
		Name:     testName,
		Password: testPassword,
	}

	//set mock
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	query := regexp.QuoteMeta("INSERT INTO \"users\" (\"name\",\"password\",\"created_at\",\"updated_at\",\"deleted_at\") VALUES ($1,$2,$3,$4,$5) RETURNING \"id\"")

	mock.ExpectBegin()
	mock.ExpectQuery(query).
		WillReturnRows(rows)
	mock.ExpectCommit()

	// init repository
	repo := UserRepository{DB: db}
	_, err = repo.Create(u)
	assert.NoError(t, err, "Expected no error, got %v", err)
}
