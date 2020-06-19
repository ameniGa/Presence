package models

type User struct {
	UserID    string `json:"UserID"`
	UserName  string `json:"UserName"`
	Password  string `json:"Password"`
	CreatedAt uint64 `json:"Password"`
}
