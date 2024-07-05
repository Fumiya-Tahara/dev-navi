package gorm

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/Fumiya-Tahara/dev-navi/internal/entity"
)

const (
	testName     = "testUser"
	testPassword = "testpass"
)

func TestUserCreate(t *testing.T) {
	t.Parallel()

	db, mock, err := SetUpDBMock()
	require.NoError(t, err, "Expected no error, got %v", err)

	u := &entity.User{
		Name:     testName,
		Password: testPassword,
	}

	// set mock
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	query := regexp.QuoteMeta(
		"INSERT INTO \"users\" (\"name\",\"password\",\"created_at\",\"updated_at\",\"deleted_at\") " +
			"VALUES ($1,$2,$3,$4,$5) RETURNING \"id\"")

	mock.ExpectBegin()
	mock.ExpectQuery(query).
		WillReturnRows(rows)
	mock.ExpectCommit()

	// init repository
	repo := UserRepository{DB: db}
	_, err = repo.Create(u)
	require.NoError(t, err, "Expected no error, got %v", err)
}

func TestUserFindById(t *testing.T) {
	t.Parallel()

	db, mock, err := SetUpDBMock()
	require.NoError(t, err, "Expected no error, got %v", err)

	// set mock
	rows := sqlmock.NewRows([]string{"id", "name", "password"}).AddRow(1, testName, testPassword)
	query := regexp.QuoteMeta(
		"SELECT * FROM \"users\" WHERE \"users\".\"id\" = $1 " +
			"AND \"users\".\"deleted_at\" IS NULL ORDER BY \"users\".\"id\" LIMIT $2")

	mock.ExpectQuery(query).
		WillReturnRows(rows)

	// init repository
	repo := UserRepository{DB: db}
	_, err = repo.FindByID(1)
	require.NoError(t, err, "Expected no error, got %v", err)
}

func TestUserUpdate(t *testing.T) {
	t.Parallel()

	db, mock, err := SetUpDBMock()
	require.NoError(t, err, "Expected no error, got %v", err)

	u := &entity.User{
		ID:       1,
		Name:     testName,
		Password: testPassword,
	}

	// set mock
	rows := sqlmock.NewRows([]string{"id", "name", "password"}).AddRow(1, testName, testPassword)
	selectQuery := regexp.QuoteMeta(
		"SELECT * FROM \"users\" WHERE \"users\".\"id\" = $1 " +
			"AND \"users\".\"deleted_at\" IS NULL ORDER BY \"users\".\"id\" LIMIT $2")

	mock.ExpectQuery(selectQuery).
		WillReturnRows(rows)

	updateQuery := regexp.QuoteMeta(
		"UPDATE \"users\" SET \"name\"=$1,\"password\"=$2,\"created_at\"=$3,\"updated_at\"=$4,\"deleted_at\"=$5 " +
			"WHERE \"users\".\"deleted_at\" IS NULL AND \"id\" = $6")

	mock.ExpectBegin()
	mock.ExpectExec(updateQuery).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// init repository
	repo := UserRepository{DB: db}
	_, err = repo.Update(u)
	require.NoError(t, err, "Expected no error, got %v", err)
}

func TestUserDelete(t *testing.T) {
	t.Parallel()

	db, mock, err := SetUpDBMock()
	require.NoError(t, err, "Expected no error, got %v", err)

	query := regexp.QuoteMeta(
		"UPDATE \"users\" SET \"deleted_at\"=$1 WHERE \"users\".\"id\" = $2 " +
			"AND \"users\".\"deleted_at\" IS NULL")

	// set mock
	mock.ExpectBegin()
	mock.ExpectExec(query).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// init repository
	repo := UserRepository{DB: db}
	err = repo.Delete(1)
	require.NoError(t, err, "Expected no error, got %v", err)
}
