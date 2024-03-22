// myapp/internal/db/DBManager_test.go
package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDBManager(t *testing.T) {
	factory := &DBManagerFactory{
		NewDBManager: func(dbType DBType) (DBManagerInterface, error) {
			return &PostgresManager{}, nil
		},
	}

	manager := GetDBManagerWithFactory(factory)
	assert.IsType(t, &PostgresManager{}, manager)
}
