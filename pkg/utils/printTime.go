package utils

import (
	"fmt"
	"time"
)

func PrintTime() {
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
