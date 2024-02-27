// myapp/internal/db/DBManager.go
package db

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DBManager struct {
	db *gorm.DB
}

var (
	once sync.Once

	manager *DBManager
)

func GetDBManager() *DBManager {
	once.Do(func() {
		var err error
		db, err := gorm.Open("postgres", "postgres://postgres:rootpassword@postgresql:5432/postgres?sslmode=disable")
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
