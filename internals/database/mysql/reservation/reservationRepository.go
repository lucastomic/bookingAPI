package reservationDB

import (
	"time"

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
	var response []domain.Reservation
	db := mysql.GetInstance()

	stmt, err := db.Prepare(findByStateRoomStmt)
	if err != nil {
		return []domain.Reservation{*domain.EmptyReservation()}, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(stateRoom.Id(), stateRoom.BoatId())
	if err != nil {
		return []domain.Reservation{}, err
	}

	for rows.Next() {
		var id int
		var name string
		var phone string
		var firstDay time.Time
		var lastDay time.Time
		var boatId int
		var stateRoomId int
		err := rows.Scan(&id, &name, &phone, &firstDay, &lastDay, &boatId, &stateRoomId)
		if err != nil {
			return []domain.Reservation{}, err
		}

		user := domain.NewUser(name, phone)
		response = append(response, *domain.NewReservation(id, user, firstDay, lastDay, boatId, stateRoomId))
	}
	return response, nil
}
