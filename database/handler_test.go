package database

import (
	"github.com/ameniGa/timeTracker/config"
	hlp "github.com/ameniGa/timeTracker/helpers"
	td "github.com/ameniGa/timeTracker/testData/database"
	"testing"
)

func TestCreate(t *testing.T) {
	tableTest := td.CreateTTHandler()
	log := hlp.GetLogger()
	for _, testCase := range tableTest {
		conf := config.Database{
			Presence: testCase.FeedbackDb,
		}
		t.Run(testCase.Name, func(t *testing.T) {
			_, err := Create(&conf, log)
			if err != nil && !testCase.HasError {
				t.Errorf("expected success , got error: %v", err)
			}
			if err == nil && testCase.HasError {
				t.Error("expected error, got nil")
			}
		})
	}
}
