package globals

const (
	DefaultPropertiesFileLocation = "./application.properties"
	DefaultValuesFileLocation     = "./base.values.yaml"
	DefaultValuesFile             = "./default.values.yaml"
	DefaultPropertiesFile         = "./default.properties.yaml"
	DefaultEnvPropertiesFile      = "./default.env.properties"
	CurrentEnvironment            = "dev"
)

// Global map to hold initial properties both application and base values,
// later on this global map will be used to lint and validate the properties
// and values.
var GlobalProperties = make(map[string]string)
