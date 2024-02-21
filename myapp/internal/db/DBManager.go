// db/DBManager.go

package db

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// 초기화를 한 번만 실행해고 나서 연결을 저장해두기 위해
// db를 전역 변수로 선언한다.
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
