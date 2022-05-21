package block

import (
	"fmt"
	"net/http"
	"os/exec"
)

func HealthCheck() {
	fmt.Println("Checking health")
	_, err := http.Get("https://" + ReadConfig.Settings.Name)
	if err != nil {
		fmt.Println(err)
		// Report error to API
		// TODO
	}

	// Image mode
	if ReadConfig.PullMode == "image" {
		output, err := exec.Command(
			"docker",
			"logs",
			"--since=30s",
			ReadConfig.Settings.Name).CombinedOutput()

		if err != nil {
			panic(fmt.Sprint(err) + ": " + string(output))
		}

		fmt.Println(string(output))
	}

}
