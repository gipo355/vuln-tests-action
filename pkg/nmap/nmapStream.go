package nmap

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/gipo355/hello-world-docker-go-action/pkg/github"
)

func SimpleNmapStream() {
	host := github.GetInputEnv("HOST")

	// this will execute the nmap, here we need to compose the command based on user input
	// e.g. use scripts, flags, host, etc
	cmd := exec.Command("nmap", "-sP", host)

	file, fileErr := os.Create("nmap.log")
	if fileErr != nil {
		log.Panic(fileErr)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// create a pipe to capture stdout and stderr
	stdout, outPipeErr := cmd.StdoutPipe()
	if outPipeErr != nil {
		log.Panic(outPipeErr)
	}

	stderr, errPipeErr := cmd.StderrPipe()
	if errPipeErr != nil {
		log.Panic(errPipeErr)
	}

	// start the command
	if cmdErr := cmd.Start(); cmdErr != nil {
		log.Panic(cmdErr)
	}

	// create a goroutine to copy the stdout in a stream
	go func() {
		_, copyStdoutErr := io.Copy(writer, stdout)
		if copyStdoutErr != nil {
			log.Panic(copyStdoutErr)
		}
	}()

	// create a goroutine to copy the stderr in a stream
	go func() {
		_, copyStderrErr := io.Copy(os.Stderr, stderr)
		if copyStderrErr != nil {
			log.Panic(copyStderrErr)
		}
	}()

	// wait for the command to finish
	if err := cmd.Wait(); err != nil {
		log.Panic(err)
	}
}
