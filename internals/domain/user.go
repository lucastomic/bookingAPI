package domain

type User struct {
	name  string `json:"name"`
	phone string `json:"phone"`
}

func NewUser(name string, phone string) User {

	return User{name, phone}
}
