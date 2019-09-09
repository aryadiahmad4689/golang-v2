package config

import (
	"github.com/tkanos/gonfig"
)

// Configuration ..
type Configuration struct {
	dbUsername string
	dbPassword string
	dbHost     string
	dbName     string
}

// GetConfig ..
func GetConfig() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/env.json", &configuration)
	return configuration
}
