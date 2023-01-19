package slave

import (
	"fmt"

	"github.com/baswilson/block/internal/shared"
)

func Restart(c *shared.Config) error {
	fmt.Println("[BLOCK] Restarting image")

	// Stop
	err := Stop(c)

	// Starting
	if err == nil {
		err = Run(c)
	}

	return err
}
