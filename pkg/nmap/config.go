package nmap

type Config struct {
	Target      string
	WriteToFile bool
}

func NewConfig(target string, writeToFile bool) *Config {
	return &Config{
		Target:      target,
		WriteToFile: writeToFile,
	}
}
