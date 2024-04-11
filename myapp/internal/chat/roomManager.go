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
	rooms      map[string]*room
	lastRoomID int
}

func getRoomManager() *roomManager {
	roomOnce.Do(func() {
		rmInstance = &roomManager{
			rooms:      make(map[string]*room),
			lastRoomID: 0,
		}
	})

	return rmInstance
}

func (rm *roomManager) checkRoom(roomID string) bool {
	_, ok := rm.rooms[roomID]
	return ok
}

func (rm *roomManager) getRoom(roomID string) *room {
	if !rm.checkRoom(roomID) {
		fmt.Println("Room with roomID", roomID, "does not exist")
		return nil
	}

	return rm.rooms[roomID]
}

// TODO: fmt 대신 별개의 로거를 사용하도록 변경
func (rm *roomManager) AddRoom(room *room) {
	if rm.checkRoom(room.roomID) {
		// fmt.Println("Room with roomID", room.roomID, "already exists")
		return
	}
	rm.rooms[room.roomID] = room
}

func (rm *roomManager) removeRoom(roomID string) {
	if !rm.checkRoom(roomID) {
		fmt.Println("Room with roomID", roomID, "does not exist")
		return
	}

	rm.rooms[roomID].closeRoom()
	delete(rm.rooms, roomID)
}

// Question: wouldn't it be better to just return room pointers?
func (rm *roomManager) getRoomIDs() []string {
	roomIDs := make([]string, 0, len(rm.rooms))
	for roomID := range rm.rooms {
		roomIDs = append(roomIDs, roomID)
	}
	return roomIDs
}

func (rm *roomManager) getNewRoomID() string {
	rm.lastRoomID++
	return fmt.Sprintf("%d", rm.lastRoomID)
}

func (rm *roomManager) createNewRoom() *room {
	roomID := rm.getNewRoomID()
	room := newRoom(roomID)
	rm.AddRoom(room)
	return room
}

func (rm *roomManager) getRoomCount() int {
	return len(rm.rooms)
}
