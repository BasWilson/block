package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/baswilson/block/internal/block"
	"github.com/go-co-op/gocron"
)

func main() {
	block.ValidateOS()
	fmt.Println("[BLOCK] running monitor...")
	s := gocron.NewScheduler(time.Local)
	s.Every(30).Seconds().Do(healthCheck)
	s.StartBlocking()
}

func healthCheck() {
	fmt.Println("Checking health")
	_, err := http.Get("http://127.0.0.1:80")
	if err != nil {
		fmt.Println(err)
		// Report error to API
		// TODO
	}

	// Image mode
	if block.ReadConfig.PullMode == "image" {
		output, err := exec.Command(
			"docker",
			"logs",
			"--since=30s",
			block.ReadConfig.Settings.Name).CombinedOutput()

		if err != nil {
			panic(fmt.Sprint(err) + ": " + string(output))
		}

		fmt.Println(string(output))
	}

}
