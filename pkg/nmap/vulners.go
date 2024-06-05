package nmap

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
)

// nmap -sV --script=~/vulscan.nse www.example.com

func (n *Client) ScanWithVulners(c chan<- error) {
	// defer close(c)
	target := n.Config.Target

	args := slices.Concat(

		[]string{
			"-sV",                    // Version detection
			"--script=nmap-vulners/", // Script to run
		},

		[]string{target},
	)

	var err error

	if n.Config.WriteToFile {
		err = n.writeToFile(args, "vulners")
		c <- err
		return
	}

	cmd := exec.Command("nmap", args...)
	log.Printf("cmd: %v", cmd)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr
	log.Printf("cmd: %v", cmd)

	err = fmt.Errorf("nmap: %w", cmd.Run())
	c <- err
}
