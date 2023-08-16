package databaseport

type RelationSaver[T any, I any] interface {
	// Save saves a relation between T and I objects in a relation-table.
	// If the relation already exists, it doesn't renturn any error.
	Save(T, I) error
}
