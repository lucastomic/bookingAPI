package clientDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type ClientRepository struct {
	mysql.CommonMysqlLogic[domain.Client, int]
}

const findByReservationStmt string = "SELECT id,name,phone,passengers FROM client JOIN client_reservation ON client.id = client_reservation.client_id WHERE client_reservation.reservation_id = ?"

func NewClientRepository() ClientRepository {
	commonBehaivor := mysql.CommonMysqlLogic[domain.Client, int]{
		IPrimitiveRepoBehaivor: clientPrimitiveRepoBehaivor{},
	}
	return ClientRepository{commonBehaivor}
}

func (repo ClientRepository) FindByReservation(reservation domain.Reservation) ([]*domain.Client, error) {
	response, err := repo.Query(findByReservationStmt, []any{reservation.Id()})
	if err != nil {
		return []*domain.Client{}, err
	}
	return response, nil
}
