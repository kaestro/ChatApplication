// session/Store.go
package session

type Store interface {
	GetSession(key string) (string, error)
	SetSession(key string, value string) error
	DeleteSession(key string) error
	IsSessionValid(key string, emailAddress string) bool
}
