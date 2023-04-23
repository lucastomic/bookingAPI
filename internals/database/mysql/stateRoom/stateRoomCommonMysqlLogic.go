package stateRoomDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

func commonMysqlLogicForStateRoom() mysql.CommonMysqlLogic[domain.StateRoom, int] {
	return mysql.CommonMysqlLogic[domain.StateRoom, int]{
		IPrimitiveRepoBehaivor: stateRoomPrimitiveRepoBehaivor{},
	}
}
