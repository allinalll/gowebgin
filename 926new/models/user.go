package models

type User struct {
	Id       int    `json:"userallId"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
