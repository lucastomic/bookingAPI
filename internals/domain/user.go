package domain

type User struct {
	name  string
	phone string
}

func NewUser(name string, phone string) User {

	return User{name, phone}
}
