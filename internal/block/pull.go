package block

import (
	"fmt"
	"os/exec"
)

func Pull() {
	fmt.Println("[BLOCK] Pulling image from")

	// Check for pull mode
	if ReadConfig.PullMode == "repo" {
		panic("PullMode 'repo' is not yet supported.")
	} else if ReadConfig.PullMode == "image" {
		image()
	}
}

func image() {
	// Validate image tag
	if len(ReadConfig.Image.Tag) == 0 {
		panic("Image.Tag is not set")
	}

	// Docker pull command
	output, err := exec.Command("docker", "pull", ReadConfig.Image.Tag).CombinedOutput()
	if err != nil {
		panic(fmt.Sprint(err) + ": " + string(output))
	}
	fmt.Println(string(output))
}
