package block

import "fmt"

type Image struct {
	Tag string `json:"tag"`
}

type Repo struct {
	Url      string `json:"url"`
	BuildCmd string `json:"build_cmd"`
	RunCmd   string `json:"run_cmd"`
}

type Settings struct {
	Port int16  `json:"port"`
	ContainerPort int16  `json:"container_port"`
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type EnvVariable struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type Config struct {
	PullMode string   `json:"pull_mode" binding:"required,oneof=repo image"`
	Image    Image    `json:"image" binding:"required,dive"`
	Repo     Repo     `json:"repo" binding:"dive"`
	Settings Settings `json:"settings" binding:"required,dive"`
	Variables []EnvVariable     `json:"variables" binding:"dive"`
}

var Configs []Config
var Base string

func AddConfig(c *Config) *Config {
	// Add config to list
	Configs = append(Configs, *c)
	
	// Write env variables to file
	WriteEnvVariablesToFile(c)

	fmt.Println("[BLOCK] Config added for: " + ": " + c.Settings.Name)
	return c
}

func RemoveConfig(c *Config) (Config, error) {
	for i, v := range Configs {
		if v.Settings.Name == c.Settings.Name {
			Configs = append(Configs[:i], Configs[i+1:]...)
			DeleteEnvVariablesFile(c)
			fmt.Println("[BLOCK] Config removed for: " + c.Settings.Name)
			return v, nil
		}
	}
	return Config{}, fmt.Errorf("[BLOCK] Config not found")
}