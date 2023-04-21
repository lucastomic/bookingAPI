package mysql

import (
	"database/sql"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
)

// Manages the interaction of the StateRoom entity with the database
type StateRoomsRepository struct {
}

const insertStateRoomStmt string = "INSERT INTO stateRoom(id, boatId) VALUES(?,?)"
const updateStateRoomStmt string = "UPDATE stateRoom SET id = ?, boatId =? WHERE id = ? AND boatId = ? "
const findStateRoomByIdStmt string = "SELECT id, boatId FROM stateRoom WHERE id = ? AND boatId = ? "

// insertStmt returns the statement to insert a new stateRoom
func (repo StateRoomsRepository) insertStmt() string {
	return insertStateRoomStmt
}

// updateStmt returns the statement to update a new stateRoom
func (repo StateRoomsRepository) updateStmt() string {
	return updateStateRoomStmt
}

// findByIdStmt returns the statement to findByIdStmt a new stateRoom
func (repo StateRoomsRepository) findByIdStmt() string {
	return findStateRoomByIdStmt
}

// persistenceValues returns an array with the fields of a stateRoom wihch will be
// persisted in the database
func (repo StateRoomsRepository) persistenceValues(stateRoom boat.StateRoom) []any {
	return []any{stateRoom.Id(), stateRoom.BoatId()}
}

// empty returns an empty stateRoom
func (repo StateRoomsRepository) empty() *boat.StateRoom {
	return boat.EmptyStateRoom()
}

// id returns the id of the stateRoom passed as argument
func (repo StateRoomsRepository) id(stateRoom boat.StateRoom) []int {
	return []int{stateRoom.BoatId(), stateRoom.Id()}
}

// isZero checks wether the stateRoom specified as paramter is a zero boat
func (repo StateRoomsRepository) isZero(stateRoom boat.StateRoom) bool {
	return stateRoom.ReservedDays() == nil && stateRoom.Id() == 0 && stateRoom.BoatId() == 0
}

func (repo StateRoomsRepository) scan(*sql.Rows) (boat.StateRoom, error) {
	//TODO:implement
}

// NewBoatRepository returns a BoatRepository
func NewStateRoomRepository() MysqlRepository[boat.StateRoom, int] {
	concretRepository := StateRoomsRepository{}
	return MysqlRepository[boat.StateRoom, int]{concretRepository}
}
