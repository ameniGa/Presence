package database

import (
	"errors"
	cfg "github.com/ameniGa/timeTracker/config"
	"github.com/ameniGa/timeTracker/database/dynamo"
	"github.com/sirupsen/logrus"
)

// PresenceHandler holds functions related to presence database
type PresenceHandler interface {

}

// Create creates db handler based on the given config
func Create(conf *cfg.Database, logger *logrus.Logger) (PresenceHandler, error) {
	var handler PresenceHandler
	switch conf.Presence.Type {
	case "dynamodb":
		handler = dynamo.NewPresenceHandler(&conf.Presence, logger)
	default:
		return nil, errors.New("Invalid DB Type")
	}
	return handler, nil
}
