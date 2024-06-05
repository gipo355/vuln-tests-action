package main

import (
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

	// githubOutput := github.GetOutputPath()
	// log.Printf("GITHUB_OUTPUT: %v", githubOutput)

	// home := os.Getenv("HOME")
	// githubWorkspace := os.Getenv("GITHUB_WORKSPACE")

	// if home != "" {
	// 	log.Println("ls $HOME")
	// 	utils.ListFolderContent(os.Getenv("HOME"))
	// }

	// if githubWorkspace != "" {
	// 	log.Println("ls $GITHUB_WORKSPACE")
	// 	utils.ListFolderContent(os.Getenv("GITHUB_WORKSPACE"))
	//
	// 	utils.PrintFileContent(githubOutput)
	//
	// 	utils.AppendToFile(githubOutput, fmt.Sprintf("time=%v\n", time.Now().Format("15:04:05")))
	//
	// 	utils.AppendToFile(githubOutput, fmt.Sprintf("arg=%v", args[0]))
	//
	// 	utils.PrintFileContent(githubOutput)
	// }

	log.Println("print pwd")
	utils.PrintPwd()

	log.Println("print env")
	utils.PrintEnvVars()

	log.Println("Executing nmap...")

	// TODO: github must move to its own package

	// github section
	gh, err := github.NewGitHubEnvironment()
	if err != nil || gh == nil {
		log.Fatal(err)
	}

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

	// TODO: nmap must move to its own package
	// nmap section

	nmapArgs := []string{"-sV", "-p", "80,443", "-oN", "nmap.log"}

	// utils.SimpleNmap()

	n := nmap.NewNmapClient(
		"localhost",
		nmapArgs,
	)
	n.WriteToFile()
}
