package block

import (
	"fmt"
	"os"
	"os/exec"
)

func Pull() {
	fmt.Println("[BLOCK] Pulling image from")

	// Check for pull mode
	BLOCK_PULL_MODE := os.Getenv("BLOCK_PULL_MODE")
	if BLOCK_PULL_MODE == "repo" {
		panic("BLOCK_PULL_MODE 'repo' is not yet supported.")
	} else if BLOCK_PULL_MODE == "image" {
		image()
	}
}

func image() {
	// Validate image tag
	BLOCK_IMAGE_TAG := os.Getenv("BLOCK_IMAGE_TAG")
	if len(BLOCK_IMAGE_TAG) == 0 {
		panic("BLOCK_IMAGE_TAG is not set")
	}

	// Docker pull command
	output, err := exec.Command("docker", "pull", BLOCK_IMAGE_TAG).CombinedOutput()
	if err != nil {
		panic(fmt.Sprint(err) + ": " + string(output))
	}
	fmt.Println(string(output))
}
