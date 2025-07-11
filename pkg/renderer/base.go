package springboot

// Base configuration of the properties file is represented as the below struct.
// configurations are filled in to the .template config file as a map[string]string
//
// Base struct is divided into several sections, each section corresponds to a specific configuration category.
// The sections are:
// - Logging: Contains configurations related to logging.
// - Server related configs (Tomcat etc.)
// - Security related configs (Keystore, Truststore etc.)
//

type Logging struct {
}

type Server struct {
}

type Security struct {
}

type Base struct {
	Logging    Logging           `yaml:"logging"`
	Server     Server            `yaml:"server"`
	Security   Security          `yaml:"security"`
	Properties map[string]string `yaml:"properties"`
}
