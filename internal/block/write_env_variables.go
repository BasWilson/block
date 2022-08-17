package block

import (
	"io/ioutil"
	"os"
	"path"
)

func WriteEnvVariablesToFile(c *Config) {
	var envString string

	for i := 0; i < len(c.Variables); i++ {
		envString += c.Variables[i].Name + "=" + c.Variables[i].Value + "\n"	
	}

	err := ioutil.WriteFile(path.Join(Base, c.Settings.Name + ".env"), []byte(envString), 0644)
	if err != nil {
		panic(err)
	}
}


func DeleteEnvVariablesFile(c *Config) {
	os.Remove(path.Join(Base, c.Settings.Name + ".env"))
}