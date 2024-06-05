package nmap

import (
	"log"
	"os"
	"os/exec"
	"slices"
)

// nmap -sV --script=~/vulscan.nse www.example.com

func (n *Client) VulnScan() {
	target := n.Config.Target

	args := slices.Concat(

		[]string{
			"-sV",                          // Version detection
			"--script=vulscan/vulscan.nse", // Script to run
		},

		[]string{target},
	)

	cmd := exec.Command("nmap", args...)

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr
	log.Printf("cmd: %v", cmd)
}
