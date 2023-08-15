package reservationDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type ReservationRepository struct {
	mysql.CommonMysqlLogic[domain.Reservation, int]
}

const findByStateRoomStmt string = "SELECT * FROM reservation WHERE stateRoomId = ? AND boatId = ?"
const findByClientStmt string = "SELECT * FROM reservation JOIN client_reservation WHERE client_reservation.client_id = ?"

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

func (repo ReservationRepository) FindByStateRoom(stateRoom domain.StateRoom) ([]domain.Reservation, error) {
	response, err := repo.Query(findByStateRoomStmt, []any{stateRoom.Id(), stateRoom.BoatId()})
	if err != nil {
		return []domain.Reservation{}, err
	}
	return response, nil
}

func (repo ReservationRepository) FindByClient(client domain.Client) ([]domain.Reservation, error) {
	response, err := repo.Query(findByClientStmt, []any{client.Id()})
	if err != nil {
		return []domain.Reservation{}, err
	}
	return response, nil
}
