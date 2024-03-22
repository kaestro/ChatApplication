// myapp/internal/db/DBManager.go
package db

import (
	"sync"
)

var (
	once    sync.Once
	manager DBManagerInterface
)

func GetDBManager() DBManagerInterface {
	once.Do(func() {
		factory := &DBManagerFactory{}
		var err error
		manager, err = factory.CreateDBManager(Postgres)
		if err != nil {
			panic(err)
		}
	})

	return manager
}
