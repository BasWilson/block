package slave

import (
	"fmt"

	"github.com/baswilson/block/internal/shared"
)

var Base string

// create variable of slice configs
var Configs = []shared.Config{}

func AddConfig(c *shared.Config) *shared.Config {
	// Add config to list
	Configs = append(Configs, *c)

	// Write env variables to file
	WriteEnvVariablesToFile(c)

	fmt.Println("[BLOCK] Config added for: " + ": " + c.Settings.Name)
	return c
}

func RemoveConfig(c *shared.Config) (shared.Config, error) {
	for i, v := range Configs {
		if v.Settings.Name == c.Settings.Name {
			Configs = append(Configs[:i], Configs[i+1:]...)
			DeleteEnvVariablesFile(c)
			fmt.Println("[BLOCK] Config removed for: " + c.Settings.Name)
			return v, nil
		}
	}
	return shared.Config{}, fmt.Errorf("[BLOCK] Config not found")
}
