package mysql

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
)

// BoatRepository manages the DB logic of the Boat struct
type BoatRepository struct {
}

const insertStmt string = "INSERT INTO boat(name) VALUES(?)"

// Save verifies if the boat already exisits in the database. If it does, it updates
// it with the modified values. If it doesn't exist, inserts it into a new row.
func (u BoatRepository) Save(boat boat.Boat) error {

	if u.alreadyExists(boat) {
		err := u.update(boat)
		return err
	}

	db := getInstance()
	stmt, err := db.Prepare(insertStmt)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(boat.Name())
	return err
}

// alreadyExists checks whether a boat has already been inserted into the DB
// Returns true if exists a row with the id of the boat passed as argument, or false otherwise
func (u BoatRepository) alreadyExists(boat boat.Boat) bool {
	currentBoat, _ := u.FindById(boat.Id())
	return currentBoat.Name() != ""
}

// update updates the modified values of boat in the database
// If an error ocurrs it returns it, otherwise return nil
func (u BoatRepository) update(boat boat.Boat) error {
	db := getInstance()

	stmt, err := db.Prepare("UPDATE boat SET name = ? WHERE id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(boat.Name(), boat.Id(), 1)

	defer stmt.Close()

	if err != nil {
		return err
	}
	return nil

}

// FindById returns a boat given its Id.
// If an error ocurrs, it returns an empty boat object with the error as second
// value. If no error ocurrs, it returns the boat as first parameter and nil as second
func (u BoatRepository) FindById(id int) (boat.Boat, error) {
	var response boat.Boat
	db := getInstance()

	stmt, err := db.Prepare("SELECT name FROM boat WHERE id = ?")

	if err != nil {
		return *boat.EmtyBoat(), err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		return *boat.EmtyBoat(), err
	}

	rows.Scan(response)

	return response, nil
}
