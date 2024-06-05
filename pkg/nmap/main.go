package nmap

type Client struct {
	Config *Config
}

func NewNmapClient(c *Config) *Client {
	return &Client{
		Config: c,
	}
}
