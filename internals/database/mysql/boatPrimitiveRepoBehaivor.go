package mysql

import (
	"database/sql"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
)

// boatPrimitiveRepoBehaivor implements the behaivor needed for implementing a CommmonMysqlRepository[boatPrimitiveRepoBehaivor,int]
type boatPrimitiveRepoBehaivor struct {
}

const insertBoatStmt string = "INSERT INTO boat(name) VALUES(?)"
const updateBoatStmt string = "UPDATE boat SET name = ? WHERE id = ?"
const findBoatByIdStmt string = "SELECT id, name FROM boat WHERE id = ?"

// insertStmt returns the statement to insert a new boat
func (b boatPrimitiveRepoBehaivor) insertStmt() string {
	return insertBoatStmt
}

// updateStmt returns the statement to update a new boat
func (b boatPrimitiveRepoBehaivor) updateStmt() string {
	return updateBoatStmt
}

// findByIdStmt returns the statement to findByIdStmt a new boat
func (b boatPrimitiveRepoBehaivor) findByIdStmt() string {
	return findBoatByIdStmt
}

// persistenceValues returns an array with the fields of a boat wihch will be
// persisted in the database
func (b boatPrimitiveRepoBehaivor) persistenceValues(boat boat.Boat) []any {
	return []any{boat.Name()}
}

// empty returns an empty boat
func (b boatPrimitiveRepoBehaivor) empty() *boat.Boat {
	return boat.EmtyBoat()
}

// id returns the id of the boat passed as argument
func (b boatPrimitiveRepoBehaivor) id(boat boat.Boat) []int {
	return []int{boat.Id()}
}

// isZero checks wether the boat specified as paramter is a zero boat
func (b boatPrimitiveRepoBehaivor) isZero(boat boat.Boat) bool {
	return boat.Name() == ""
}

// scan scans the boat inside the row passed by argument
func (repo boatPrimitiveRepoBehaivor) scan(row *sql.Rows) (boat.Boat, error) {
	var id int
	var name string
	var stateRooms []boat.StateRoom = []boat.StateRoom{}
	err := row.Scan(&id, &name)
	if err != nil {
		return *boat.EmtyBoat(), err
	}
	return *boat.NewBoatWithId(id, name, stateRooms), nil
}

func (repo boatPrimitiveRepoBehaivor) updateRelations(boat *boat.Boat) error {
	stateRoomRepo := NewStateRoomRepository()
	boatStateRooms, err := stateRoomRepo.FindByBoatId(boat.Id())
	if err != nil {
		return err
	}
	boat.SetStateRooms(boatStateRooms)
	return nil
}
