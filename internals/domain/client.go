package domain

type Client struct {
	id         int
	name       string
	phone      string
	passengers int
}

func NewClient(name string, phone string, passengers int) *Client {
	return &Client{name: name, phone: phone, passengers: passengers}
}

func NewClientWithId(id int, name string, phone string, passengers int) *Client {
	return &Client{id: id, name: name, phone: phone, passengers: passengers}
}

func (c Client) Id() int {
	return c.id
}

func (c *Client) SetId(id int) {
	c.id = id
}

func (c *Client) Passengers() int {
	return c.passengers
}

func (c Client) Name() string {
	return c.name
}

func (c Client) Phone() string {
	return c.phone
}
