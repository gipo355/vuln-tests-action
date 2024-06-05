package nmap

import (
	"log"
	"os/exec"
)

type Client struct {
	Config *Config
}

func NewNmapClient(c *Config) (*Client, error) {
	path, err := exec.LookPath("ls")
	if err != nil {
		return nil, err
	}

	log.Println(path) // bin/ls

	return &Client{
		Config: c,
	}, nil
}
