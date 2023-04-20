package database

// Repository is the implementation of the Repository Pattern for a struct
// T with an ID of type I
type Repository[T any, I any] interface {
	// Save takes an object of type T and inserts it on the database or
	// updates it with the modified values if it already exists
	// If ocurrs an error, it returns it, and nil otherwise.
	Save(T) error
	// FindById retrieves the object of type T with the ID specified as argument
	// It returns the T object and nil if there is no error, or, a nil T object with
	// the error otherwise
	FindById(I) (T, error)
}
