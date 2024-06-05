package nmap

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
)

// nmap -sV --script=~/vulscan.nse www.example.com

func (n *Client) ScanWithVulscan() error {
	target := n.Config.Target

	args := slices.Concat(

		[]string{
			"-sV",                          // Version detection
			"--script=vulscan/vulscan.nse", // Script to run
		},

		[]string{target},
	)

	if n.Config.WriteToFile {
		return n.writeToFile(args, "vulscan")
	}

	cmd := exec.Command("nmap", args...)
	log.Printf("cmd: %v", cmd)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr
	log.Printf("cmd: %v", cmd)

	return fmt.Errorf("nmap: %w", cmd.Run())
}
