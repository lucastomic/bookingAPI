package main

import (
	"fmt"

	"github.com/lucastomic/naturalYSalvajeRent/internals/database"
	boatDB "github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql/boat"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

func main() {
	var repo database.Repository[domain.Boat, int] = boatDB.NewBoatRepository()
	boat, err := repo.FindById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(boat)

}
