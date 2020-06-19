package slack_test

import (
	. "github.com/ameniGa/timeTracker/notification/slack"
	td "github.com/ameniGa/timeTracker/testData/slack"
	"testing"
)

func Test_SendMessage(t *testing.T) {
	for _, tc := range td.CreateTTSendMsg() {
		t.Run(tc.Name, func(t *testing.T) {
			slack := NewSlackHandler(tc.Conf)
			err := slack.SendMessage(tc.Channel, tc.Message, tc.Username)
			if err != nil && !tc.HasError {
				t.Errorf("expected success got error: %v", err)
			}
			if err == nil && tc.HasError {
				t.Errorf("expected failure got nil")
			}
		})
	}
}
