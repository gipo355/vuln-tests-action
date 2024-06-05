package nmap

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
)

// Scan runs nmap with the provided arguments to stdout
func (n *Client) DirectScan(nmapArgs []string) error {
	target := n.Config.Target

	args := slices.Concat(nmapArgs, []string{
		target, // target
	})

	if n.Config.WriteToFile {
		return n.writeToFile(nmapArgs, "direct")
	}

	cmd := exec.Command("nmap", args...)
	log.Printf("cmd: %v", cmd)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return fmt.Errorf("nmap: %w", cmd.Run())
}
