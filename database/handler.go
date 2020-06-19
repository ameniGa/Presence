package database

import (
	"context"
	"errors"
	cfg "github.com/ameniGa/timeTracker/config"
	"github.com/ameniGa/timeTracker/database/dynamo"
	"github.com/sirupsen/logrus"
	"time"
)

// PresenceHandler holds functions related to presence database
type PresenceHandler interface {

	// DbAddUser creates a user for the given id and with the given username and save it in the database
	DbAddUser(ctx context.Context, userID, userName string, ch chan<- error)

	// DbAddEntry creates a new entry for the given userId save it in the database
	DbAddEntry(ctx context.Context, userID string, ch chan<- error)

	// DbAddExit updates a exit time for the given userId save it in the database
	DbAddExit(ctx context.Context, userID string, ch chan<- error)
}

// Create creates db handler based on the given config
func Create(conf *cfg.Database, timeout time.Duration, logger *logrus.Logger) (PresenceHandler, error) {
	var handler PresenceHandler
	switch conf.Presence.Type {
	case "dynamodb":
		handler = dynamo.NewPresenceHandler(&conf.Presence, timeout, logger)
	default:
		return nil, errors.New("Invalid DB Type")
	}
	return handler, nil
}
