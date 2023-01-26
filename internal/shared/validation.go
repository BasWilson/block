package shared

import "runtime"

func ValidateOS() {
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		panic("Block can only execute on Linux & Darwin for now, your os: " + runtime.GOOS)
	}
}
