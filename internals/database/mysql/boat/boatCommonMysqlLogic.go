package boatDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

func newBoatCommonMysqlLogic() mysql.CommonMysqlLogic[domain.Boat, int] {
	return mysql.CommonMysqlLogic[domain.Boat, int]{
		IPrimitiveRepoBehaivor: boatPrimitiveRepoBehaivor{},
	}
}
