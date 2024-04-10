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

// RoomManager는 방의 유무를 확인, 생성, 제거, 조회를 담당한다.
// Singleton 객체로 구현되어 있다.
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

// TODO: fmt 대신 별개의 로거를 사용하도록 변경
func (rm *RoomManager) AddRoom(room *Room) {
	if rm.CheckRoom(room.roomID) {
		// fmt.Println("Room with roomID", room.roomID, "already exists")
		return
	}
	rm.rooms[room.roomID] = room
}

func (rm *RoomManager) RemoveRoom(roomID string) {
	if !rm.CheckRoom(roomID) {
		fmt.Println("Room with roomID", roomID, "does not exist")
		return
	}

	rm.rooms[roomID].closeRoom()
	delete(rm.rooms, roomID)
}

// Question: wouldn't it be better to just return room pointers?
func (rm *RoomManager) GetRoomIDs() []string {
	roomIDs := make([]string, 0, len(rm.rooms))
	for roomID := range rm.rooms {
		roomIDs = append(roomIDs, roomID)
	}
	return roomIDs
}
