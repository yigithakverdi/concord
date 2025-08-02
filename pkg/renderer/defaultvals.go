package renderer

type Data struct {
	Config ConfigType `yaml:"config"`
}

type ConfigType struct {
	Server  map[string]interface{} `yaml:"server"`
	Logging map[string]interface{} `yaml:"logging"`
	// Management map[string] `yaml:"management"`
}
