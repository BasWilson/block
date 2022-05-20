package main

import (
	"fmt"

	block "github.com/baswilson/block/internal/block"
)

func main() {
	block.ValidateOS()
	fmt.Println("[BLOCK] initializing block")
	block.Pull()
	block.Run()
}
