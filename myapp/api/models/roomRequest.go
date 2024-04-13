// myapp/api/models/roomRequest.go
package models

type RoomRequest struct {
	RoomName       string `json:"roomName"`
	LoginSessionID string `json:"loginSessionID"`
	EmailAddress   string `json:"emailAddress"`
	Password       string `json:"password"`
}

func NewRoomRequest(roomName, loginSessionID, emailAddress, password string) RoomRequest {
	return RoomRequest{
		RoomName:       roomName,
		LoginSessionID: loginSessionID,
		EmailAddress:   emailAddress,
		Password:       password,
	}
}
