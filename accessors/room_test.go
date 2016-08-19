package accessors

// import (
// 	"fmt"
// 	"testing"
//
// 	"gopkg.in/DATA-DOG/go-sqlmock.v1"
// )
//
// func TestGetAllRooms(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM buildings`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "shortname"}).AddRow(1, "Information Technology Building", "ITB").AddRow(2, "Crabtree Building", "CTB"))
// 	mock.ExpectQuery(`SELECT \* FROM rooms`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "vlan", "building"}).AddRow(1, "1100A", 600, 1).AddRow(2, "1000", 650, 2))
//
// 	_, err = accessorGroup.GetAllRooms()
// 	if err != nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
//
// func TestGetAllRoomsFail(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM buildings`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "shortname"}).AddRow(1, "Information Technology Building", "ITB").AddRow(2, "Crabtree Building", "CTB"))
// 	mock.ExpectQuery(`SELECT \* FROM rooms`).WillReturnError(fmt.Errorf("ERROR"))
//
// 	_, err = accessorGroup.GetAllRooms()
// 	if err == nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
//
// func TestGetRoomByID(test *testing.T) {
// 	rows := sqlmock.NewRows([]string{"id", "name", "building", "vlan"}).
// 		AddRow(1, "1100A", 1, 600)
//
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM rooms WHERE id=\?`).WillReturnRows(rows)
//
// 	_, err = accessorGroup.GetRoomByID(1)
// 	if err != nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
//
// func TestGetRoomByIDFail(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM rooms WHERE id=\?`).WillReturnError(fmt.Errorf("ERROR"))
//
// 	_, err = accessorGroup.GetRoomByID(1)
// 	if err == nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
//
// func TestGetRoomsByBuilding(test *testing.T) {
// 	rows := sqlmock.NewRows([]string{"id", "name", "building", "vlan"}).
// 		AddRow(1, "1100A", 1, 600).
// 		AddRow(2, "1100B", 1, 650)
//
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM rooms WHERE building=\?`).WillReturnRows(rows)
//
// 	_, err = accessorGroup.GetRoomsByBuilding(1)
// 	if err != nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
//
// func TestGetRoomsByBuildingFail(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM rooms WHERE building=\?`).WillReturnError(fmt.Errorf("ERROR"))
//
// 	_, err = accessorGroup.GetRoomsByBuilding(1)
// 	if err == nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
//
// func TestGetRoomByBuildingAndName(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM buildings WHERE shortname=\?`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "shortname"}).AddRow(1, "Information Technology Building", "ITB"))
// 	mock.ExpectQuery(`SELECT \* FROM rooms WHERE building=\? AND name=\?`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "building", "vlan"}).AddRow(1, "1100A", 1, 600))
//
// 	_, err = accessorGroup.GetRoomByBuildingAndName("ITB", "1100A")
// 	if err != nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
//
// func TestGetRoomByBuildingAndNameFail(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM buildings WHERE shortname=\?`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "shortname"}).AddRow(1, "Information Technology Building", "ITB"))
// 	mock.ExpectQuery(`SELECT \* FROM rooms WHERE building=\? AND name=\?`).WillReturnError(fmt.Errorf("ERROR"))
//
// 	_, err = accessorGroup.GetRoomByBuildingAndName("ITB", "1100A")
// 	if err == nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
//
// func TestMakeRoom(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM buildings WHERE shortname=\?`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "shortname"}).AddRow(1, "Information Technology Building", "ITB"))
// 	mock.ExpectExec(`INSERT INTO rooms \(name, building, vlan\) VALUES \(\?, \?, \?\)`).WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectQuery(`SELECT \* FROM buildings WHERE shortname=\?`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "shortname"}).AddRow(1, "Information Technology Building", "ITB"))
// 	mock.ExpectQuery(`SELECT \* FROM rooms WHERE building=\? AND name=\?`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "building", "vlan"}).AddRow(1, "1100A", 1, 600))
//
// 	_, err = accessorGroup.MakeRoom("1100A", "ITB", 600)
// 	if err != nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
//
// func TestMakeRoomFail(test *testing.T) { // FIX IN A MINUTE
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM buildings WHERE shortname=\?`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "shortname"}).AddRow(1, "Information Technology Building", "ITB"))
// 	mock.ExpectExec(`INSERT INTO rooms \(name, building, vlan\) VALUES \(\?, \?, \?\)`).WillReturnError(fmt.Errorf("ERROR"))
//
// 	_, err = accessorGroup.MakeRoom("1100A", "ITB", 600)
// 	if err == nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
