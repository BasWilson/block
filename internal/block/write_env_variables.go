package block

import (
	"io/ioutil"
	"os"
)

func WriteEnvVariablesToFile(c *Config) {
	var envString string

	for i := 0; i < len(c.Variables); i++ {
		envString += c.Variables[i].Name + "=" + c.Variables[i].Value + "\n"	
	}

	err := ioutil.WriteFile(c.Settings.Name + ".env", []byte(envString), 0644)
	if err != nil {
		panic(err)
	}
}


func DeleteEnvVariablesFile(c *Config) {
	os.Remove(c.Settings.Name + ".env")
}