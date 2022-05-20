package block

import (
	"fmt"
	"os"
	"os/exec"
)

func Run() {
	fmt.Println("[BLOCK] Running image")

	// Validate image tag
	BLOCK_IMAGE_TAG := os.Getenv("BLOCK_IMAGE_TAG")
	if len(BLOCK_IMAGE_TAG) == 0 {
		panic("BLOCK_IMAGE_TAG is not set")
	}

	// Docker run command
	output, err := exec.Command(
		"docker",
		"run",
		"--env-file", "build/ci/dev/sample.env",
		BLOCK_IMAGE_TAG).CombinedOutput()

	if err != nil {
		panic(fmt.Sprint(err) + ": " + string(output))
	}

	fmt.Println(string(output))
}
