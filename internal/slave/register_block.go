package slave

import (
	"fmt"
	"net/http"
	"os"
)

func RegisterBlock() {
	endpoint := os.Getenv("BLOCK_REGISTRATION_ENDPOINT")
	if len(endpoint) == 0 {
		fmt.Println("[BLOCK] Not registering block, no BLOCK_REGISTRATION_ENDPOINT set")
		return
	}

	fmt.Println("[BLOCK] Registering")

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("[BLOCK] Registration complete", res.Status)
}
