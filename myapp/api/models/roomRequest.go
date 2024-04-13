// myapp/api/models/roomRequest.go
package models

type RoomRequest struct {
	RoomName       string `json:"roomName"`
	LoginSessionID string `json:"loginSessionID"`
	EmailAddress   string `json:"emailAddress"`
	Password       string `json:"password"`
}
