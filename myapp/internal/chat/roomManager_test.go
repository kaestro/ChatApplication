// myapp/internal/chat/roomManager_test.go
package chat

import (
	"strconv"
	"testing"
)

func TestRoomManagerCycle(t *testing.T) {
	rm := GetRoomManager()

	// Test AddRoom
	rm.AddRoom(sampleRoom)
	if !rm.CheckRoom(sampleRoomID) {
		t.Errorf("AddRoom failed, expected roomID 123 to exist")
	}

	// Test GetRoom
	gotRoom := rm.GetRoom(sampleRoomID)
	if gotRoom != sampleRoom {
		t.Errorf("GetRoom failed, expected %v, got %v", sampleRoom, gotRoom)
	}

	// Test RemoveRoom
	rm.RemoveRoom(sampleRoomID)
	if rm.CheckRoom(sampleRoomID) {
		t.Errorf("RemoveRoom failed, expected roomID %s to be removed", sampleRoomID)
	}
}

func TestRoomManagerCapacity(t *testing.T) {
	rm := GetRoomManager()

	// Test AddRoom
	for i := 0; i < maxRooms; i++ {
		room := &Room{roomID: string(rune(i))}
		rm.AddRoom(room)
	}

	// Test AddRoom exceeding capacity
	for i := 0; i < maxRooms; i++ {
		if !rm.CheckRoom(string(rune(i))) {
			t.Errorf("AddRoom failed, expected roomID %d to exist", i)
		}
	}
}

func TestRoomManagerGetRoomIDs(t *testing.T) {
	rm := GetRoomManager()

	// Test AddRoom
	for i := 0; i < maxRooms; i++ {
		room := &Room{roomID: string(rune(i))}
		rm.AddRoom(room)
	}

	// Test GetRoomIDs
	roomIDs := rm.GetRoomIDs()
	if len(roomIDs) != maxRooms {
		t.Errorf("GetRoomIDs failed, expected %d roomIDs, got %d", maxRooms, len(roomIDs))
		return
	}

	for _, roomID := range roomIDs {
		if !rm.CheckRoom(roomID) {
			t.Errorf("GetRoomIDs failed, expected roomID %s to exist", roomID)
		}
	}

	t.Logf("GetRoomIDs passed, expected %d roomIDs", maxRooms)
}

func TestRoomManagerGetNewRoomID(t *testing.T) {
	rm := GetRoomManager()
	roomID := rm.getNewRoomID()

	if strconv.Itoa(rm.lastRoomID) != roomID {
		t.Errorf("getNewRoomID failed, expected roomID %s to exist", roomID)
	}
}

func TestRoomManagerCreateRoom(t *testing.T) {
	rm := GetRoomManager()
	room := rm.createRoom()

	if !rm.CheckRoom(room.roomID) {
		t.Errorf("createRoom failed, expected roomID %s to exist", room.roomID)
	}
}

func TestRoomManagerGetRoomCount(t *testing.T) {
	rm := GetRoomManager()

	// Test AddRoom
	for i := 0; i < maxRooms; i++ {
		room := &Room{roomID: string(rune(i))}
		rm.AddRoom(room)
	}

	roomCount := rm.getRoomCount()
	if roomCount != maxRooms {
		t.Errorf("getRoomCount failed, expected %d rooms, got %d", maxRooms, roomCount)
	}
}
