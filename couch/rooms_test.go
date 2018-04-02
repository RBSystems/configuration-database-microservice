package couch

import "testing"

func TestRoom(t *testing.T) {
	defer setupDatabase(t)(t)
}
