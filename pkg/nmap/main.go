package nmap

type Client struct {
	Config *Config
}

func NewNmapClient(host string, args []string) *Client {
	return &Client{
		Config: NewConfig(host, args),
	}
}
