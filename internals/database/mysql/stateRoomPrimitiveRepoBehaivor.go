package mysql

import (
	"database/sql"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/reservation"
)

// stateRoomPrimitiveRepoBehaivor implements the behaivor needed for implementing a
// CommmonMysqlRepository[stateRoomPrimitiveRepoBehaivor,int]
type stateRoomPrimitiveRepoBehaivor struct {
}

const insertStateRoomStmt string = "INSERT INTO stateRoom(id, boatId) VALUES(?,?)"
const updateStateRoomStmt string = "UPDATE stateRoom SET id = ?, boatId =? WHERE id = ? AND boatId = ? "
const findStateRoomByIdStmt string = "SELECT id, boatId FROM stateRoom WHERE id = ? AND boatId = ? "

// insertStmt returns the statement to insert a new stateRoom
func (repo stateRoomPrimitiveRepoBehaivor) insertStmt() string {
	return insertStateRoomStmt
}

// updateStmt returns the statement to update a new stateRoom
func (repo stateRoomPrimitiveRepoBehaivor) updateStmt() string {
	return updateStateRoomStmt
}

// findByIdStmt returns the statement to findByIdStmt a new stateRoom
func (repo stateRoomPrimitiveRepoBehaivor) findByIdStmt() string {
	return findStateRoomByIdStmt
}

// persistenceValues returns an array with the fields of a stateRoom wihch will be
// persisted in the database
func (repo stateRoomPrimitiveRepoBehaivor) persistenceValues(stateRoom boat.StateRoom) []any {
	return []any{stateRoom.Id(), stateRoom.BoatId()}
}

// empty returns an empty stateRoom
func (repo stateRoomPrimitiveRepoBehaivor) empty() *boat.StateRoom {
	return boat.EmptyStateRoom()
}

// id returns the id of the stateRoom passed as argument
func (repo stateRoomPrimitiveRepoBehaivor) id(stateRoom boat.StateRoom) []int {
	return []int{stateRoom.BoatId(), stateRoom.Id()}
}

// isZero checks wether the stateRoom specified as paramter is a zero boat
func (repo stateRoomPrimitiveRepoBehaivor) isZero(stateRoom boat.StateRoom) bool {
	return stateRoom.ReservedDays() == nil && stateRoom.Id() == 0 && stateRoom.BoatId() == 0
}

// scan scans the stateRoom inside the row passed by argument
func (repo stateRoomPrimitiveRepoBehaivor) scan(row *sql.Rows) (boat.StateRoom, error) {
	var id int
	var boatId int

	var reservedDays []reservation.Reservation = []reservation.Reservation{}
	err := row.Scan(&id, &boatId)
	if err != nil {
		return *boat.EmptyStateRoom(), nil
	}

	return *boat.NewStateRoom(id, boatId, reservedDays), nil
}

func (repo stateRoomPrimitiveRepoBehaivor) updateRelations(boat *boat.StateRoom) error {
	return nil
}
