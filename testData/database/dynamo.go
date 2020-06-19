package database

import (
	"github.com/ameniGa/timeTracker/config"
	faker "github.com/bxcodec/faker/v3"
	"time"
)

// valid and invalid config
var (
	validConfig, invalidConfig *config.Presence
	conf                       *config.Config
)

func init() {
	conf, _ = config.LoadConfig()
	validConfig = &conf.Database.Presence
	invalidConfig = &config.Presence{
		TableName: "unknown-table",
	}

}

type TTAddUser struct {
	Name     string
	DBConf   *config.Presence
	Timeout  time.Duration
	UserID   string
	Username string
	HasError bool
}

func CreateTTAddUser() []TTAddUser {
	data := []TTAddUser{
		{
			Name:     "Valid Request",
			DBConf:   validConfig,
			Timeout:  conf.Server.Deadline,
			Username: faker.Username(),
			UserID:   faker.UUIDDigit(),
			HasError: false,
		},
		{
			Name:     "invalid Request: invalid config",
			DBConf:   invalidConfig,
			Timeout:  conf.Server.Deadline,
			Username: faker.Username(),
			UserID:   faker.UUIDDigit(),
			HasError: true,
		},
		{
			Name:     "invalid Request: invalid Username",
			DBConf:   validConfig,
			Timeout:  conf.Server.Deadline,
			Username: "",
			UserID:   faker.UUIDDigit(),
			HasError: true,
		},
	}
	return data
}
