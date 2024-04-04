// myapp/internal/chat/roomManager.go
package chat

import (
	"fmt"
	"sync"
)

var (
	roomOnce    sync.Once
	roomManager *RoomManager
)

type RoomManager struct {
	rooms map[string]*Room
}

func GetRoomManager() *RoomManager {
	roomOnce.Do(func() {
		roomManager = &RoomManager{
			rooms: make(map[string]*Room),
		}
	})

	return roomManager
}

func (rm *RoomManager) CheckRoom(roomID string) bool {
	_, ok := rm.rooms[roomID]
	return ok
}

func (rm *RoomManager) GetRoom(roomID string) *Room {
	if !rm.CheckRoom(roomID) {
		fmt.Println("Room with roomID", roomID, "does not exist")
		return nil
	}

	return rm.rooms[roomID]
}

func (rm *RoomManager) AddRoom(room *Room) {
	if rm.CheckRoom(room.roomID) {
		fmt.Println("Room with roomID", room.roomID, "already exists")
		return
	}
	rm.rooms[room.roomID] = room
}

func (rm *RoomManager) RemoveRoom(roomID string) {
	if !rm.CheckRoom(roomID) {
		fmt.Println("Room with roomID", roomID, "does not exist")
		return
	}
	delete(rm.rooms, roomID)
}
