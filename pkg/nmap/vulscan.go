package nmap

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"slices"
	"sync"
)

// nmap -sV --script=~/vulscan.nse www.example.com

func (n *Client) ScanWithVulscan(c chan<- error, wg *sync.WaitGroup) {
	// defer close(c)
	target := n.Config.Target

	defer wg.Done()

	args := slices.Concat(

		[]string{
			"-sV",               // Version detection
			"--script=vulscan/", // Script to run
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
