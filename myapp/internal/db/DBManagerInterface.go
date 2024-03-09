// myapp/internal/db/DBManagerInterface.go
package db

type DBManagerInterface interface {
	Create(value interface{}) error
	ReadAll(out interface{}) error
	Read(out interface{}, field string, value interface{}) error
	Update(value interface{}, attrs ...interface{}) error
	Delete(value interface{}) error
}
