package nmap

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
)

// nmap -sV --script=~/vulscan.nse www.example.com

func (n *Client) ScanWithVulscan(c chan<- error) {
	target := n.Config.Target

	args := slices.Concat(

		[]string{
			"-sV",                          // Version detection
			"--script=vulscan/vulscan.nse", // Script to run
		},

		[]string{target},
	)

	if n.Config.WriteToFile {
		c <- n.writeToFile(args, "vulscan")
		return
	}

	cmd := exec.Command("nmap", args...)
	log.Printf("cmd: %v", cmd)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr
	log.Printf("cmd: %v", cmd)

	c <- fmt.Errorf("nmap: %w", cmd.Run())
}
