package nmap

type Config struct {
	Host string
	Args []string
}

func NewConfig(host string, args []string) *Config {
	return &Config{
		Host: host,
		Args: args,
	}
}
