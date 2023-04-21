package mysql

import "database/sql"

// ConcretRepository is the implementation of a concret repository, useful to be passed
// as argument of the MysqlRepository template.
// The type T is the struct which the repository handles (for example, for a BoatRepository, it would be Boat)
// and I is the ID type of the struct T
type ConcretRepository[T any, I any] interface {
	// insertStmt returns a string with the SQL statement for inserting a T object in the database
	// For example,
	// INSERT INTO boat(name) VALUES(?)
	insertStmt() string
	// updateStmt returns a string with the SQL statement for update a T object in the database
	// given its Id. For example,
	// UPDATE boat SET name = ? WHERE id = ?
	updateStmt() string
	// findByIdStmt returns a string with the SQL statement for selecting a T object in the database
	// given its Id. For exmaple,
	// SELECT name FROM boat WHERE id = ?
	findByIdStmt() string
	// persistenceValues returns an array with the values to be persisted in the same order as
	// the given in the insertStmt.
	//
	// For exmaple, if the insertStmt is this:
	// INSERT INTO boat(name, stateRoomsNumber, colour) VALUES(?,?,?)
	// the persistenceValues must be
	// []any{boat.Name(),boat.StateRoomsNumber(),boat.Color()}
	persistenceValues(T) []any
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
	id(T) []I
	// empty returns a Zero object of type T
	empty() *T
	// isZero checks whether tge T object passed as argument is a Zero object
	isZero(T) bool
	scan(*sql.Rows) (T, error)
}
