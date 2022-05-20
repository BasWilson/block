package block

import "os"

func Env() string {
	return os.Getenv("BLOCK_ENV")
}
