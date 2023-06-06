package domain

type User struct {
	email    string
	password string
	boats    []Boat
}

func (u User) Email() string {
	return u.email
}
func (u User) Password() string {
	return u.password
}
func (u User) Boats() []Boat {
	return u.boats
}
func EmptyUser() *User {
	return &User{}
}
func (u User) SetBoats(boats []Boat) {
	u.boats = boats
}
func NewUser(email string, password string, boats []Boat) User {
	return User{email, password, boats}
}
