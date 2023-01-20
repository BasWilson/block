package main

import (
	"fmt"
	"path"

	"github.com/baswilson/block/internal/master"
	"github.com/baswilson/block/internal/shared"
	"github.com/joho/godotenv"
)

func main() {
	shared.ValidateOS()

	fmt.Println("[BLOCK_MASTER] initializing block_master")

	shared.SetBasePath()

	godotenv.Load(path.Join(shared.Base, "build/ci/dev/sim.env"))

	// Start up webserver
	master.SetupRouter()
}
