package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// get args
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("No args provided")
	}

	name := args[0]

	githubOutput := os.Getenv("GITHUB_OUTPUT")

	date := time.Now().Format("2006-01-02 15:04:05")

	newEnv := fmt.Sprintf("\ndate=%s", date)

	newGithubOutput := githubOutput + newEnv

	os.Setenv("GITHUB_OUTPUT", newGithubOutput)

	log.Printf("Hello, %v!", name)
}
