package slave

import (
	"fmt"
	"os/exec"
	"path"
	"strconv"

	"github.com/baswilson/block/internal/shared"
)

var ImageId string

func Run(c *shared.Config) error {
	fmt.Println("[BLOCK] Running image")

	// Validate image tag
	if len(c.Image.Tag) == 0 {
		return fmt.Errorf("Settings.Image.Tag is not set")
	}

	var output []byte
	var err error

	// Add config to list
	SetConfig(c)

	if c.Settings.Type == "webservice" {
		output, err = exec.Command(
			"docker",
			"run",
			"-d",
			"--name", c.Settings.Name,
			"-p", strconv.Itoa(int(c.Settings.ContainerPort))+":"+strconv.Itoa(int(c.Settings.Port)),
			"--env-file", path.Join(shared.Base, c.Settings.Name+".env"),
			"--restart=always",
			"--network=block_network",
			c.Image.Tag).CombinedOutput()
	} else if c.Settings.Type == "worker" {
		output, err = exec.Command(
			"docker",
			"run",
			"-d",
			"--name", c.Settings.Name,
			"--env-file", path.Join(shared.Base, c.Settings.Name+".env"),
			"--restart=always",
			c.Image.Tag).CombinedOutput()
	}

	if err != nil {
		return fmt.Errorf(fmt.Sprint(err) + ": " + string(output))
	}

	ImageId = string(output)
	fmt.Println("[BLOCK] " + c.Settings.Name + " running: " + ImageId)

	return nil
}
