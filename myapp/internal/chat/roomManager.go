// myapp/internal/chat/roomManager.go
package chat

import (
	"fmt"
	"sync"
)

var (
	roomOnce   sync.Once
	rmInstance *roomManager
)

// roomManager는 방의 유무를 확인, 생성, 제거, 조회를 담당한다.
// Singleton 객체로 구현되어 있다.
type roomManager struct {
	rooms map[string]*room
}

func getRoomManager() *roomManager {
	roomOnce.Do(func() {
		rmInstance = &roomManager{
			rooms: make(map[string]*room),
		}
	})

	return rmInstance
}

func (rm *roomManager) checkRoom(roomName string) bool {
	_, ok := rm.rooms[roomName]
	return ok
}

func (rm *roomManager) getRoom(roomName string) *room {
	if !rm.checkRoom(roomName) {
		fmt.Println("Room with roomID", roomName, "does not exist")
		return nil
	}

	return rm.rooms[roomName]
}

// TODO: fmt 대신 별개의 로거를 사용하도록 변경
func (rm *roomManager) addRoom(room *room) {
	if rm.checkRoom(room.roomName) {
		// fmt.Println("Room with roomID", room.roomID, "already exists")
		return
	}
	rm.rooms[room.roomName] = room
}

func (rm *roomManager) removeRoom(roomName string) {
	if !rm.checkRoom(roomName) {
		fmt.Println("Room with roomID", roomName, "does not exist")
		return
	}

	rm.rooms[roomName].closeRoom()
	delete(rm.rooms, roomName)
}

// Question: wouldn't it be better to just return room pointers?
func (rm *roomManager) getRoomNames() []string {
	roomNames := make([]string, 0, len(rm.rooms))
	for roomName := range rm.rooms {
		roomNames = append(roomNames, roomName)
	}
	return roomNames
}

func (rm *roomManager) createNewRoom(roomName string) *room {
	//roomID := rm.getNewRoomID()
	room := newRoom(roomName)
	rm.addRoom(room)
	return room
}

func (rm *roomManager) getRoomCount() int {
	return len(rm.rooms)
}

func (rm *roomManager) clearRooms() {
	rm.rooms = make(map[string]*room)
}
