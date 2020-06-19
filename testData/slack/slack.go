package slack

import "github.com/ameniGa/timeTracker/config"

type TTSendMsg struct {
	Name     string
	Channel  string
	Message  string
	Username string
	Conf     config.Slack
	HasError bool
}

func CreateTTSendMsg() []TTSendMsg {
	validConf, _ := config.LoadConfig()
	return []TTSendMsg{
		{
			Name:    "valid sending",
			Conf:    validConf.Notification.Slack,
			Channel: "presence",
			Message: "Good Morning",
		},
		{
			Name: "invalid api url",
			Conf: config.Slack{
				Apis:     map[string]string{"send": "balala"},
				ApiToken: validConf.Notification.Slack.ApiToken,
			},
			Channel:  "presence",
			Message:  "Good Morning",
			HasError: true,
		},
		{
			Name: "invalid api token",
			Conf: config.Slack{
				Apis:     validConf.Notification.Slack.Apis,
				ApiToken: "blabla",
			},
			Channel:  "presence",
			Message:  "Good Morning",
			HasError: true,
		},
		{
			Name:     "invalid channel",
			Conf:     config.Slack{},
			Channel:  "blabla",
			Message:  "Good Morning",
			HasError: true,
		},
		{
			Name:     "empty channel",
			Conf:     validConf.Notification.Slack,
			Channel:  "",
			Message:  "Good Morning",
			HasError: true,
		},
		{
			Name:     "empty message",
			Conf:     validConf.Notification.Slack,
			Channel:  "presence",
			Message:  "",
			HasError: true,
		},
		{
			Name:     "send private message",
			Conf:     validConf.Notification.Slack,
			Channel:  "presence",
			Message:  "Good morning",
			Username: "amani",
		},
	}
}
