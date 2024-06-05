package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gipo355/hello-world-docker-go-action/utils"
)

func main() {
	// get args
	// args are used to pass input to the golang cli program, not to nmap
	// we will use env vars to pass input to nmap or possibly args with --flag like --host=localhost
	// or env INPUT_HOST=localhost
	args := os.Args[1:]
	if len(args) == 0 {
		log.Printf("No args provided")
	} else {
		utils.PrintHello(args[0])

		for _, arg := range args {
			log.Printf("arg: %v", arg)
		}
	}

	// all inputs are passed as env vars
	// inputWhoToGreet := os.Getenv("INPUT_WHO-TO-GREET")
	// args to be passed to the action entrypoint must be passed to args in gh action

	// sets output using deprecated method ::set-output
	// utils.PrintTime()

	// https://stackoverflow.com/questions/71357973/github-actions-set-two-output-names-from-custom-action-in-golang-code
	githubOutput := utils.GithubOutput()
	log.Printf("GITHUB_OUTPUT: %v", githubOutput)

	home := os.Getenv("HOME")
	githubWorkspace := os.Getenv("GITHUB_WORKSPACE")

	log.Println("ls .")
	utils.Ls(".")

	log.Println("ls ..")
	utils.Ls("..")

	log.Println("ls /")
	utils.Ls("/")

	if home != "" {
		log.Println("ls $HOME")
		utils.Ls(os.Getenv("HOME"))
	}

	if githubWorkspace != "" {
		log.Println("ls $GITHUB_WORKSPACE")
		utils.Ls(os.Getenv("GITHUB_WORKSPACE"))

		utils.PrintFileContent(githubOutput)

		utils.AppendToFile(githubOutput, fmt.Sprintf("time=%v\n", time.Now().Format("15:04:05")))

		utils.AppendToFile(githubOutput, fmt.Sprintf("arg=%v", args[0]))

		utils.PrintFileContent(githubOutput)
	}

	// doesn't exist
	// log.Println("ls $RUNNER_WORKSPACE")
	// ls(os.Getenv("RUNNER_WORKSPACE"))

	log.Println("print pwd")
	utils.PrintPwd()

	log.Println("print env")
	utils.PrintEnv()

	log.Println("Executing nmap...")
	// utils.SimpleNmap()
	utils.SimpleNmapStream()
}
