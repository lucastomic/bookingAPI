package database

type Repository[T any, I any] interface {
	Save(T)
	FindById(I)
}
