package main

import (
	"fmt"
	"path"

	"github.com/baswilson/block/internal/shared"
	"github.com/baswilson/block/internal/slave"
	"github.com/joho/godotenv"
)

func main() {
	shared.ValidateOS()

	fmt.Println("[BLOCK_SLAVE] initializing slave")

	shared.SetBasePath()

	godotenv.Load(path.Join(shared.Base, "build/ci/dev/sim.env"))

	// Register block with block registry
	go slave.RegisterBlock()

	// Start monitoring all services running under the block
	go slave.StartHealthMonitoring()

	// Start up webserver
	slave.SetupRouter()
}
