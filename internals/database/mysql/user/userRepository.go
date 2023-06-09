package userDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type UserRepository struct {
	mysql.CommonMysqlLogic[domain.User, string]
}

func NewUserRepository() UserRepository {
	commonBehaivor := commonMysqlLogicForUser()
	return UserRepository{commonBehaivor}
}
