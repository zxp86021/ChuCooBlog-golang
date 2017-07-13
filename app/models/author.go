package models

type Author struct {
	Username string `json:"username"`
	Password string `json:"-"` // password 不 response 出去
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Address  string `json:"address"`
}