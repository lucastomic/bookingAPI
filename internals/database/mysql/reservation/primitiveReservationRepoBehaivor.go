package reservationDB

import (
	"database/sql"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timeParser"
)

type reservationPrimitiveRepoBehaivor struct {
}

const insertStmt string = "INSERT INTO reservation(name, phone,firstDay,lastDay,boatId,stateRoomId) VALUES(?,?,?,?,?,?)"
const updateStmt string = "UPDATE reservation SET name = ?,phone = ?, firstDay = ?, lastDay = ?, boatId = ?, stateRoomId = ? WHERE id = ?"
const findByIdStmt string = "SELECT * FROM reservation WHERE id = ?"
const findAllStmt string = "SELECT * FROM reservation"
const removeStmt string = "DELETE FROM reservation WHERE id = ?"

func (repo reservationPrimitiveRepoBehaivor) InsertStmt() string {
	return insertStmt
}

func (repo reservationPrimitiveRepoBehaivor) RemoveStmt() string {
	return removeStmt
}

// UpdateStmt returns the SQL statement for updating a reservation in the database.
func (repo reservationPrimitiveRepoBehaivor) UpdateStmt() string {
	return updateStmt
}

// FindByIdStmt returns the SQL statement for finding a reservation by its ID in the database.
func (repo reservationPrimitiveRepoBehaivor) FindByIdStmt() string {
	return findByIdStmt
}
func (repo reservationPrimitiveRepoBehaivor) FindAllStmt() string {
	return findAllStmt
}

// PersistenceValues returns a slice of type []any that contains the values of the reservation's properties to be persisted in the database.
func (repo reservationPrimitiveRepoBehaivor) PersistenceValues(reservation domain.Reservation) []any {
	return []any{reservation.UserName(), reservation.UserPhone(), reservation.FirstDay(), reservation.LastDay(), reservation.BoatId(), reservation.StateRoomId()}
}

// Id returns a slice of type []int that contains the ID of the reservation to be used as a parameter in database operations.
func (repo reservationPrimitiveRepoBehaivor) Id(reservation domain.Reservation) []int {
	return []int{reservation.Id()}
}

// Empty returns a pointer to an empty domain.Reservation object, which can be used as a placeholder or default value.
func (repo reservationPrimitiveRepoBehaivor) Empty() *domain.Reservation {
	return domain.EmptyReservation()
}

// IsZero returns a boolean value indicating whether the reservation object has all its properties set to zero values, which typically indicates that it is empty or uninitialized.
func (repo reservationPrimitiveRepoBehaivor) IsZero(reservation domain.Reservation) bool {
	return reservation.Id() == 0 && reservation.BoatId() == 0 && reservation.FirstDay().IsZero() && reservation.LastDay().IsZero() && reservation.StateRoomId() == 0
}

// Scan takes a pointer to a sql.Rows object as input and scans the rows to populate a domain.Reservation object, which is then returned
// along with any error that may occur during scanning.
func (repo reservationPrimitiveRepoBehaivor) Scan(rows *sql.Rows) (domain.Reservation, error) {
	var id, boatId, stateRoomId int
	var name, phone, firstDay, lastDay string

	err := rows.Scan(&id, &name, &phone, &firstDay, &lastDay, &boatId, &stateRoomId)
	if err != nil {
		return *domain.EmptyReservation(), err
	}

	user := domain.NewUser(name, phone)
	firstDayParsed, _ := timeParser.ParseFromString(firstDay)
	lastDayParsed, _ := timeParser.ParseFromString(lastDay)

	return *domain.NewReservation(id, user, firstDayParsed, lastDayParsed, boatId, stateRoomId), nil
}

// UpdateRelations takes a pointer to a domain.Reservation object as input and updates any related
// entities or relationships in the database. Currently, it returns nil and does not perform any actual updates.
func (repo reservationPrimitiveRepoBehaivor) UpdateRelations(reservation *domain.Reservation) error {
	return nil
}
