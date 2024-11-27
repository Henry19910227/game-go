package game

type Channel struct {
	channel map[string][]*Client
}

func NewChannel() *Channel {
	return &Channel{
		channel: make(map[string][]*Client),
	}
}

func (c *Channel) Add(name string, client *Client) {
	_, ok := c.channel[name]
	if !ok {
		c.channel[name] = make([]*Client, 0)
	}
	c.channel[name] = append(c.channel[name], client)
}
