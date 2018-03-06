package structs

import (
	"log"
	"testing"

	"github.com/fatih/color"
)

/*
	Basically we just want to make sure that there's some parity between the definitions and the go structs defined.
*/

func TestBuilding(t *testing.T) {
	log.Printf("Testing buildings")
	var b Building
	err := UnmarshalFromFile("./definitions/building.json", &b)
	if err != nil {
		t.FailNow()
	}
	log.Printf(color.HiGreenString("Pass"))
}

func TestDevice(t *testing.T) {
	log.Printf("Testing device")
	var d Device
	err := UnmarshalFromFile("./definitions/device.json", &d)
	if err != nil {
		t.FailNow()
	}
	log.Printf(color.HiGreenString("Pass"))
}

func TestDeviceType(t *testing.T) {
	log.Printf("Testing devicetypes")
	var dt DeviceType
	err := UnmarshalFromFile("./definitions/device-type.json", &dt)
	if err != nil {
		t.FailNow()
	}
	log.Printf(color.HiGreenString("Pass"))
}

func TestRoom(t *testing.T) {
	log.Printf("Testing rooms")
	var r Room
	err := UnmarshalFromFile("./definitions/room.json", &r)
	if err != nil {
		t.FailNow()
	}
	log.Printf(color.HiGreenString("Pass"))
}

func TestRoomConfiguration(t *testing.T) {
	log.Printf("Testing RoomConfiguration")
	var r RoomConfiguration
	err := UnmarshalFromFile("./definitions/room-configuration.json", &r)
	if err != nil {
		t.FailNow()
	}
	log.Printf(color.HiGreenString("Pass"))
}
