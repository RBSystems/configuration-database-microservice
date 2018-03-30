package couch

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/byuoitav/configuration-database-microservice/log"
	"github.com/byuoitav/configuration-database-microservice/structs"
	"github.com/stretchr/testify/assert"
)

var testDir = `./test-data`

func setupDatabase(t *testing.T) func(t *testing.T) {
	log.CFG.OutputPaths = []string{}
	tmp, _ := log.CFG.Build()
	log.L = tmp.Sugar()

	t.Log("Setting up database for testing")

	//set up our environment variables
	oldCouchAddress := os.Getenv("COUCH_ADDRESS")
	oldCouchUsername := os.Getenv("COUCH_USERNAME")
	oldCouchPassword := os.Getenv("COUCH_PASSWORD")
	oldLoggingLocation := os.Getenv("LOGGING_FILE_LOCATION")

	os.Setenv("COUCH_ADDRESS", os.Getenv("COUCH_TESTING_ADDRESS"))
	os.Setenv("COUCH_USERNAME", os.Getenv("COUCH_TESTING_USERNAME"))
	os.Setenv("COUCH_PASSWORD", os.Getenv("COUCH_TESTING_PASSWORD"))
	os.Setenv("LOGGING_FILE_LOCATION", os.Getenv("TEST_LOGGING_FILE_LOCATION"))

	//now we go and set up the database

	//find all of the setup files to be read in

	files, err := ioutil.ReadDir(testDir)
	if err != nil {
		msg := fmt.Sprintf("Couldn't read the database setup director: %v", err.Error())
		t.Log(msg)
		t.FailNow()
	}

	setupScriptRegex := regexp.MustCompile(`setup_([A-Z,a-z]+)`)

	//wipe out the current databases.
	wipeDatabases()

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		//check it for a setup
		matches := setupScriptRegex.FindStringSubmatch(f.Name())
		if len(matches) == 0 {
			continue
		}
		t.Logf("Reading in %v", f.Name())

		switch matches[1] {
		case "buildings":
			building := structs.Building{}
			//add a building
			err := structs.UnmarshalFromFile(testDir+"/"+f.Name(), &building)
			if err != nil {
				t.Logf("couldn't set up database. Error reading in %v: %v", f.Name(), err.Error())
				t.FailNow()
			}

			_, err = CreateBuilding(building)
			if err != nil {
				t.Logf("couldn't set up database. Error creating building %v: %v", f.Name(), err.Error())
				t.FailNow()
			}
		case "devicetype":
			dt := structs.DeviceType{}
			//add a building
			err := structs.UnmarshalFromFile(testDir+"/"+f.Name(), &dt)
			if err != nil {
				t.Logf("couldn't set up database. Error reading in %v: %v", f.Name(), err.Error())
				t.FailNow()
			}

			_, err = CreateDeviceType(dt)
			if err != nil {
				t.Logf("couldn't set up database. Error creating devicetype %v: %v", f.Name(), err.Error())
				t.FailNow()
			}
		case "room":
			dt := structs.Room{}
			//add a building
			err := structs.UnmarshalFromFile(testDir+"/"+f.Name(), &dt)
			if err != nil {
				t.Logf("couldn't set up database. Error reading in %v: %v", f.Name(), err.Error())
				t.FailNow()
			}

			_, err = CreateRoom(dt)
			if err != nil {
				t.Logf("couldn't set up database. Error creating room %v: %v", f.Name(), err.Error())
				t.FailNow()
			}
		case "device":
			dt := structs.Device{}
			//add a building
			err := structs.UnmarshalFromFile(testDir+"/"+f.Name(), &dt)
			if err != nil {
				t.Logf("couldn't set up database. Error reading in %v: %v", f.Name(), err.Error())
				t.FailNow()
			}

			_, err = CreateDevice(dt)
			if err != nil {
				t.Logf("couldn't set up database. Error creating device %v: %v", f.Name(), err.Error())
				t.FailNow()
			}
		case "roomconfig":
			dt := structs.RoomConfiguration{}
			//add a building
			err := structs.UnmarshalFromFile(testDir+"/"+f.Name(), &dt)
			if err != nil {
				t.Logf("couldn't set up database. Error reading in %v: %v", f.Name(), err.Error())
				t.FailNow()
			}

			_, err = CreateRoomConfiguration(dt)
			if err != nil {
				t.Logf("couldn't set up database. Error creating roomconfiguration %v: %v", f.Name(), err.Error())
				t.FailNow()
			}
		}
	}

	return func(tarp *testing.T) {
		os.Setenv("COUCH_ADDRESS", oldCouchAddress)
		os.Setenv("COUCH_USERNAME", oldCouchUsername)
		os.Setenv("COUCH_PASSWORD", oldCouchPassword)
		os.Setenv("LOGGING_FILE_LOCATION", oldLoggingLocation)
	}
}

func wipeDatabases() {
	MakeRequest("DELETE", "buildings", "", nil, nil)
	MakeRequest("DELETE", "rooms", "", nil, nil)
	MakeRequest("DELETE", "room_configurations", "", nil, nil)
	MakeRequest("DELETE", "devices", "", nil, nil)
	MakeRequest("DELETE", "device_types", "", nil, nil)

	MakeRequest("PUT", "buildings", "", nil, nil)
	MakeRequest("PUT", "rooms", "", nil, nil)
	MakeRequest("PUT", "room_configurations", "", nil, nil)
	MakeRequest("PUT", "devices", "", nil, nil)
	MakeRequest("PUT", "device_types", "", nil, nil)
}

func TestBuilding(t *testing.T) {
	defer setupDatabase(t)(t)

	t.Run("Building Create", testBuildingCreate)
	t.Run("Building Create Duplicate", testBuildingCreateDuplicate)
	t.Run("Building Update", testBuildingUpdate)
	t.Run("Building Delete", testBuildingDelete)
}

func testBuildingCreate(t *testing.T) {

	building := structs.Building{}
	//add a building
	err := structs.UnmarshalFromFile(testDir+"/new_building.json", &building)
	if err != nil {
		t.Logf("Error reading in %v: %v", "new_building.json", err.Error())
		t.Fail()
	}

	_, err = CreateBuilding(building)
	if err != nil {
		t.Logf("Error creating building %v: %v", "new_building.json", err.Error())
		t.Fail()
	}
}

func testBuildingCreateDuplicate(t *testing.T) {

	building := structs.Building{}
	//add a building
	err := structs.UnmarshalFromFile(testDir+"/setup_buildings_a.json", &building)
	if err != nil {
		t.Logf("Error reading in %v: %v", "setup_buildings_a.json", err.Error())
		t.Fail()
	}

	_, err = CreateBuilding(building)
	if err == nil {
		t.Logf("Creation succeeded when should have failed.")
		t.Fail()
	}
}

func testBuildingUpdate(t *testing.T) {

	building, err := GetBuildingByID("AAA")
	if err != nil {
		t.Logf("Couldn't get building: %v", err.Error())
		t.Fail()
	}

	currentlen := len(building.Tags)
	building.Tags = append(building.Tags, "pootingmonsterpenguins")
	newDescription := "No! Your great grandaughter had to be a CROSS DRESSER!"
	building.Description = newDescription
	rev := building.Rev

	building.Rev = ""

	//try to fail without rev
	_, err = CreateBuilding(building)
	if err == nil {
		t.Log("Succeeded when it shouldn't have. Failed on rev being null")
		t.FailNow()
	}

	building.Rev = rev

	b, err := CreateBuilding(building)
	if err != nil {
		t.Logf("Failed update: %v", err.Error())
		t.FailNow()
	}
	assert.Equal(t, b.Description, newDescription)
	assert.Equal(t, len(building.Tags), (currentlen + 1))

}

func testBuildingDelete(t *testing.T) {
	err := DeleteBuilding("BBB")
	assert.Nil(t, err)

	err = DeleteBuilding("ZZZ")
	assert.NotNil(t, err)
}
