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
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type Config struct {
	PullMode string   `json:"pull_mode" binding:"required,oneof=repo image"`
	Image    Image    `json:"image" binding:"required,dive"`
	Repo     Repo     `json:"repo" binding:"dive"`
	Settings Settings `json:"settings" binding:"required,dive"`
}

var Configs []Config

func AddConfig(c *Config) *Config {
	// Add config to list
	Configs = append(Configs, *c)
	fmt.Println("[BLOCK] Config added", c)
	return c
}

func RemoveConfig(c *Config) (Config, error) {
	for i, v := range Configs {
		if v.Settings.Name == c.Settings.Name {
			Configs = append(Configs[:i], Configs[i+1:]...)
			fmt.Println("[BLOCK] Config removed", c)
			return v, nil
		}
	}
	return Config{}, fmt.Errorf("[BLOCK] Config not found")
}