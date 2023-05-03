package stateRoomDB

import (
	"database/sql"

	reservationDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/reservation"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// stateRoomPrimitiveRepoBehaivor implements the behaivor needed for implementing a
// CommmonMysqlRepository[stateRoomPrimitiveRepoBehaivor,int]
type stateRoomPrimitiveRepoBehaivor struct {
}

const insertStateRoomStmt string = "INSERT INTO stateRoom(id, boatId) VALUES(?,?)"
const updateStateRoomStmt string = "UPDATE stateRoom SET id = ?, boatId =? WHERE id = ? AND boatId = ? "
const findStateRoomByIdStmt string = "SELECT FROM stateRoom WHERE id = ? AND boatId = ? "
const findAllstmt string = "SELECT * FROM stateRoom"
const removeStmt string = "DELETE FROM stateRoom WHERE id = ? AND boatId = ? "

// insertStmt returns the SQL statement to insert a new stateRoom
func (repo stateRoomPrimitiveRepoBehaivor) InsertStmt() string {
	return insertStateRoomStmt
}

// RemoveStmt returns the SQL statement to remove a stateRoom
func (repo stateRoomPrimitiveRepoBehaivor) RemoveStmt() string {
	return removeStmt
}

// updateStmt returns the SQL statement to update a new stateRoom
func (repo stateRoomPrimitiveRepoBehaivor) UpdateStmt() string {
	return updateStateRoomStmt
}

// findByIdStmt returns the SQL statement to find a stateRoom by id
func (repo stateRoomPrimitiveRepoBehaivor) FindByIdStmt() string {
	return findStateRoomByIdStmt
}

// findAllStmt returns the SQL statement to find all the staterooms
func (repo stateRoomPrimitiveRepoBehaivor) FindAllStmt() string {
	return findAllstmt
}

// persistenceValues returns an array with the fields of a stateRoom wihch will be
// persisted in the database
func (repo stateRoomPrimitiveRepoBehaivor) PersistenceValues(stateRoom domain.StateRoom) []any {
	return []any{stateRoom.Id(), stateRoom.BoatId()}
}

// empty returns an empty stateRoom
func (repo stateRoomPrimitiveRepoBehaivor) Empty() *domain.StateRoom {
	return domain.EmptyStateRoom()
}

// id returns the id of the stateRoom passed as argument
func (repo stateRoomPrimitiveRepoBehaivor) Id(stateRoom domain.StateRoom) []int {
	return []int{stateRoom.BoatId(), stateRoom.Id()}
}

// isZero checks wether the stateRoom specified as paramter is a zero boat
func (repo stateRoomPrimitiveRepoBehaivor) IsZero(stateRoom domain.StateRoom) bool {
	return stateRoom.ReservedDays() == nil && stateRoom.Id() == 0 && stateRoom.BoatId() == 0
}

// scan scans the stateRoom inside the row passed by argument
func (repo stateRoomPrimitiveRepoBehaivor) Scan(row *sql.Rows) (domain.StateRoom, error) {
	var id int
	var boatId int

	var reservedDays []domain.Reservation = []domain.Reservation{}
	err := row.Scan(&id, &boatId)
	if err != nil {
		return *domain.EmptyStateRoom(), err
	}

	return *domain.NewStateRoom(id, boatId, reservedDays), nil
}

func (repo stateRoomPrimitiveRepoBehaivor) UpdateRelations(stateRoom *domain.StateRoom) error {
	reservationRepo := reservationDB.NewReservationRepository()
	stateRoomReservations, err := reservationRepo.FindByStateRoom(*stateRoom)
	if err != nil {
		return err
	}
	stateRoom.SetReservedDays(stateRoomReservations)
	return nil
}
