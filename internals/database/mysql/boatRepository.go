package mysql

import (
	"database/sql"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
)

// Manages the interaction of the boat entity with the database
type BoatRepository struct {
}

const insertBoatStmt string = "INSERT INTO boat(name) VALUES(?)"
const updateBoatStmt string = "UPDATE boat SET name = ? WHERE id = ?"
const findBoatByIdStmt string = "SELECT id, name FROM boat WHERE id = ?"

// insertStmt returns the statement to insert a new boat
func (b BoatRepository) insertStmt() string {
	return insertBoatStmt
}

// updateStmt returns the statement to update a new boat
func (b BoatRepository) updateStmt() string {
	return updateBoatStmt
}

// findByIdStmt returns the statement to findByIdStmt a new boat
func (b BoatRepository) findByIdStmt() string {
	return findBoatByIdStmt
}

// persistenceValues returns an array with the fields of a boat wihch will be
// persisted in the database
func (b BoatRepository) persistenceValues(boat boat.Boat) []any {
	return []any{boat.Name()}
}

// empty returns an empty boat
func (b BoatRepository) empty() *boat.Boat {
	return boat.EmtyBoat()
}

// id returns the id of the boat passed as argument
func (b BoatRepository) id(boat boat.Boat) []int {
	return []int{boat.Id()}
}

// isZero checks wether the boat specified as paramter is a zero boat
func (b BoatRepository) isZero(boat boat.Boat) bool {
	return boat.Name() == ""
}

// getScaneableEntity returns the a ScaneableBoat
func (repo BoatRepository) scan(*sql.Rows) (boat.Boat, error) {
	//TODO:implement
}

// NewBoatRepository returns a BoatRepository
func NewBoatRepository() MysqlRepository[boat.Boat, int] {
	concretRepository := BoatRepository{}
	return MysqlRepository[boat.Boat, int]{concretRepository}
}
