package config

import (
	"github.com/jinzhu/configor"
	"path"
	"runtime"
	"time"
)

const ConfEnvVar = "CONFIGOR_ENV"

type Server struct {
	Type     string
	Host     string        `default:"localhost" env:"SERVER_HOST"`
	Port     string        `default:"50077" env:"SERVER_PORT"`
	Deadline time.Duration `default:"5" env:"GRPC_DEADLINE"`
}

type Presence struct {
	Type      string
	TableName string `env:"Presence_DB_TABLE_NAME"`
}
type Database struct {
	Presence Presence
}
type Camera struct {
	DeviceID int
}
type Config struct {
	Tag      string // indicates the config environment prod or dev
	Server   Server
	Database Database
	Camera   Camera
}

// LoadConfig sets the application config
// uses CONFIGOR_ENV to set environment,
// if CONFIGOR_ENV not set, environment will be production by default
// and it will be test when running tests with go test
// otherwise it can be set to test manually
func LoadConfig() (*Config, error) {
	var configFilePath string
	config := configor.New(&configor.Config{})
	switch config.GetEnvironment() {
	case "test", "development":
		configFilePath = "./config.development.yml"
	default:
		configFilePath = "./config.production.yml"
	}
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), configFilePath)
	conf := new(Config)
	err := config.Load(conf, filepath)
	return conf, err
}
