package databaseport

type repository[T, I any] interface {
	Save(T) error
	Remove(T) error
	FindById(...I) (T, error)
	FindAll() ([]T, error)
}
