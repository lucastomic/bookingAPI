package stateRoomDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// StateRoomRepository manages all the interactions between the StateRoom entity and the database
type StateRoomRepository struct {
	mysql.CommonMysqlLogic[domain.StateRoom, int]
}

const findStateRoomByBoatIDStmt string = "SELECT id FROM stateRoom WHERE boatId = ? "

func NewStateRoomRepository() StateRoomRepository {
	commonBehaivor := commonMysqlLogicForStateRoom()
	return StateRoomRepository{commonBehaivor}
}
func (repo StateRoomRepository) FindByBoatId(boatId int) ([]domain.StateRoom, error) {
	//TODO: MOVE ALL THIS LOGIC TO A METHOD QUERY() IN CommonMysqlLogic TO BE ABLE TO REUSE THE scan() METHOD
	var response []domain.StateRoom

	db := mysql.GetInstance()

	stmt, err := db.Prepare(findStateRoomByBoatIDStmt)
	if err != nil {
		return []domain.StateRoom{*domain.EmptyStateRoom()}, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(boatId)
	if err != nil {
		return []domain.StateRoom{*domain.EmptyStateRoom()}, err
	}

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return []domain.StateRoom{*domain.EmptyStateRoom()}, err
		}

		stateRoom := domain.NewStateRoom(id, boatId, []domain.Reservation{})
		repo.UpdateRelations(stateRoom)
		response = append(response, *stateRoom)
	}
	return response, nil
}
