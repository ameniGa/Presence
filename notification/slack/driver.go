package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ameniGa/timeTracker/config"
	mdl "github.com/ameniGa/timeTracker/models"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type Slack struct {
	Conf   config.Slack
	client *http.Client
}

func NewSlackHandler(conf config.Slack) Slack {
	return Slack{
		Conf:   conf,
		client: &http.Client{},
	}
}

func (s Slack) SendMessage(channel, message, userName string) error {
	if channel == "" {
		return errors.New("please specify the channel")
	}
	if message == "" {
		return errors.New("please specify a message")
	}

	body := mdl.SendMsgInput{
		Text:     message,
		Channel:  channel,
		Username: userName,
		Token:    s.Conf.ApiToken,
	}
	jsonValue, _ := json.Marshal(body)

	request, _ := http.NewRequest("POST", s.Conf.Apis["send"], bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", s.Conf.ApiToken))

	response, err := s.client.Do(request)
	if err != nil {
		return errors.Wrap(err, "failed to send message")
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "failed to retrieve response")
	}
	return parseResponse(responseData)
}

func parseResponse(responseData []byte) error {
	var res interface{}
	err := json.Unmarshal(responseData, &res)
	if err != nil {
		return errors.Wrap(err, "failed to get the response")
	}
	output := res.(map[string]interface{})
	val, ok := output["error"]
	if !ok {
		return nil
	}
	return errors.New(val.(string))
}
