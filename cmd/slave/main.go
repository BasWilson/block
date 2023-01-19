package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/baswilson/block/internal/shared"
	"github.com/baswilson/block/internal/slave"
	"github.com/baswilson/block/internal/slave/webserver"
	"github.com/joho/godotenv"
)

func main() {
	shared.ValidateOS()

	fmt.Println("[BLOCK] initializing slave")

	executable, _ := os.Executable()
	slave.Base = strings.Split(executable, "/bin/slave")[0]

	godotenv.Load(path.Join(slave.Base, "build/ci/dev/sim.env"))

	// Register block with block registry
	go slave.RegisterBlock()

	// Start monitoring all services running under the block
	go slave.StartHealthMonitoring()

	// Start up webserver
	webserver.SetupRouter()
}
