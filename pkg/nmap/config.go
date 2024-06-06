package nmap

type Config struct {
	Target        string
	Port          string
	OutputDir     string
	WriteToFile   bool
	GenerateSarif bool
}

func NewConfig(writeToFile, sarif bool, target, outputDir, port string) *Config {
	return &Config{
		Target:        target,
		Port:          port,
		WriteToFile:   writeToFile,
		GenerateSarif: sarif,
		OutputDir:     outputDir,
	}
}
