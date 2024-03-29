package slave

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
)

func StartHealthMonitoring() {
	s := gocron.NewScheduler(time.Local)
	s.Every(30).Seconds().Do(healthCheck)
	s.StartBlocking()
}

// Checks if the webservice returns a response.
// Result gets reported to API
func healthCheck() {
	c := Config
	if c.Settings.Type != "webservice" {
		return
	}

	fmt.Println("[BLOCK] Checking health for " + c.Settings.Name)
	_, err := http.Get("http://" + c.Settings.Name)
	if err != nil {
		fmt.Println("Webservice is offline. " + err.Error())
		// Report error to API
		// TODO
	} else {
		// Report success to API
		// TODO
		fmt.Println("Webservice is online")
	}
}
