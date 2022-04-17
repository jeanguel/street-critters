package config

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Application struct {
		Environment  string `toml:"environment"`
		LogDirectory string `toml:"log-directory"`
	} `toml:"application"`
	Server struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
	} `toml:"api-server"`
	Database struct {
		PSQLConnectionString string `toml:"psql-connection-string"`
		TLSCert              string `toml:"tls-cert"`
	} `toml:"database"`
}

// Config
var Config Configuration

func readConfigFile() {
	// Default project config file path
	configPath := "./sc.config.toml"
	envConfigPath := os.Getenv("STREET_CRITTERS_CONFIG_PATH")

	// The default config path can be overridden by indicating the
	//  config path in the environment variable
	if len(envConfigPath) > 0 {
		configPath = envConfigPath
	}

	c, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if _, err := toml.Decode(string(c), &Config); err != nil {
		panic(err)
	}
}
