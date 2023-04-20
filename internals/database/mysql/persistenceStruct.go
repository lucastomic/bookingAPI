package mysql

type ConcretRepository[T any, I any] interface {
	insertStmt() string
	updateStmt() string
	findByIdStmt() string
	persistenceValues(T) []any
	id(T) I
	empty() *T
}
