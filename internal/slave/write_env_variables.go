package slave

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/baswilson/block/internal/shared"
)

func WriteEnvVariablesToFile(c *shared.Config) {
	var envString string

	for i := 0; i < len(c.Variables); i++ {
		envString += c.Variables[i].Name + "=" + c.Variables[i].Value + "\n"
	}

	err := ioutil.WriteFile(path.Join(Base, c.Settings.Name+".env"), []byte(envString), 0644)
	if err != nil {
		panic(err)
	}
}

func DeleteEnvVariablesFile(c *shared.Config) {
	os.Remove(path.Join(Base, c.Settings.Name+".env"))
}
