package couch

import "testing"

func TestRoom(t *testing.T) {
	defer setupDatabase(t)(t)
}

func testRoomCreateWithNewConfiguration(t *testing.T) {

	room := structs.Room{}
	err := structs.UnmarshallFromFile(testDir + "/new_room_a.json")
	if err != nil {
		t.Logf("Error reading in %v: %v", "new_room_a.json", err.Error())
		t.Fail()
	}

	newrm, err = CreateRoom(room)
}
