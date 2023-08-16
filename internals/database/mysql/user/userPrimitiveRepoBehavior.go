package userDB

import (
	"database/sql"

	databaseport "github.com/lucastomic/naturalYSalvajeRent/internals/database/ports"
	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type userPrimitiveRepoBehaivor struct {
	boatRepo databaseport.IBoatRepository
}

const insertStmt string = "INSERT INTO user(email,password) VALUES(?,?)"
const updateStmt string = "UPDATE user SET email = ?, password = ? WHERE email = ?"
const findByIdStmt string = "SELECT email, password FROM user WHERE email = ?"
const findAllStmt string = "SELECT email, password FROM user"
const removeStmt string = "DELETE FROM user WHERE email = ?"

// insertStmt returns the statement to insert a new user
func (b userPrimitiveRepoBehaivor) InsertStmt() string {
	return insertStmt
}

// RemoveStmt returns rhe statement to remove a user
func (b userPrimitiveRepoBehaivor) RemoveStmt() string {
	return removeStmt
}

// updateStmt returns the statement to update a new user
func (b userPrimitiveRepoBehaivor) UpdateStmt() string {
	return updateStmt
}

// findByIdStmt returns the statement to findByIdStmt a new user
func (b userPrimitiveRepoBehaivor) FindByIdStmt() string {
	return findByIdStmt
}

// findByIdStmt returns the statement to findByIdStmt a new user
func (b userPrimitiveRepoBehaivor) FindAllStmt() string {
	return findAllStmt
}

func (b userPrimitiveRepoBehaivor) PersistenceValues(user domain.User) []any {
	return []any{user.Email(), user.Password()}
}

// empty returns an empty user
func (b userPrimitiveRepoBehaivor) Empty() *domain.User {
	return domain.EmptyUser()
}

// id returns the id of the user passed as argument
func (b userPrimitiveRepoBehaivor) Id(user domain.User) []string {
	return []string{user.Email()}
}

func (repo userPrimitiveRepoBehaivor) ModifyId(user *domain.User, id int64) {
	//Doesn't have to modify it
}

// isZero checks wether the user specified as paramter is a zero user
func (b userPrimitiveRepoBehaivor) IsZero(user domain.User) bool {
	return user.Email() == ""
}

// scan scans the user inside the row passed by argument
func (repo userPrimitiveRepoBehaivor) Scan(row *sql.Rows) (domain.User, error) {
	var email string
	var password string
	var boats []*domain.Boat = []*domain.Boat{}
	err := row.Scan(&email, &password)
	if err != nil {
		return *domain.EmptyUser(), err
	}
	return domain.NewUser(email, password, boats), nil
}

func (repo userPrimitiveRepoBehaivor) UpdateRelations(user *domain.User) error {
	boats, err := repo.boatRepo.FindByUser(user.Email())
	if err != nil {
		return err
	}
	user.SetBoats(boats)
	return nil
}

// SaveChildsChanges takes all the staterooms in the user and save their changes in the datanase (or
// inserts a new stateroom if it's a new one)
func (repo userPrimitiveRepoBehaivor) SaveChildsChanges(user *domain.User) error {
	for _, boat := range user.Boats() {
		err := repo.boatRepo.Save(boat)
		if err != nil {
			return err
		}
	}
	return nil
}
func (repo userPrimitiveRepoBehaivor) SaveRelations(user *domain.User) error {
	return nil
}
