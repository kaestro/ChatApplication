// myapp/internal/db/DBManager.go
package db

import (
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBManager struct {
	db *gorm.DB
}

var (
	once sync.Once

	manager DBManagerInterface
)

func GetDBManager() DBManagerInterface {
	once.Do(func() {
		var err error
		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			dbURL = "postgres://postgres:rootpassword@localhost:5432/postgres?sslmode=disable" // default value
		}
		db, err := gorm.Open("postgres", dbURL)
		if err != nil {
			panic(err)
		}

		manager = &DBManager{
			db: db,
		}
	})

	return manager
}

func (m *DBManager) Create(value interface{}) error {
	return m.db.Create(value).Error
}

func (m *DBManager) ReadAll(out interface{}) error {
	return m.db.Find(out).Error
}

func (m *DBManager) Read(out interface{}, field string, value interface{}) error {
	return m.db.Where(field+" = ?", value).First(out).Error
}

func (m *DBManager) Update(value interface{}, attrs ...interface{}) error {
	return m.db.Model(value).Update(attrs...).Error
}

func (m *DBManager) Delete(value interface{}) error {
	return m.db.Delete(value).Error
}
