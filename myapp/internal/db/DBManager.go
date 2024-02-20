// db/DBManager.go

package db

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	once sync.Once

	db *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open("postgres", "postgres://postgres:rootpassword@localhost:5432/postgres?sslmode=disable")
		if err != nil {
			panic(err)
		}
	})

	return db
}
