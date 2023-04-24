package stateRoomDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// StateRoomRepository manages all the interactions between the StateRoom entity and the database
type StateRoomRepository struct {
	mysql.CommonMysqlLogic[domain.StateRoom, int]
}

const findStateRoomByBoatIDStmt string = "SELECT * FROM stateRoom WHERE boatId = ? "

func NewStateRoomRepository() StateRoomRepository {
	commonBehaivor := commonMysqlLogicForStateRoom()
	return StateRoomRepository{commonBehaivor}
}
func (repo StateRoomRepository) FindByBoatId(boatId int) ([]domain.StateRoom, error) {
	response, err := repo.Query(findStateRoomByBoatIDStmt, []any{boatId})
	if err != nil {
		return []domain.StateRoom{}, err
	}
	return response, nil
}
