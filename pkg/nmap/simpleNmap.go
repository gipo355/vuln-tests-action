package nmap

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func (n *NmapConfig) SimpleNmap() {
	// host := github.GetInputEnv("HOST")
	host := n.Host
	args := strings.Join(n.Args, " ")

	// this will execute the nmap, here we need to compose the command based on user input
	// e.g. use scripts, flags, host, etc
	cmd := exec.Command("nmap", args, host)

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
