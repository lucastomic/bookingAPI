package mysql

type StateRoomRepository struct {
}

// Save verifies if the boat already exisits in the database. If it does, it updates
// it with the modified values. If it doesn't exist, inserts it into a new row.
// func (repo StateRoomRepository) Save(stateRoom boat.StateRoom) error {
// 	currentBoat, _ := repo.FindById(stateRoom.Id())
// 	if currentBoat.Name() != "" {
// 		err := repo.update(stateRoom)
// 		return err
// 	}
//
// 	db := getInstance()
// 	stmt, err := db.Prepare("INSERT INTO boat(name) VALUES(?)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
//
// 	_, err = stmt.Exec(stateRoom.Name())
// 	return err
// }
//
// // update updates the modified values of boat in the database
// // If an error ocurrs it returns it, otherwise return nil
// func (u BoatRepository) update(boat boat.Boat) error {
// 	db := getInstance()
//
// 	stmt, err := db.Prepare("UPDATE boat SET name = ? WHERE id = ?")
//
// 	if err != nil {
// 		return err
// 	}
//
// 	_, err = stmt.Exec(boat.Name(), boat.Id(), 1)
//
// 	defer stmt.Close()
//
// 	if err != nil {
// 		return err
// 	}
// 	return nil
//
// }
//
// // FindById returns a boat given its Id.
// // If an error ocurrs, it returns an empty boat object with the error as second
// // value. If no error ocurrs, it returns the boat as first parameter and nil as second
// func (u BoatRepository) FindById(id int) (boat.Boat, error) {
// 	var response boat.Boat
// 	db := getInstance()
//
// 	stmt, err := db.Prepare("SELECT name FROM boat WHERE id = ?")
//
// 	if err != nil {
// 		return *boat.EmtyBoat(), err
// 	}
// 	defer stmt.Close()
//
// 	rows, err := stmt.Query(id)
//
// 	if err != nil {
// 		return *boat.EmtyBoat(), err
// 	}
//
// 	rows.Scan(response)
//
// 	return response, nil
// }
