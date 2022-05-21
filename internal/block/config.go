package block

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Image struct {
	Tag string
}

type Repo struct {
	Url      string `json:"url"`
	BuildCmd string `json:"build_cmd"`
	RunCmd   string `json:"run_cmd"`
}

type Settings struct {
	Port int16  `json:"port"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Config struct {
	PullMode string   `json:"pull_mode"`
	Image    Image    `json:"image"`
	Repo     Repo     `json:"repo"`
	Settings Settings `json:"settings"`
}

var ReadConfig Config

func ParseConfig() Config {
	jsonFile, err := os.Open("configs/example.block.config.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &ReadConfig)

	return ReadConfig
}
