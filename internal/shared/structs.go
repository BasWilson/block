package shared

type Image struct {
	Tag string `json:"tag"`
}

type Repo struct {
	Url      string `json:"url"`
	BuildCmd string `json:"build_cmd"`
	RunCmd   string `json:"run_cmd"`
}

type Settings struct {
	Port          int16  `json:"port"`
	ContainerPort int16  `json:"container_port"`
	Name          string `json:"name" binding:"required"`
	Type          string `json:"type" binding:"required"`
}

type LoadManager struct {
	MachineType       string `json:"machineType"`       // low, medium, high
	MaxInstances      int    `json:"maxInstances"`      // max number of instances
	MinInstances      int    `json:"minInstances"`      // min number of instances
	UnitsPerInstance  int    `json:"unitsPerInstance"`  // units per instance
	AutoScaleDown     bool   `json:"autoScaleDown"`     // auto scale down
	AutoScaleUp       bool   `json:"autoScaleUp"`       // auto scale up
	TerminationScript string `json:"terminationScript"` // script to run when terminating instance
}

type EnvVariable struct {
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type Config struct {
	PullMode  string        `json:"pull_mode" binding:"required,oneof=repo image"`
	Image     Image         `json:"image" binding:"required,dive"`
	Repo      Repo          `json:"repo" binding:"dive"`
	Settings  Settings      `json:"settings" binding:"required,dive"`
	Variables []EnvVariable `json:"variables" binding:"dive"`
}
