package main

import (
	"fmt"
	"time"

	block "github.com/baswilson/block/internal/block"
	"github.com/go-co-op/gocron"
)

func main() {
	block.ValidateOS()
	fmt.Println("[BLOCK] initializing block")
	block.ParseConfig()
	block.Pull()
	block.Run()

	fmt.Println("[BLOCK] running monitor...")
	s := gocron.NewScheduler(time.Local)
	s.Every(30).Seconds().Do(block.HealthCheck)
	s.StartBlocking()
}
