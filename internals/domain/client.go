package domain

type Client struct {
	id    int
	name  string
	phone string
}

func NewClient(name string, phone string) *Client {
	return &Client{name: name, phone: phone}
}

func NewClientWithId(id int, name string, phone string) *Client {
	return &Client{id: id, name: name, phone: phone}
}

func (c Client) Id() int {
	return c.id
}

func (c Client) Name() string {
	return c.name
}
func (c Client) Phone() string {
	return c.phone
}
