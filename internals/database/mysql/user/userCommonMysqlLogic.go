package userDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

func commonMysqlLogicForUser() mysql.CommonMysqlLogic[domain.User, string] {
	return mysql.CommonMysqlLogic[domain.User, string]{
		IPrimitiveRepoBehaivor: userPrimitiveRepoBehaivor{},
	}
}
