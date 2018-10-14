package parser

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

// LogFormat : Format specification for error log entries
type LogFormat struct {
	Time     time.Time
	Loglevel string
	Pid      int
	Tid      int
	Source   string
	Apr      string
	Client   string
	Message  string
	Request  string
}

// ConfigFile : Define path and filename (without extension) of config file
type ConfigFile struct {
	Path     string
	Filename string
}

// ConfigContent : Entries in config file
type ConfigContent struct {
	One string
	Two string
}

// C : Stores ConfigContent
var C ConfigContent

// Config : Handle configuration file
func Config(path, filename string) {
	config := ConfigFile{path, filename}
	viper.AddConfigPath(config.Path)
	viper.SetConfigName(config.Filename)

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Get config entry from Configuration file
	json := viper.Sub("config")

	err := json.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
