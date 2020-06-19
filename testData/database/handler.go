package database

import "github.com/ameniGa/timeTracker/config"

// TTCreateHandler represents table test structure of CreateHandler test
type TTCreateHandler struct {
	Name     string
	Presence config.Presence
	HasError bool
}

// CreateTTHandler creates table test for CreateHandler test
func CreateTTHandler() []TTCreateHandler {
	testTableName := "dynamodb"
	return []TTCreateHandler{
		{
			Name:     "valid config",
			Presence: config.Presence{Type: testTableName},
			HasError: false,
		},
		{
			Name:     "unsupported db ",
			Presence: config.Presence{Type: "unknown"},
			HasError: true,
		},
	}
}
