package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gipo355/vuln-tests-action/pkg/github"
	"github.com/gipo355/vuln-tests-action/pkg/nmap"
	"github.com/gipo355/vuln-tests-action/pkg/utils"
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

	log.Println("print pwd")
	utils.PrintPwd()

	// log.Println("print env")
	// utils.PrintEnvVars()

	// TODO: github must move to its own package

	// github section
	gh, err := github.NewGitHubEnvironment()
	if err == nil {
		// some logging
		log.Println("github output", gh.GITHUB_OUTPUT)
		log.Println("github state", gh.GITHUB_STATE)
		log.Println("github workspace", gh.GITHUB_WORKSPACE)
		log.Println("home", gh.HOME)
		log.Println("ls .")
		utils.ListFolderContent(".")
		log.Println("ls ..")
		utils.ListFolderContent("..")
		log.Println("ls /")
		utils.ListFolderContent("/")

		err = gh.SetOutput("time", time.Now().Format("15:04:05"))
		if err != nil {
			log.Fatal(err)
		}
		err = gh.SetOutput("arg", args[0])
		if err != nil {
			log.Fatal(err)
		}
	}

	// TODO: nmap must move to its own package
	// nmap section

	log.Println("Executing nmap...")

	n, err := nmap.NewNmapClient(
		&nmap.Config{
			Target:        "localhost",
			WriteToFile:   true,
			GenerateSarif: false,
			OutputDir:     "nmap-reports",
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Testing
	// TODO: remove hardcoded args
	nmapArgs := []string{"-sP"}

	channels := []chan error{
		make(chan error),
		// make(chan error),
		// make(chan error),
	}

	// directChan := make(chan error)
	go n.DirectScan(nmapArgs, channels[0])

	// vulscanChan := make(chan error)
	// go n.ScanWithVulscan(channels[1])

	// vulnerChan := make(chan error)
	// go n.ScanWithVulners(channels[2])

	// for i := 0; i < 3; i++ {
	for i := 0; i < len(channels); i++ {
		select {
		// case directErr := <-directChan:
		case directErr := <-channels[0]:
			if directErr != nil {
				log.Panic(fmt.Errorf("error direct scanning: %w", directErr))
			}
			log.Println("direct scan finished")

			// case vulnerErr := <-channels[1]:
			// 	if vulnerErr != nil {
			// 		log.Panic(fmt.Errorf("error scanning with vulners: %w", vulnerErr))
			// 	}
			// 	log.Println("vulners scan finished")
			//
			// case vulscanErr := <-channels[2]:
			// 	if vulscanErr != nil {
			// 		log.Panic(fmt.Errorf("error scanning with vulscan: %w", vulscanErr))
			// 	}
			// 	log.Println("vulscan scan finished")
		}
	}

	log.Println("nmap finished")

	// BUG: doesn't wait for the goroutines to finish
	n.ConverToJSON()

	for _, ch := range channels {
		close(ch)
	}
}
