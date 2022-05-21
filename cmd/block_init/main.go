package main

import (
	"fmt"

	block "github.com/baswilson/block/internal/block"
)

func main() {
	block.ValidateOS()
	fmt.Println("[BLOCK] initializing block")

	// Parse and add config to memory
	block.ParseConfig()

	// Pull image / source
	block.Pull()

	// Run image
	block.Run()

	// Start monitoring health
	fmt.Println("[BLOCK] running monitor...")
	block.StartHealthMonitoring()
}
