// myapp/internal/chat/roomManager_test.go
package chat

import (
	"testing"
)

var (
	maxRooms     = 100
	sampleRoomID = "123"
	sampleRoom   = &Room{roomID: sampleRoomID}
)

func TestRoomManager(t *testing.T) {
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
