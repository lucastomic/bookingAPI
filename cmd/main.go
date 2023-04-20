package main

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database"
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
)

func main() {
	var repo database.Repository[boat.Boat, int] = mysql.BoatRepository{}
	boat := boat.NewBoat(1, "Sol de Mayo", []boat.StateRoom{})
	repo.Save(*boat)

}
