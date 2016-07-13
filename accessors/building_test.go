package accessors

import (
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetAllBuildings(test *testing.T) {
	rows := sqlmock.NewRows([]string{"id", "name", "shortname"}).
		AddRow(1, "Information Technology Building", "ITB").
		AddRow(2, "Crabtree Building", "CTB")

	database, mock, err := sqlmock.New()
	if err != nil {
		test.Fatal(err)
	}

	// Constructs a new accessor group and connects it to the mock database
	accessorGroup := new(AccessorGroup)
	accessorGroup.Database = database

	mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(rows)

	buildings, err := accessorGroup.GetAllBuildings()
	if err != nil {
		test.Error(err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		test.Error(err)
	}
}
