// myapp/internal/db/DBManager.go
package db

import (
	"sync"
)

var (
	once    sync.Once
	manager DBManagerInterface
)

func GetDBManagerWithFactory(factory *DBManagerFactory) DBManagerInterface {
	once.Do(func() {
		var err error
		manager, err = factory.CreateDBManager(Postgres)
		if err != nil {
			panic(err)
		}
	})

	return manager
}

func GetDBManager() DBManagerInterface {
	return GetDBManagerWithFactory(&DBManagerFactory{})
}
