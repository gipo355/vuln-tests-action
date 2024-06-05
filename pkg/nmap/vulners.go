package nmap

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
)

// nmap -sV --script=~/vulscan.nse www.example.com

func (n *Client) ScanVulners() error {
	target := n.Config.Target

	args := slices.Concat(

		[]string{
			"-sV",                    // Version detection
			"--script=nmap-vulners/", // Script to run
		},

		[]string{target},
	)

	if n.Config.WriteToFile {
		return n.writeToFile(args, "vulners")
	}

	cmd := exec.Command("nmap", args...)
	log.Printf("cmd: %v", cmd)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr
	log.Printf("cmd: %v", cmd)

	return fmt.Errorf("nmap: %w", cmd.Run())
}
