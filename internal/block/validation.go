package block

import "runtime"

func ValidateOS() {
	if runtime.GOOS != "linux" {
		panic("Block can only execute on Linux for now")
	}
}
