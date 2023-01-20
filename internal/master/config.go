package master

import (
	"fmt"

	"github.com/baswilson/block/internal/shared"
)

var Config = shared.Config{}

func AddConfig(c *shared.Config) *shared.Config {
	Config = *c

	// Write env variables to file
	shared.WriteEnvVariablesToFile(c)

	fmt.Println("[BLOCK] Config applied successfully")
	return c
}
