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

	// https://stackoverflow.com/questions/71357973/github-actions-set-two-output-names-from-custom-action-in-golang-code
	githubOutput := os.Getenv("GITHUB_OUTPUT")
	log.Printf("GITHUB_OUTPUT: %v", githubOutput)

	date := time.Now().Format("2006-01-02 15:04:05")
	// newEnv := fmt.Sprintf("\ndate=%s", date)
	// newGithubOutput := githubOutput + newEnv
	// os.Setenv("GITHUB_OUTPUT", newGithubOutput)

	// fmt.Printf(`::set-output name=repo_tag::%s`, value)
	// fmt.Print("\n")
	// fmt.Printf(`::set-output name=ecr_tag::%s`, "v"+value)
	// fmt.Print("\n")
	fmt.Printf(`::set-output date=%s`, date)
	fmt.Print("\n")

	log.Printf("Hello, %v!", name)
}
