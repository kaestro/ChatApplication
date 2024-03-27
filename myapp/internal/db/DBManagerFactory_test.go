// myapp/internal/db/DBManagerFactory_test.go
package db

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDBManager(t *testing.T) {
	f := &DBManagerFactory{
		NewDBManager: func(dbType DBType) (DBManagerInterface, error) {
			if dbType == "invalid" {
				return nil, errors.New("invalid database type")
			}
			return &PostgresManager{}, nil
		},
	}

	manager, err := f.CreateDBManager(Postgres)
	assert.NoError(t, err)
	assert.IsType(t, &PostgresManager{}, manager)

	_, err = f.CreateDBManager("invalid")
	assert.Error(t, err)
	assert.Equal(t, errors.New("invalid database type"), err)
}
