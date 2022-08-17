package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	block "github.com/baswilson/block/internal/block"
	"github.com/baswilson/block/internal/webserver"
	"github.com/joho/godotenv"
)

func main() {
	block.ValidateOS()
	
	fmt.Println("[BLOCK] initializing block")

	executable, _ := os.Executable();
	block.Base = strings.Split(executable, "/bin/block_init")[0]

	godotenv.Load(path.Join(block.Base, "build/ci/dev/sim.env"))

	// Start up webserver
	webserver.SetupRouter()

	// Start monitoring all services running under the block
	block.StartHealthMonitoring()
}
