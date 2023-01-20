package shared

import (
	"os"
	"strings"
)

var Base string

func SetBasePath() {
	executable, _ := os.Executable()
	Base = strings.Split(executable, "/bin/master")[0]
}
