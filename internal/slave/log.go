package slave

import (
	"fmt"
	"os/exec"
)

func Log(containerName string) (string, error) {
	fmt.Println("[BLOCK] Logging container")

	// Check for pull mode
	output, err := exec.Command("docker", "logs", "--timestamps", containerName).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf(fmt.Sprint(err) + ": " + string(output))
	}

	return string(output), nil
}
