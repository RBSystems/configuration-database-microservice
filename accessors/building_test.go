package accessors

// import (
// 	"fmt"
// 	"testing"
//
// 	"gopkg.in/DATA-DOG/go-sqlmock.v1"
// )
//
// func TestGetAllBuildings(test *testing.T) {
// 	rows := sqlmock.NewRows([]string{"id", "name", "shortname"}).
// 		AddRow(1, "Information Technology Building", "ITB").
// 		AddRow(2, "Crabtree Building", "CTB")
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
// 	mock.ExpectQuery(`SELECT \* FROM buildings`).WillReturnRows(rows)
//
// 	_, err = accessorGroup.GetAllBuildings()
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
// func TestGetAllBuildingsFail(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM buildings`).WillReturnError(fmt.Errorf("ERROR"))
//
// 	_, err = accessorGroup.GetAllBuildings()
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
// func TestGetBuildingByID(test *testing.T) {
// 	rows := sqlmock.NewRows([]string{"id", "name", "shortname"}).
// 		AddRow(1, "Information Technology Building", "ITB")
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
// 	mock.ExpectQuery(`SELECT \* FROM buildings WHERE id=\?`).WillReturnRows(rows)
//
// 	_, err = accessorGroup.GetBuildingByID(1)
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
// func TestGetBuildingByIDFail(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM buildings WHERE id=\?`).WillReturnError(fmt.Errorf("ERROR"))
//
// 	_, err = accessorGroup.GetBuildingByID(1)
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
// func TestGetBuildingByShortname(test *testing.T) {
// 	rows := sqlmock.NewRows([]string{"id", "name", "shortname"}).
// 		AddRow(1, "Information Technology Building", "ITB")
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
// 	mock.ExpectQuery(`SELECT \* FROM buildings WHERE shortname=\?`).WillReturnRows(rows)
//
// 	_, err = accessorGroup.GetBuildingByShortname("ITB")
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
// func TestGetBuildingByShortnameFail(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectQuery(`SELECT \* FROM buildings WHERE shortname=\?`).WillReturnError(fmt.Errorf("ERROR"))
//
// 	_, err = accessorGroup.GetBuildingByShortname("ITB")
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
// func TestMakeBuilding(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectExec(`INSERT INTO buildings \(name, shortname\) VALUES \(\?, \?\)`).WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectQuery(`SELECT \* FROM buildings WHERE shortname=\?`).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "shortname"}).AddRow(1, "Information Technology Building", "ITB"))
//
// 	_, err = accessorGroup.MakeBuilding("Information Technology Building", "ITB")
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
// func TestMakeBuildingFail(test *testing.T) {
// 	database, mock, err := sqlmock.New()
// 	if err != nil {
// 		test.Fatal(err)
// 	}
//
// 	// Constructs a new accessor group and connects it to the mock database
// 	accessorGroup := new(AccessorGroup)
// 	accessorGroup.Database = database
//
// 	mock.ExpectExec(`INSERT INTO buildings \(name, shortname\) VALUES \(\?, \?\)`).WillReturnError(fmt.Errorf("ERROR"))
//
// 	_, err = accessorGroup.MakeBuilding("Information Technology Building", "ITB")
// 	if err == nil {
// 		test.Error(err)
// 	}
//
// 	err = mock.ExpectationsWereMet()
// 	if err != nil {
// 		test.Error(err)
// 	}
// }
