package configuration

import (
	log "github.com/sirupsen/logrus"
)

var (
	defaultLogLevel log.Level = log.DebugLevel
)

// InitLogger ...
func InitLogger(ll string) {
	logLevel, err := log.ParseLevel(ll)
	if err != nil {
		Logger.SetLevel(defaultLogLevel)
		Logger.Errorf("Invalid Loglevel '%s' passed as argument, Defaulting to '%v' logging ", ll, defaultLogLevel)
	}
	Logger.SetLevel(logLevel)
}

// Logger ...
var Logger = log.New()
