package config

import (
	"fmt"
	"os"

	log "github.com/gothew/l-og"
)

const (
	mongoDBURIStr = "mongodb://%s:%s@%s"
)

var (
	MongoDBURI string
)

type Config struct {
	logger *log.Logger
}

func NewConfig() Config {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		Level: log.DebugLevel,
	})
	return Config{logger: logger}
}

func (c *Config) InitializeAppConfig() {
	// mongo db config
	dbServer := os.Getenv("MONGODB_SERVER")
	dbUsername, dbPassword := os.Getenv("MONGODB_USERNAME"), os.Getenv("MONGODB_PASSWORD")
	MongoDBURI = fmt.Sprintf(mongoDBURIStr, dbUsername, dbPassword, dbServer)
	c.logger.Infof("mongodb server URI is : %s", dbServer)
}

func (c *Config) Logger() *log.Logger {
	return c.logger
}
