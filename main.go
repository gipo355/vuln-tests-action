package main

import (
	"log"
	"os"
)

func main() {
	// get args
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("No args provided")
	}

	name := args[0]

	log.Printf("Hello, %v!", name)
}
