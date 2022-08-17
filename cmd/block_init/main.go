package main

import (
	"fmt"

	block "github.com/baswilson/block/internal/block"
	"github.com/baswilson/block/internal/webserver"
	"github.com/joho/godotenv"
)

func main() {
	block.ValidateOS()
	
	fmt.Println("[BLOCK] initializing block")
	
	godotenv.Load("build/ci/dev/sim.env")

	// Start up webserver
	webserver.SetupRouter()

	// Start monitoring all services running under the block
	block.StartHealthMonitoring()
}
