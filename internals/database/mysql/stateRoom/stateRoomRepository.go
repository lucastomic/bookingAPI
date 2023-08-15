package stateRoomDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// StateRoomRepository manages all the interactions between the StateRoom entity and the database
type StateRoomRepository struct {
	mysql.CommonMysqlLogic[domain.StateRoom, int]
}

const findStateRoomByBoatIDStmt string = "SELECT * FROM stateRoom WHERE boatId = ? "
const findByReservationIDStmt string = "SELECT * FROM stateRoom JOIN stateRoom_reservation ON stateRoom.id = stateRoom_reservation.stateroom_id AND stateRoom.boatId = stateRoom_reservation.boat_id WHERE = stateRoom_reservation.reservation_id = ?"

func NewStateRoomRepository(
	reservationRepo databaseport.IReservationRepository,
	stateroomReservationRepo databaseport.RelationSaver[domain.StateRoom, domain.Reservation],
) StateRoomRepository {
	commonBehaivor := mysql.CommonMysqlLogic[domain.StateRoom, int]{
		stateRoomPrimitiveRepoBehaivor{reservationRepo, stateroomReservationRepo},
	}
	return StateRoomRepository{commonBehaivor}
}
func (repo StateRoomRepository) FindByBoatId(boatId int) ([]domain.StateRoom, error) {
	response, err := repo.Query(findStateRoomByBoatIDStmt, []any{boatId})
	if err != nil {
		return []domain.StateRoom{}, err
	}
	return response, nil
}
func (repo StateRoomRepository) FindByReservation(reservation domain.Reservation) ([]domain.StateRoom, error) {
	response, err := repo.Query(findByReservationIDStmt, []any{reservation.Id()})
	if err != nil {
		return []domain.StateRoom{}, err
	}
	return response, nil
}
