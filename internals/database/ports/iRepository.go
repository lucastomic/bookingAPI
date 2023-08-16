package databaseport

type Repository[T, I any] interface {
	Save(*T) error
	Remove(T) error
	// If the object is not found given its ID, it returns a nil T object
	FindById(...I) (T, error)
	FindAll() ([]*T, error)
}
