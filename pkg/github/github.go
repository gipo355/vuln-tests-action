package github

import (
	"log"
	"os"
	"strings"
)

type Vars struct {
	Workspace  string
	Home       string
	StatePath  string
	OutputPath string
}

func NewGitHubEnvironment() *Vars {
	return &Vars{
		Workspace:  os.Getenv("GITHUB_WORKSPACE"),
		Home:       os.Getenv("HOME"),
		StatePath:  os.Getenv("GITHUB_STATE"),
		OutputPath: os.Getenv("GITHUB_OUTPUT"),
	}
}

func GetStatePath() string {
	return os.Getenv("GITHUB_STATE")
}

func GetOutputPath() string {
	return os.Getenv("GITHUB_OUTPUT")
}

// GetInputEnv returns the value of the input env var, requires "UPPER-VAR".
func GetInputEnv(name string) string {
	upperName := strings.ToUpper(name)

	return os.Getenv("INPUT_" + upperName)
}

// SetOutput sets the output of the action, requires "$name=$value" format.
func SetOutput(content string) {
	path := GetOutputPath()

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		log.Panic(err)
	}
}

// SetState sets the state of the action, requires TODO:
func SetState(content string) {
	path := GetStatePath()

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		log.Panic(err)
	}
}
