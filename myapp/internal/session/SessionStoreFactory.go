package session

type StoreFactory interface {
	Create(sessionTypeNum SessionType) SessionStore
}
