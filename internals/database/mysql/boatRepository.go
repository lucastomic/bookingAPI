package mysql

import "github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"

// BoatRepository manages all the interactions between the Boat entity and the database
type BoatRepository struct {
	CommonMysqlRepository[boat.Boat, int]
}

func NewBoatRepository() BoatRepository {
	commonBehaivor := newBoatCommonMysqlRepository()
	return BoatRepository{commonBehaivor}
}
