package clientDB

import (
	"database/sql"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain"
)

type clientPrimitiveRepoBehaivor struct {
}

const insertStmt string = "INSERT INTO client(name, phone) VALUES(?,?)"
const updateStmt string = "UPDATE client SET name = ?, phone =? WHERE id = ? "
const findStmt string = "SELECT * FROM client WHERE id = ?"
const findAllstmt string = "SELECT * FROM client"
const removeStmt string = "DELETE FROM client WHERE id = ?"

func (repo clientPrimitiveRepoBehaivor) InsertStmt() string {
	return insertStmt
}
func (repo clientPrimitiveRepoBehaivor) RemoveStmt() string {
	return removeStmt
}
func (repo clientPrimitiveRepoBehaivor) UpdateStmt() string {
	return updateStmt
}
func (repo clientPrimitiveRepoBehaivor) FindByIdStmt() string {
	return findStmt
}
func (repo clientPrimitiveRepoBehaivor) FindAllStmt() string {
	return findAllstmt
}

func (repo clientPrimitiveRepoBehaivor) PersistenceValues(client domain.Client) []any {
	return []any{client.Name(), client.Phone()}
}

func (repo clientPrimitiveRepoBehaivor) Empty() *domain.Client {
	return &domain.Client{}
}

func (repo clientPrimitiveRepoBehaivor) Id(client domain.Client) []int {
	return []int{client.Id()}
}

func (repo clientPrimitiveRepoBehaivor) ModifyId(client *domain.Client, id int64) {
	client.SetId(int(id))
}

func (repo clientPrimitiveRepoBehaivor) IsZero(client domain.Client) bool {
	return client.Id() == 0 && client.Name() == "" && client.Phone() == ""
}

// scan scans the stateRoom inside the row passed by argument
func (repo clientPrimitiveRepoBehaivor) Scan(row *sql.Rows) (domain.Client, error) {
	var id int
	var name string
	var phone string

	err := row.Scan(&id, &name, &phone)
	if err != nil {
		return domain.Client{}, err
	}

	return *domain.NewClientWithId(id, name, phone), nil
}

func (repo clientPrimitiveRepoBehaivor) UpdateRelations(client *domain.Client) error {
	return nil
}

func (repo clientPrimitiveRepoBehaivor) SaveChildsChanges(client *domain.Client) error {

	return nil
}

func (repo clientPrimitiveRepoBehaivor) SaveRelations(client *domain.Client) error {
	return nil
}
