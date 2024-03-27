// myapp/internal/db/PostgresManager.go
package db

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresManager struct {
	db *gorm.DB
}

func NewPostgresManager() (DBManagerInterface, error) {
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:rootpassword@localhost:5432/postgres?sslmode=disable" // default value
	}
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	return &PostgresManager{
		db: db,
	}, nil
}

func (m *PostgresManager) Create(value interface{}) error {
	return m.db.Create(value).Error
}

func (m *PostgresManager) ReadAll(out interface{}) error {
	return m.db.Find(out).Error
}

func (m *PostgresManager) Read(out interface{}, field string, value interface{}) error {
	return m.db.Where(field+" = ?", value).First(out).Error
}

func (m *PostgresManager) Update(value interface{}, attrs ...interface{}) error {
	return m.db.Model(value).Update(attrs...).Error
}

func (m *PostgresManager) Delete(value interface{}) error {
	return m.db.Delete(value).Error
}
