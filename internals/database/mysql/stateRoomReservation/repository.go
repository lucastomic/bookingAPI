package stateroomreservation

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

const insertClientHasReservationStmt string = "INSERT INTO stateRoom_reservation(stateroom_id, reservation_id, boat_id) VALUES(?,?,?)"

type Repository struct {
}

func (repo Repository) Save(stateRoom domain.StateRoom, reservation domain.Reservation) error {
	err := mysql.ExecStmt(insertClientHasReservationStmt, []any{stateRoom.Id(), reservation.Id(), stateRoom.BoatId()})
	if err != nil {
		return err
	}
	return nil
}
