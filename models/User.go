package models

type User struct {
	UserID      string `json:"UserID"`
	UserName    string `json:"UserName"`
	Password    string `json:"Password"`
	CreatedAt   uint64 `json:"CreatedAt"`
	PassChanged bool   `json:"PassChanged"`
}

type UserInfo struct {
	UserID      string `json:"UserID"`
	UserName    string `json:"UserName"`
	Password    string `json:"Password"`
	CreatedAt   uint64 `json:"CreatedAt"`
	PassChanged bool   `json:"PassChanged"`
}

type UserWithError struct {
	UserInfo UserInfo
	Error    error
}

type AuthUserInput struct {
	UserID   string
	Password string
}

type GenericRes struct {
	Status string
	Error  error
}
