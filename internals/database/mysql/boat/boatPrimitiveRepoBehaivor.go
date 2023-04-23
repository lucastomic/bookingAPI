package boatDB

import (
	"database/sql"

	stateRoomDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/stateRoom"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// boatPrimitiveRepoBehaivor implements the behaivor needed for implementing a CommmonMysqlRepository[boatPrimitiveRepoBehaivor,int]
type boatPrimitiveRepoBehaivor struct {
}

const insertBoatStmt string = "INSERT INTO boat(name) VALUES(?)"
const updateBoatStmt string = "UPDATE boat SET name = ? WHERE id = ?"
const findBoatByIdStmt string = "SELECT id, name FROM boat WHERE id = ?"

// insertStmt returns the statement to insert a new boat
func (b boatPrimitiveRepoBehaivor) InsertStmt() string {
	return insertBoatStmt
}

// updateStmt returns the statement to update a new boat
func (b boatPrimitiveRepoBehaivor) UpdateStmt() string {
	return updateBoatStmt
}

// findByIdStmt returns the statement to findByIdStmt a new boat
func (b boatPrimitiveRepoBehaivor) FindByIdStmt() string {
	return findBoatByIdStmt
}

// persistenceValues returns an array with the fields of a boat wihch will be
// persisted in the database
func (b boatPrimitiveRepoBehaivor) PersistenceValues(boat domain.Boat) []any {
	return []any{boat.Name()}
}

// empty returns an empty boat
func (b boatPrimitiveRepoBehaivor) Empty() *domain.Boat {
	return domain.EmtyBoat()
}

// id returns the id of the boat passed as argument
func (b boatPrimitiveRepoBehaivor) Id(boat domain.Boat) []int {
	return []int{boat.Id()}
}

// isZero checks wether the boat specified as paramter is a zero boat
func (b boatPrimitiveRepoBehaivor) IsZero(boat domain.Boat) bool {
	return boat.Name() == ""
}

// scan scans the boat inside the row passed by argument
func (repo boatPrimitiveRepoBehaivor) Scan(row *sql.Rows) (domain.Boat, error) {
	var id int
	var name string
	var stateRooms []domain.StateRoom = []domain.StateRoom{}
	err := row.Scan(&id, &name)
	if err != nil {
		return *domain.EmtyBoat(), err
	}
	return *domain.NewBoatWithId(id, name, stateRooms), nil
}

func (repo boatPrimitiveRepoBehaivor) UpdateRelations(boat *domain.Boat) error {
	stateRoomRepo := stateRoomDB.NewStateRoomRepository()
	boatStateRooms, err := stateRoomRepo.FindByBoatId(boat.Id())
	if err != nil {
		return err
	}
	boat.SetStateRooms(boatStateRooms)
	return nil
}
