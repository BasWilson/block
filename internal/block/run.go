package block

import (
	"fmt"
	"os/exec"
	"strconv"
)

func Run() {
	fmt.Println("[BLOCK] Running image")

	// Validate image tag
	if len(ReadConfig.Image.Tag) == 0 {
		panic("Settings.Image.Tag is not set")
	}

	// Remove old instance
	exec.Command(
		"docker",
		"rm",
		"-f",
		ReadConfig.Settings.Name).CombinedOutput()

	// Docker run command
	output, err := exec.Command(
		"docker",
		"run",
		"-d",
		"--name", ReadConfig.Settings.Name,
		"-p", "80:"+strconv.Itoa(int(ReadConfig.Settings.Port)),
		"--env-file", "build/ci/dev/sample.env",
		"--restart=always",
		"--network=block_network",
		ReadConfig.Image.Tag).CombinedOutput()

	if err != nil {
		panic(fmt.Sprint(err) + ": " + string(output))
	}

	fmt.Println(string(output))
}
