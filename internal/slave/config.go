package slave

import (
	"fmt"

	"github.com/baswilson/block/internal/shared"
)

// create variable of slice configs
var Config = shared.Config{}

func SetConfig(c *shared.Config) *shared.Config {
	// Add config to list
	Config = *c

	// Write env variables to file
	shared.WriteEnvVariablesToFile(c)

	fmt.Println("[BLOCK] Updated config")
	return c
}
