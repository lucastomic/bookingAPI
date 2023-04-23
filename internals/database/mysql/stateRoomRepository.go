package mysql

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/reservation"
)

// StateRoomRepository manages all the interactions between the StateRoom entity and the database
type StateRoomRepository struct {
	CommonMysqlRepository[boat.StateRoom, int]
}

const findStateRoomByBoatIDStmt string = "SELECT id FROM stateRoom WHERE boatId = ? "

func NewStateRoomRepository() StateRoomRepository {
	commonBehaivor := newStateRoomCommonMysqlRepository()
	return StateRoomRepository{commonBehaivor}
}
func (repo StateRoomRepository) FindByBoatId(boatId int) ([]boat.StateRoom, error) {
	var response []boat.StateRoom
	db := getInstance()

	stmt, err := db.Prepare(findStateRoomByBoatIDStmt)
	if err != nil {
		return []boat.StateRoom{*boat.EmptyStateRoom()}, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(boatId)
	if err != nil {
		return []boat.StateRoom{*boat.EmptyStateRoom()}, err
	}

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return []boat.StateRoom{*boat.EmptyStateRoom()}, err
		}
		response = append(response, *boat.NewStateRoom(id, boatId, []reservation.Reservation{}))
	}
	return response, nil
}
