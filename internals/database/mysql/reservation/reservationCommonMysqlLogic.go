package reservationDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

func newReservationCommonMysqlLogic() mysql.CommonMysqlLogic[domain.Reservation, int] {
	return mysql.CommonMysqlLogic[domain.Reservation, int]{
		IPrimitiveRepoBehaivor: reservationPrimitiveRepoBehaivor{},
	}
}
