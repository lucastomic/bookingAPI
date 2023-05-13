package domain

type User struct {
	email string
	phone string
}

func NewUser(email string, phone string) User {
	return User{email, phone}
}
