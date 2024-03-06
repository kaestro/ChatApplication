// myapp/internal/session/SessionStore.go
package session

type SessionType int

const (
	LoginSession SessionType = 0
	OtherSession SessionType = -1
)

type SessionStore interface {
	GetSession(key string) (string, error)
	SetSession(key string, value string) error
	DeleteSession(key string) error
	IsSessionValid(key string, emailAddress string) bool
}
