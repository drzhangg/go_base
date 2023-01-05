package demo2

type Interface interface {
	AppsV1() string
	CoreV1() string
}

type Client struct {
}

func (c *Client) AppsV1() string {
	return "apps/v1"
}

func (c *Client) CoreV1() string {
	return "core/v1"
}

func NewClient() *Client {
	c := Client{}
	return &c
}
