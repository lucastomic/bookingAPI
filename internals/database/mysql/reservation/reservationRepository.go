package reservationDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type ReservationRepository struct {
	mysql.CommonMysqlLogic[domain.Reservation, int]
}

const findByStateRoomStmt string = "SELECT * FROM reservation WHERE stateRoomId = ? AND boatId = ?"

func NewReservationRepository() ReservationRepository {
	commonBehaivor := newReservationCommonMysqlLogic()
	return ReservationRepository{commonBehaivor}
}
func (repo ReservationRepository) FindByStateRoom(stateRoom domain.StateRoom) ([]domain.Reservation, error) {
	response, err := repo.Query(findByStateRoomStmt, []any{stateRoom.Id(), stateRoom.BoatId()})
	if err != nil {
		return []domain.Reservation{}, err
	}
	return response, nil
}
