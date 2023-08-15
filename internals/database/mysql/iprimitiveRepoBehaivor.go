package mysql

import "database/sql"

// IPrimitiveRepoBehaivor is the behaivor of a repository needed for implementing the Template CommmonMysqlRepository
// The type T is the struct which the repository handles (for example, for a BoatRepository, it would be Boat)
// and I is the ID type of the struct T
type IPrimitiveRepoBehaivor[T any, I any] interface {
	// RemoveStmt returns a string with the SQL statement for removing a T object from the database
	// given its Id
	// For example,
	// DELETE FROM boat WHERE id = ?;
	RemoveStmt() string
	// insertStmt returns a string with the SQL statement for inserting a T object in the database
	// For example,
	// INSERT INTO boat(name) VALUES(?)
	InsertStmt() string
	// updateStmt returns a string with the SQL statement for update a T object in the database
	// given its Id. For example,
	// UPDATE boat SET name = ? WHERE id = ?
	UpdateStmt() string
	// findByIdStmt returns a string with the SQL statement for selecting a T object in the database
	// given its Id. For exmaple,
	// SELECT * FROM boat WHERE id = ?
	FindByIdStmt() string
	// FindAll returns a string with the SQL statement for selecting all the  T object from the database
	// For exmaple,
	// SELECT * FROM boat
	FindAllStmt() string
	// persistenceValues returns an array with the values to be persisted in the same order as
	// the given in the insertStmt.
	//
	// For exmaple, if the insertStmt is this:
	// INSERT INTO boat(name, stateRoomsNumber, colour) VALUES(?,?,?)
	// the persistenceValues must be
	// []any{boat.Name(),boat.StateRoomsNumber(),boat.Color()}
	PersistenceValues(T) []any
	// id returns the id of the T object specified as argument. This ID must be wrapped
	// in a slice, even if is a single element. This way, if the struct has only one ID called "id", it
	// would return []{id}, and, if the struct has a compound id, for example id1 and id2, it would return []{id1,id2}
	// The order of the ids in the slice must be the same as the ones in the findByidStmt and updateStmt.
	//
	// For example, given this statements:
	// UPDATE boat SET name = ? WHERE id1 = ? AND id2 = ?
	// SELECT name FROM boat WHERE id1 = ? AND id2 = ?
	// the slice returned must look like this:
	// []{id1,id2}
	// and not like this:
	// []{id2,id1}
	Id(T) []I
	// empty returns a Zero object of type T
	Empty() *T
	// isZero checks whether tge T object passed as argument is a Zero object
	IsZero(T) bool
	// scan scans the T object in the row passed as argument and returns it parsed into
	// the struct. If there is an error, it returns it as second value
	// It doesn't update de relations of the T object, only the primary elements
	// For example, given a Boat{id int, name string, staeRooms []StateRoom}
	// It will scan the values id and name, but not the relation OneToMany staeRooms
	Scan(*sql.Rows) (T, error)
	// updateRelations takes a T object and updates all its relations (OneToOne, OneToMany, ManyToOne and ManyToMany)
	UpdateRelations(*T) error
	// SaveChildsChanges takes a T object and save the changes in all their childs. If a new child was inserted into
	// the database or if one of them was modified, it persists the changes in the database
	SaveChildsChanges(*T) error
	// SaveRealtions takes a T object and saves all its relations (OneToOne, OneToMany, ManyToOne and ManyToMany)
	SaveRelations(*T) error
}
