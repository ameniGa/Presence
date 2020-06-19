package database

import (
	"github.com/ameniGa/timeTracker/config"
	faker "github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"time"
)

// valid and invalid config
var (
	validConfig, invalidConfig *config.Presence
	Conf                       *config.Config
)

func init() {
	Conf, _ = config.LoadConfig()
	validConfig = &Conf.Database.Presence
	invalidConfig = &config.Presence{
		UserTableName: "unknown-table",
		TimeTableName: "unknown-table",
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
			Timeout:  Conf.Server.Deadline,
			Username: faker.Username(),
			UserID:   uuid.New().String(),
			HasError: false,
		},
		{
			Name:     "invalid Request: invalid config",
			DBConf:   invalidConfig,
			Timeout:  Conf.Server.Deadline,
			Username: faker.Username(),
			UserID:   uuid.New().String(),
			HasError: true,
		},
		{
			Name:     "invalid Request: invalid Username",
			DBConf:   validConfig,
			Timeout:  Conf.Server.Deadline,
			Username: "",
			UserID:   uuid.New().String(),
			HasError: true,
		},
		{
			Name:     "invalid Request: invalid timeout",
			DBConf:   validConfig,
			Timeout:  0,
			Username: faker.Username(),
			UserID:   uuid.New().String(),
			HasError: true,
		},
		{
			Name:     "invalid Request: invalid id",
			DBConf:   validConfig,
			Timeout:  Conf.Server.Deadline,
			Username: faker.Username(),
			UserID:   "invalid id ",
			HasError: true,
		},
	}
	return data
}

type TTAddInOut struct {
	Name     string
	DBConf   *config.Presence
	Timeout  time.Duration
	UserID   string
	HasError bool
}

func CreateTTAddInOut() ([]TTAddInOut, []TTAddInOut) {
	validUserID := uuid.New().String()
	in := []TTAddInOut{
		{
			Name:     "Valid Request",
			DBConf:   validConfig,
			Timeout:  Conf.Server.Deadline,
			UserID:   validUserID,
			HasError: false,
		},
		{
			Name:     "invalid Request: invalid Conf ",
			DBConf:   invalidConfig,
			Timeout:  Conf.Server.Deadline,
			UserID:   validUserID,
			HasError: true,
		},
		{
			Name:     "invalid Request: missing id",
			DBConf:   validConfig,
			Timeout:  Conf.Server.Deadline,
			UserID:   "",
			HasError: true,
		},
		{
			Name:     "invalid Request: invalid id",
			DBConf:   validConfig,
			Timeout:  Conf.Server.Deadline,
			UserID:   "invalid id",
			HasError: true,
		},
		{
			Name:     "invalid Request: invalid timeout",
			DBConf:   validConfig,
			Timeout:  1,
			UserID:   validUserID,
			HasError: true,
		},
	}
	out := []TTAddInOut{
		{
			Name:     "Valid Request",
			DBConf:   validConfig,
			Timeout:  Conf.Server.Deadline,
			UserID:   validUserID,
			HasError: false,
		},
		{
			Name:     "invalid Request: invalid Conf ",
			DBConf:   invalidConfig,
			Timeout:  Conf.Server.Deadline,
			UserID:   validUserID,
			HasError: true,
		},
		{
			Name:     "invalid Request: missing id",
			DBConf:   validConfig,
			Timeout:  Conf.Server.Deadline,
			UserID:   "",
			HasError: true,
		},
		{
			Name:     "invalid Request: invalid timeout",
			DBConf:   validConfig,
			Timeout:  1,
			UserID:   validUserID,
			HasError: true,
		},
		{
			Name:     "invalid Request: invalid id",
			DBConf:   validConfig,
			Timeout:  Conf.Server.Deadline,
			UserID:   "invalidId",
			HasError: true,
		},
	}
	return in, out
}

type TTGetUser struct {
	Name         string
	DBConf       *config.Presence
	Timeout      time.Duration
	UserID       string
	ExpectedName string
	HasError     bool
}

func CreateTTGetUser(validUserID string) []TTGetUser {
	return []TTGetUser{
		{
			Name:         "Valid Request",
			DBConf:       validConfig,
			Timeout:      Conf.Server.Deadline,
			UserID:       validUserID,
			ExpectedName: "Test",
			HasError:     false,
		},
		{
			Name:         "invalid Request: invalid config",
			DBConf:       invalidConfig,
			Timeout:      Conf.Server.Deadline,
			UserID:       validUserID,
			ExpectedName: "Test",
			HasError:     true,
		},
		{
			Name:         "invalid Request: invalid id",
			DBConf:       validConfig,
			Timeout:      Conf.Server.Deadline,
			UserID:       "invalidUserID",
			ExpectedName: "Test",
			HasError:     true,
		},
		{
			Name:         "invalid context timeout",
			DBConf:       validConfig,
			Timeout:      0,
			UserID:       validUserID,
			ExpectedName: "Test",
			HasError:     true,
		},
	}
}
