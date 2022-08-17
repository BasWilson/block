package block

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var ImageId string


func Run(c *Config) (error) {
	fmt.Println("[BLOCK] Running image")

	// Validate image tag
	if len(c.Image.Tag) == 0 {
		return fmt.Errorf("Settings.Image.Tag is not set")
	}
	
	var output []byte
	var err error

	if (c.Settings.Type == "webservice") {
		output, err = exec.Command(
			"docker",
			"run",
			"-d",
			"--name", c.Settings.Name,
			"-p", "80:"+strconv.Itoa(int(c.Settings.Port)),
			"--env-file", os.Getenv("BLOCK_ENV_PATH"),
			"--restart=always",
			"--network=block_network",
			c.Image.Tag).CombinedOutput()
	} else if (c.Settings.Type == "worker") {
		output, err = exec.Command(
			"docker",
			"run",
			"-d",
			"--name", c.Settings.Name,
			"--env-file", os.Getenv("BLOCK_ENV_PATH"),
			"--restart=always",
			c.Image.Tag).CombinedOutput()
	}

	if err != nil {
		return fmt.Errorf(fmt.Sprint(err) + ": " + string(output))
	}

	// Add config to list
	AddConfig(c)

	ImageId = string(output)
	fmt.Println("[BLOCK] " + c.Settings.Name + " running: " + ImageId)

	return nil
}
