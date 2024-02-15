package db

import (
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

var (
	once sync.Once

	db *sql.DB
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("postgres", "postgres://postgres:rootpassword@localhost:5432/postgres?sslmode=disable")
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}
	})

	return db
}
