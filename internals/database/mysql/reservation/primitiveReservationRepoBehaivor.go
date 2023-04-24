package reservationDB

import (
	"database/sql"
	"time"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type reservationPrimitiveRepoBehaivor struct {
}

const insertStmt string = "INSERT INTO reservation(name, phone,firstDay,lastDay,boatId,stateRoomId) VALUES(?,?,?,?,?,?)"
const updateStmt string = "UPDATE reservation SET name = ?,phone = ?, firstDay = ?, lastDay = ?, boatId = ?, stateRoomId = ? WHERE id = ?"
const findByIdStmt string = "SELECT * FROM reservation WHERE id = ?"

func (repo reservationPrimitiveRepoBehaivor) InsertStmt() string {
	return insertStmt
}
func (repo reservationPrimitiveRepoBehaivor) UpdateStmt() string {
	return updateStmt
}
func (repo reservationPrimitiveRepoBehaivor) FindByIdStmt() string {
	return findByIdStmt
}
func (repo reservationPrimitiveRepoBehaivor) PersistenceValues(reservation domain.Reservation) []any {
	return []any{reservation.UserName(), reservation.UserPhone(), reservation.FirstDay(), reservation.LastDay(), reservation.BoatId(), reservation.StateRoomId()}
}
func (repo reservationPrimitiveRepoBehaivor) Id(reservation domain.Reservation) []int {
	return []int{reservation.Id()}
}

func (repo reservationPrimitiveRepoBehaivor) Empty() *domain.Reservation {
	return domain.EmptyReservation()
}

func (repo reservationPrimitiveRepoBehaivor) IsZero(reservation domain.Reservation) bool {
	return reservation.Id() == 0 && reservation.BoatId() == 0 && reservation.FirstDay().IsZero() && reservation.LastDay().IsZero() && reservation.StateRoomId() == 0
}
func (repo reservationPrimitiveRepoBehaivor) Scan(rows *sql.Rows) (domain.Reservation, error) {
	var id int
	var name string
	var phone string
	var firstDay string
	var lastDay string
	var boatId int
	var stateRoomId int
	err := rows.Scan(&id, &name, &phone, &firstDay, &lastDay, &boatId, &stateRoomId)
	if err != nil {
		return *domain.EmptyReservation(), err
	}

	user := domain.NewUser(name, phone)
	//TODO CHEANGES DAYS
	return *domain.NewReservation(id, user, time.Now(), time.Now(), boatId, stateRoomId), nil
}

func (repo reservationPrimitiveRepoBehaivor) UpdateRelations(reservation *domain.Reservation) error {
	return nil
}
