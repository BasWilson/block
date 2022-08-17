package block

import (
	"fmt"
)


func Restart(c *Config) error {
	fmt.Println("[BLOCK] Restarting image")

	// Stop
	err := Stop(c)

	// Starting
	if err == nil {
		err = Run(c)
	}

	return err
}
