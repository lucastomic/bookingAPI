package domain

type Client struct {
	email string
	phone string
}

func NewClient(email string, phone string) Client {
	return Client{email, phone}
}
