package github

import (
	"fmt"
	"os"
	"strings"
)

// func GetStatePath() string {
// 	return os.Getenv("GITHUB_STATE")
// }
//
// func GetOutputPath() string {
// 	return os.Getenv("GITHUB_OUTPUT")
// }

// GetInputEnv returns the value of the input env var, requires "UPPER-VAR".
func (e *Environment) GetUserInputFromEnv(name string) string {
	upperName := strings.ToUpper(name)

	return os.Getenv("INPUT_" + upperName)
}

// SetOutput sets the output of the action, requires "$name=$value" format.
func (e *Environment) SetOutput(name, value string) error {
	path := e.GITHUB_OUTPUT
	if path == "" {
		return fmt.Errorf("Output path not set")
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	content := fmt.Sprintf("%s=%s\n", name, value)

	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return nil
}

// SetState sets the state of the action, requires TODO:
func (e *Environment) SetState(content string) error {
	path := e.GITHUB_STATE
	if path == "" {
		return fmt.Errorf("State path not set")
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		return err
	}

	return nil
}
