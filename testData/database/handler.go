package database

import "github.com/ameniGa/timeTracker/config"

// TTCreateHandler represents table test structure of CreateHandler test
type TTCreateHandler struct {
	Name         string
	FeedbackDb   config.Presence
	HasError     bool
}

// CreateTTHandler creates table test for CreateHandler test
func CreateTTHandler() []TTCreateHandler {
	testTableName := "dynamodb"
	return []TTCreateHandler{
		{
			Name:         "valid config",
			FeedbackDb:   config.Presence{Type: testTableName},
			HasError:     false,
		},
		{
			Name:         "unsupported db ",
			FeedbackDb:   config.Presence{Type: "unknown"},
			HasError:     true,
		},
	}
}


