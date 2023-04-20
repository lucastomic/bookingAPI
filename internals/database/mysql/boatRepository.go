package mysql

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
)

type BoatRepository struct {
}

const insertBoatStmt string = "INSERT INTO boat(name) VALUES(?)"
const updateBoatStmt string = "UPDATE boat SET name = ? WHERE id = ?"
const findBoatByIdStmt string = "SELECT name FROM boat WHERE id = ?"

func (b BoatRepository) insertStmt() string {
	return insertBoatStmt
}

func (b BoatRepository) updateStmt() string {
	return updateBoatStmt
}

func (b BoatRepository) findByIdStmt() string {
	return findBoatByIdStmt
}

func (b BoatRepository) persistenceValues(boat boat.Boat) []any {
	return []any{boat.Name}
}

func (b BoatRepository) empty() *boat.Boat {
	return boat.EmtyBoat()
}

func (b BoatRepository) id(boat boat.Boat) int {
	return boat.Id()
}

func NewBoatRepository() MysqlRepository[boat.Boat, int] {
	concretRepository := BoatRepository{}
	return MysqlRepository[boat.Boat, int]{concretRepository}
}
