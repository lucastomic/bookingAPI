package main

import (
	"fmt"

	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
)

func main() {

	repo := databaseport.NewBoatRepository()

	boat, err := repo.FindById(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(boat.StateRooms()[0].ReservedDays()[0].FirstDay().Date())

}
