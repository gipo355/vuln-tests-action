package nmap

type NmapConfig struct {
	Host string
	Args []string
}

func NewNmapClient(host string, args []string) *NmapConfig {
	return &NmapConfig{
		Host: host,
		// Args: []string{"-sP", host},
		Args: args,
	}
}
