package slave

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/baswilson/block/internal/shared"
)

func Pull(c *shared.Config) error {
	fmt.Println("[BLOCK] Pulling image")

	// Check for pull mode
	if c.PullMode == "repo" {
		return fmt.Errorf("PullMode 'repo' is not yet supported")
	} else if c.PullMode == "image" {
		err := image(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func image(c *shared.Config) error {
	// Validate image tag
	if len(c.Image.Tag) == 0 {
		panic("Image.Tag is not set")
	}

	// Check if image is in registry
	re := regexp.MustCompile(`[^/]+\.[^/.]+/([^/.]+/)?[^/.]+(:.+)?`)
	if re.MatchString(c.Image.Tag) {

		// Remove old image
		exec.Command("docker", "image", "rm", "-f", c.Image.Tag, "||", "true").Run()

		// Docker pull command
		output, err := exec.Command("docker", "pull", c.Image.Tag).CombinedOutput()
		if err != nil {
			return fmt.Errorf(fmt.Sprint(err) + ": " + string(output))
		}
		fmt.Println(string(output))
	}

	return nil
}
