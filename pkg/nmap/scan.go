package nmap

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
)

// Scan runs nmap with the provided arguments to stdout
func (n *Client) DirectScan(userArgs []string) error {
	target := n.Config.Target

	args := slices.Concat(userArgs, []string{
		target, // target
	})

	if n.Config.WriteToFile {
		return n.writeToFile(userArgs)
	}

	cmd := exec.Command("nmap", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return fmt.Errorf("nmap: %w", cmd.Run())
}
