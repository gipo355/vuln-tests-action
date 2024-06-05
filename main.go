package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func GithubState() string {
	return os.Getenv("GITHUB_STATE")
}

func GithubOutput() string {
	return os.Getenv("GITHUB_OUTPUT")
}

func main() {
	// get args
	args := os.Args[1:]
	if len(args) == 0 {
		log.Panic("No args provided")
	}
	log.Printf("args: %v", args)

	// https://stackoverflow.com/questions/71357973/github-actions-set-two-output-names-from-custom-action-in-golang-code
	githubOutput := GithubOutput()
	log.Printf("GITHUB_OUTPUT: %v", githubOutput)

	// all inputs are passed as env vars
	// inputWhoToGreet := os.Getenv("INPUT_WHO-TO-GREET")
	// args to be passed to the action entrypoint must be passed to args in gh action

	// sets output using deprecated method ::set-output
	// printTime()

	log.Println("ls .")
	ls(".")
	log.Println("ls ..")
	ls("..")
	log.Println("ls /")
	ls("/")
	log.Println("ls $HOME")
	ls(os.Getenv("HOME"))
	log.Println("ls $GITHUB_WORKSPACE")
	ls(os.Getenv("GITHUB_WORKSPACE"))
	// /github/home

	log.Println("ls $RUNNER_WORKSPACE")
	ls(os.Getenv("RUNNER_WORKSPACE"))
	printPwd()
	printEnv()
	printFileContent(githubOutput)
	appendToFile(githubOutput, fmt.Sprintf("time=%v\n", time.Now().Format("15:04:05")))
	appendToFile(githubOutput, fmt.Sprintf("arg=%v", args[0]))
	printFileContent(githubOutput)

	printHello(args[0])
}

func printHello(arg string) {
	name := arg
	log.Printf("Hello, %v!", name)
}

func appendToFile(path, content string) {
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

func printPwd() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("PWD: %v", pwd)
}

func printEnv() {
	for _, env := range os.Environ() {
		log.Printf("ENV: %v", env)
	}
}

func ls(path string) {
	cmd := exec.Command("ls", "-la", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func printFileContent(path string) {
	buf, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	println(string(buf))
}

func printTime() {
	// date := currentTime.Now().Format("2006-01-02 15:04:05")
	currentTime := time.Now().Format("15:04:05")
	// newEnv := fmt.Sprintf("\ndate=%s", date)
	// newGithubOutput := githubOutput + newEnv
	// os.Setenv("GITHUB_OUTPUT", newGithubOutput)

	// https://github.blog/changelog/2022-10-11-github-actions-deprecating-save-state-and-set-output-commands/

	// fmt.Printf(`::set-output name=repo_tag::%s`, value)
	// fmt.Print("\n")
	// fmt.Printf(`::set-output name=ecr_tag::%s`, "v"+value)
	// fmt.Print("\n")

	// WARN: deprecated
	fmt.Printf(`::set-output name=time::%s`, currentTime)
	fmt.Print("\n")

	// DEPRECATED ::set-output
	// must use echo
	// - name: Save state
	// run: echo "{name}={value}" >> $GITHUB_STATE
	//
	// - name: Set output
	// run: echo "{name}={value}" >> $GITHUB_OUTPUT

	// https://github.com/orgs/community/discussions/38570

	// must write to external file under $GITHUB_OUTPUT

	// https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions#setting-an-output-parameter

	// shCmd := fmt.Sprintf("'echo time=%s'", currentTime)
	// cmd := exec.Command("sh", "-c", shCmd, ">>", githubOutput)
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
