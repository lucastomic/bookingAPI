package databaseport

type RelationSaver[T any, I any] interface {
	Save(T, I) error
}
