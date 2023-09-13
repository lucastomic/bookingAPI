package domain

type User struct {
	email        string
	hashPassword string
	boats        []*Boat
}

func (u User) Email() string {
	return u.email
}

func (u User) Password() string {
	return u.hashPassword
}

func (u User) Boats() []*Boat {
	return u.boats
}

func EmptyUser() *User {
	return &User{}
}

func (u *User) SetBoats(boats []*Boat) {
	u.boats = boats
}

func NewUser(email string, password string, boats []*Boat) User {
	return User{email, password, boats}
}

func NewUserWithoutBoats(email string, password string) User {
	boats := []*Boat{}
	return User{email, password, boats}
}
