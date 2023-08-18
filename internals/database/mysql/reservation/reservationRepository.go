package reservationDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type ReservationRepository struct {
	mysql.CommonMysqlLogic[domain.Reservation, int]
}

const findByStateRoomStmt string = "SELECT id,firstDay,lastDay,isOpen,boatId FROM reservation JOIN stateRoom_reservation ON stateRoom_reservation.reservation_id = reservation.id WHERE stateRoom_reservation.stateroom_id = ? AND stateRoom_reservation.boat_id = ?"
const findByClientStmt string = "SELECT id,firstDay,lastDay,isOpen,boatId FROM reservation JOIN client_reservation ON reservation.id = client_reservation.reservation_id WHERE client_reservation.client_id = ?"

func NewReservationRepository(
	clientRepo databaseport.IClientRepository,
	clientReservationRepo databaseport.RelationSaver[domain.Client, domain.Reservation],
) ReservationRepository {
	commonBehaivor := mysql.CommonMysqlLogic[domain.Reservation, int]{
		IPrimitiveRepoBehaivor: reservationPrimitiveRepoBehaivor{
			clientRepository:      clientRepo,
			clientReservationRepo: clientReservationRepo,
		},
	}
	return ReservationRepository{commonBehaivor}
}

func (repo ReservationRepository) FindByStateRoom(stateRoom domain.StateRoom) ([]*domain.Reservation, error) {
	response, err := repo.Query(findByStateRoomStmt, []any{stateRoom.Id(), stateRoom.BoatId()})
	if err != nil {
		return []*domain.Reservation{}, err
	}
	return response, nil
}

func (repo ReservationRepository) FindByClient(client domain.Client) ([]*domain.Reservation, error) {
	response, err := repo.Query(findByClientStmt, []any{client.Id()})
	if err != nil {
		return []*domain.Reservation{}, err
	}
	return response, nil
}
