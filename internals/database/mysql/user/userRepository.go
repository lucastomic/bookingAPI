package userDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type UserRepository struct {
	mysql.CommonMysqlLogic[domain.User, string]
}

func NewUserRepository(boatRepo databaseport.IBoatRepository) UserRepository {
	commonBehaivor := mysql.CommonMysqlLogic[domain.User, string]{
		IPrimitiveRepoBehaivor: userPrimitiveRepoBehaivor{boatRepo},
	}
	return UserRepository{commonBehaivor}
}
