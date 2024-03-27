// myapp/internal/db/PostgresManager_test.go
package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func newMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		return nil, nil, err
	}
	return gormDB, mock, nil
}

func TestCreate(t *testing.T) {
	gormDB, mock, err := newMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	m := &PostgresManager{db: gormDB}
	err = m.Create(&struct{}{})

	assert.NoError(t, err)
}

func TestReadAll(t *testing.T) {
	gormDB, mock, err := newMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("^SELECT (.+) FROM").WillReturnRows(rows)

	m := &PostgresManager{db: gormDB}
	err = m.ReadAll(&[]struct{}{})

	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	gormDB, mock, err := newMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectQuery("^SELECT (.+) FROM").WillReturnRows(rows)

	m := &PostgresManager{db: gormDB}
	err = m.Read(&struct{}{}, "id", 1)

	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	gormDB, mock, err := newMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	m := &PostgresManager{db: gormDB}
	err = m.Update(&struct{}{}, "field", "value")

	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	gormDB, mock, err := newMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	m := &PostgresManager{db: gormDB}
	err = m.Delete(&struct{}{})

	assert.NoError(t, err)
}
