package config

import (
	"github.com/tkanos/gonfig"
)

// Configuration ..
type Configuration struct {
	DbUsername string
	DbPassword string
	DbHost     string
	DbName     string
}

// GetConfig ..
func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/env.json", &configuration)
	return configuration
}
