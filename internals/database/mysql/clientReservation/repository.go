package clientreservation

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

const insertClientHasReservationStmt string = "INSERT IGNORE INTO client_reservation(client_id, reservation_id) VALUES(?,?)"

type Repository struct {
}

func (repo Repository) Save(client domain.Client, reservation domain.Reservation) error {
	err := mysql.ExecStmt(insertClientHasReservationStmt, []any{client.Id(), reservation.Id()})
	if err != nil {
		return err
	}
	return nil
}
