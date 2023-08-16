package boatDB

import (
	"github.com/lucastomic/naturalYSalvajeRent/internals/database/mysql"
	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

// BoatRepository manages all the interactions between the Boat entity and the database
type BoatRepository struct {
	mysql.CommonMysqlLogic[domain.Boat, int]
}

const findBoatsByOwnerStmt string = "SELECT * FROM boat WHERE owner = ? "

func NewBoatRepository(
	stateRoomRepo databaseport.IStateRoomRepository,
) BoatRepository {
	commonBehaivor := mysql.CommonMysqlLogic[domain.Boat, int]{
		IPrimitiveRepoBehaivor: boatPrimitiveRepoBehaivor{
			stateRoomRepo: stateRoomRepo,
		},
	}
	return BoatRepository{commonBehaivor}
}

func (repo BoatRepository) FindByUser(email string) ([]*domain.Boat, error) {
	response, err := repo.Query(findBoatsByOwnerStmt, []any{email})
	if err != nil {
		return []*domain.Boat{}, err
	}
	return response, nil
}
