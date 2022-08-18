package block

import (
	"fmt"
	"os/exec"
)


func Stop(c *Config) error {
	fmt.Println("[BLOCK] Stopping image")

	// Remove old instance
	output, err := exec.Command(
		"docker",
		"rm",
		"-f",
		c.Settings.Name).CombinedOutput()

	if err != nil {
		return fmt.Errorf(fmt.Sprint(err) + ": " + string(output))
	}

	// Remove image
	exec.Command("docker", "image", "rm", "-f", c.Image.Tag, "||", "true").Run()


	// Remove config from list
	RemoveConfig(c)

	ImageId = ""

	return nil
}
