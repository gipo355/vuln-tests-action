package nmap

type Config struct {
	Target        string
	OutputDir     string
	WriteToFile   bool
	GenerateSarif bool
}

func NewConfig(target string, writeToFile, sarif bool, outputDir string) *Config {
	return &Config{
		Target:        target,
		WriteToFile:   writeToFile,
		GenerateSarif: sarif,
		OutputDir:     outputDir,
	}
}
