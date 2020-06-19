package models

type SendMsgInput struct {
	Text     string `json:"text,omitempty"`
	Channel  string `json:"channel,omitempty"`
	Username string `json:"username,omitempty"`
	Token    string `json:"token,omitempty"`
}
