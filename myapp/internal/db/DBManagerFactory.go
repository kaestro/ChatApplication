// myapp/internal/db/DBManagerFactory.go
package db

import "errors"

type DBType string

const (
	Postgres DBType = "postgres"
	// 여기에 다른 데이터베이스 타입을 추가할 수 있습니다.
)

type DBManagerFactory struct{}

func (f *DBManagerFactory) CreateDBManager(dbType DBType) (DBManagerInterface, error) {
	switch dbType {
	case Postgres:
		return NewPostgresManager()
	default:
		return nil, errors.New("invalid database type")
	}
}
