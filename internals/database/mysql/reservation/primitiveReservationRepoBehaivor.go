package reservationDB

import (
	"database/sql"
	"time"

	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
	"github.com/lucastomic/naturalYSalvajeRent/internals/timesimplified"
)

type reservationPrimitiveRepoBehaivor struct {
	clientRepository      databaseport.IClientRepository
	clientReservationRepo databaseport.RelationSaver[domain.Client, domain.Reservation]
}

const insertStmt string = "INSERT INTO reservation(firstDay,lastDay,passengers,isOpen,boatId) VALUES(?,?,?,?,?)"
const updateStmt string = "UPDATE reservation SET firstDay = ?, lastDay = ?, passengers = ?, isOpen = ?, boatId = ? WHERE id = ?"
const findByIdStmt string = "SELECT * FROM reservation WHERE id = ?"
const findAllStmt string = "SELECT * FROM reservation"
const removeStmt string = "DELETE FROM reservation WHERE id = ?"

func (repo reservationPrimitiveRepoBehaivor) InsertStmt() string {
	return insertStmt
}
func (repo reservationPrimitiveRepoBehaivor) RemoveStmt() string {
	return removeStmt
}
func (repo reservationPrimitiveRepoBehaivor) UpdateStmt() string {
	return updateStmt
}
func (repo reservationPrimitiveRepoBehaivor) FindByIdStmt() string {
	return findByIdStmt
}
func (repo reservationPrimitiveRepoBehaivor) FindAllStmt() string {
	return findAllStmt
}

func (repo reservationPrimitiveRepoBehaivor) PersistenceValues(reservation domain.Reservation) []any {
	return []any{time.Time(reservation.FirstDay()), time.Time(reservation.LastDay()), reservation.Passengers(), reservation.IsOpen(), reservation.BoatId()}
}

func (repo reservationPrimitiveRepoBehaivor) Id(reservation domain.Reservation) []int {
	return []int{reservation.Id()}
}

func (repo reservationPrimitiveRepoBehaivor) ModifyId(reservation *domain.Reservation, id int64) {
	reservation.SetId(int(id))
}

func (repo reservationPrimitiveRepoBehaivor) Empty() *domain.Reservation {
	return domain.EmptyReservation()
}

func (repo reservationPrimitiveRepoBehaivor) IsZero(reservation domain.Reservation) bool {
	return reservation.Id() == 0 && reservation.FirstDay().IsZero() && reservation.LastDay().IsZero()
}

func (repo reservationPrimitiveRepoBehaivor) Scan(rows *sql.Rows) (domain.Reservation, error) {
	var id, boatId, passengers int
	var firstDay, lastDay string
	var isOpen bool

	err := rows.Scan(&id, &firstDay, &lastDay, &passengers, &isOpen, &boatId)
	if err != nil {
		return *domain.EmptyReservation(), err
	}

	firstDayParsed, _ := timesimplified.FromString(firstDay)
	lastDayParsed, _ := timesimplified.FromString(lastDay)

	if err != nil {
		return *domain.EmptyReservation(), err
	}

	return *domain.NewReservationWithoutClient(id, firstDayParsed, lastDayParsed, isOpen, passengers, boatId), nil
}

func (repo reservationPrimitiveRepoBehaivor) UpdateRelations(reservation *domain.Reservation) error {
	clients, err := repo.clientRepository.FindByReservation(*reservation)
	if err != nil {
		return err
	}
	reservation.SetClients(clients)
	return nil
}

func (repo reservationPrimitiveRepoBehaivor) SaveChildsChanges(reservation *domain.Reservation) error {
	for _, client := range reservation.Clients() {
		err := repo.clientRepository.Save(client)
		if err != nil {
			return err
		}
	}
	return nil
}
func (repo reservationPrimitiveRepoBehaivor) SaveRelations(reservation *domain.Reservation) error {
	for _, client := range reservation.Clients() {
		err := repo.clientReservationRepo.Save(*client, *reservation)
		if err != nil {
			return err
		}
	}
	return nil
}
