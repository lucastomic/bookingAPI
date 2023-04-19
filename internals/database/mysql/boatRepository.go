package database

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
)

type BoatRepository struct {
}

func (u BoatRepository) Save(boat boat.Boat) {
	db := getInstance()
	stmt, err := db.Prepare("INSERT INTO (name) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	stmt.Exec(boat.Name)
}
