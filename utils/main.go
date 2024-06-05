package utils

import (
	"log"
	"os"
	"os/exec"
)

func GithubState() string {
	return os.Getenv("GITHUB_STATE")
}

func GithubOutput() string {
	return os.Getenv("GITHUB_OUTPUT")
}

func GetInputEnv(name string) string {
	return os.Getenv("INPUT_" + name)
}

func PrintHello(arg string) {
	name := arg

	log.Printf("Hello, %v!", name)
}

func AppendToFile(path, content string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Panic(err)
	}
}

func PrintPwd() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PWD: %v", pwd)
}

func PrintEnv() {
	for _, env := range os.Environ() {
		log.Printf("ENV: %v", env)
	}
}

func Ls(path string) {
	cmd := exec.Command("ls", "-la", path)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func PrintFileContent(path string) {
	buf, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	println(string(buf))
}
