package boatDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// BoatRepository manages all the interactions between the Boat entity and the database
type BoatRepository struct {
	mysql.CommonMysqlLogic[domain.Boat, int]
}

func NewBoatRepository() BoatRepository {
	commonBehaivor := newBoatCommonMysqlLogic()
	return BoatRepository{commonBehaivor}
}
