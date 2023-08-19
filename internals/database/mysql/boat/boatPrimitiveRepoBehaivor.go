package boatDB

import (
	"database/sql"

	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// boatPrimitiveRepoBehaivor implements the behaivor needed for implementing a CommmonMysqlRepository[boatPrimitiveRepoBehaivor,int]
type boatPrimitiveRepoBehaivor struct {
	stateRoomRepo databaseport.IStateRoomRepository
}

const insertBoatStmt string = "INSERT INTO boat(name,max_capacity,owner) VALUES(?,?,?)"
const updateBoatStmt string = "UPDATE boat SET name = ?, max_capacity = ?, owner = ? WHERE id = ?"
const findBoatByIdStmt string = "SELECT * FROM boat WHERE id = ?"
const findAllStmt string = "SELECT * FROM boat"
const removeStmt string = "DELETE FROM boat WHERE id = ?"

// insertStmt returns the statement to insert a new boat
func (b boatPrimitiveRepoBehaivor) InsertStmt() string {
	return insertBoatStmt
}

// RemoveStmt returns rhe statement to remove a boat
func (b boatPrimitiveRepoBehaivor) RemoveStmt() string {
	return removeStmt
}

// updateStmt returns the statement to update a new boat
func (b boatPrimitiveRepoBehaivor) UpdateStmt() string {
	return updateBoatStmt
}

// findByIdStmt returns the statement to findByIdStmt a new boat
func (b boatPrimitiveRepoBehaivor) FindByIdStmt() string {
	return findBoatByIdStmt
}

// findByIdStmt returns the statement to findByIdStmt a new boat
func (b boatPrimitiveRepoBehaivor) FindAllStmt() string {
	return findAllStmt
}

// persistenceValues returns an array with the fields of a boat wihch will be
// persisted in the database
func (b boatPrimitiveRepoBehaivor) PersistenceValues(boat domain.Boat) []any {
	return []any{boat.Name(), boat.MaxCapacity(), boat.Owner()}
}

// empty returns an empty boat
func (b boatPrimitiveRepoBehaivor) Empty() *domain.Boat {
	return domain.EmptyBoat()
}

// id returns the id of the boat passed as argument
func (b boatPrimitiveRepoBehaivor) Id(boat domain.Boat) []int {
	return []int{boat.Id()}
}

func (repo boatPrimitiveRepoBehaivor) ModifyId(boat *domain.Boat, id int64) {
	boat.SetId(int(id))
}

// isZero checks wether the boat specified as paramter is a zero boat
func (b boatPrimitiveRepoBehaivor) IsZero(boat domain.Boat) bool {
	return boat.Name() == ""
}

// scan scans the boat inside the row passed by argument
func (repo boatPrimitiveRepoBehaivor) Scan(row *sql.Rows) (domain.Boat, error) {
	var id, maxCapacity int
	var name, owner string
	var stateRooms []*domain.StateRoom = []*domain.StateRoom{}
	err := row.Scan(&id, &name, &owner, &maxCapacity)
	if err != nil {
		return *domain.EmptyBoat(), err
	}
	return *domain.NewBoatWithId(id, name, stateRooms, owner, maxCapacity), nil
}

func (repo boatPrimitiveRepoBehaivor) UpdateRelations(boat *domain.Boat) error {
	boatStateRooms, err := repo.stateRoomRepo.FindByBoatId(boat.Id())
	if err != nil {
		return err
	}
	boat.SetStateRooms(boatStateRooms)
	return nil
}

// SaveChildsChanges takes all the staterooms in the boat and save their changes in the datanase (or
// inserts a new stateroom if it's a new one)
func (repo boatPrimitiveRepoBehaivor) SaveChildsChanges(boat *domain.Boat) error {
	for _, stateRoom := range boat.StateRooms() {
		err := repo.stateRoomRepo.Save(stateRoom)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo boatPrimitiveRepoBehaivor) SaveRelations(boat *domain.Boat) error {
	return nil
}
