package nmap

import (
	"log"
	"os"
	"os/exec"

	"github.com/gipo355/hello-world-docker-go-action/pkg/github"
)

func SimpleNmap() {
	host := github.GetInputEnv("HOST")

	// this will execute the nmap, here we need to compose the command based on user input
	// e.g. use scripts, flags, host, etc
	cmd := exec.Command("nmap", "-sP", host)

	file, fileErr := os.Create("nmap.log")
	if fileErr != nil {
		log.Panic(fileErr)
	}
	defer file.Close()

	// this is a writer, we want to write to a file with bufio
	// cmd.Stdout = os.Stdout
	cmd.Stdout = file

	cmd.Stderr = os.Stderr

	if cmdErr := cmd.Run(); cmdErr != nil {
		log.Panic(cmdErr)
	}
}
