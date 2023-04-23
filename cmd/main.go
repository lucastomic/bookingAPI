package main

import (
	"fmt"

	"github.com/lucastomic/naturalYSalvajeRent/internals/database"
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
)

func main() {
	var repo database.Repository[boat.Boat, int] = mysql.NewBoatRepository()
	boat, err := repo.FindById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(boat)

}
