// myapp/internal/chat/roomManager_test.go
package chat

import (
	"strconv"
	"testing"
)

func TestRoomManagerCycle(t *testing.T) {
	rm := getRoomManager()

	// Test AddRoom
	rm.addRoom(sampleRoom)
	if !rm.checkRoom(sampleRoomID) {
		t.Errorf("AddRoom failed, expected roomID 123 to exist")
	}

	// Test GetRoom
	gotRoom := rm.getRoom(sampleRoomID)
	if gotRoom != sampleRoom {
		t.Errorf("GetRoom failed, expected %v, got %v", sampleRoom, gotRoom)
	}

	// Test RemoveRoom
	rm.removeRoom(sampleRoomID)
	if rm.checkRoom(sampleRoomID) {
		t.Errorf("RemoveRoom failed, expected roomID %s to be removed", sampleRoomID)
	}
}

func TestRoomManagerCapacity(t *testing.T) {
	rm := getRoomManager()

	// Test AddRoom
	for i := 0; i < maxRooms; i++ {
		room := &room{roomName: strconv.Itoa(i)}
		rm.addRoom(room)
	}

	// Test AddRoom exceeding capacity
	for i := 0; i < maxRooms; i++ {
		if !rm.checkRoom(strconv.Itoa(i)) {
			t.Errorf("AddRoom failed, expected roomID %d to exist", i)
		}
	}
}

func TestRoomManagerGetRoomIDs(t *testing.T) {
	rm := getRoomManager()

	// Test AddRoom
	for i := 0; i < maxRooms; i++ {
		room := &room{roomName: strconv.Itoa(i)}
		rm.addRoom(room)
	}

	// Test GetRoomIDs
	roomIDs := rm.getRoomNames()
	if len(roomIDs) != maxRooms {
		t.Errorf("GetRoomIDs failed, expected %d roomIDs, got %d", maxRooms, len(roomIDs))
		return
	}

	for _, roomID := range roomIDs {
		if !rm.checkRoom(roomID) {
			t.Errorf("GetRoomIDs failed, expected roomID %s to exist", roomID)
		}
	}

	t.Logf("GetRoomIDs passed, expected %d roomIDs", maxRooms)
}

func TestRoomManagerCreateRoom(t *testing.T) {
	rm := getRoomManager()
	room := rm.createNewRoom(sampleRoom.roomName)

	if !rm.checkRoom(room.roomName) {
		t.Errorf("createRoom failed, expected roomID %s to exist", room.roomName)
	}
}

func TestRoomManagerGetRoomCount(t *testing.T) {
	rm := getRoomManager()

	// Test AddRoom
	for i := 0; i < maxRooms; i++ {
		room := &room{roomName: string(rune(i))}
		rm.addRoom(room)
	}

	roomCount := rm.getRoomCount()
	if roomCount != maxRooms {
		t.Errorf("getRoomCount failed, expected %d rooms, got %d", maxRooms, roomCount)
	}
}
