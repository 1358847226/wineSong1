package models

type LoginLog struct {
	Id       int    `json:"id"`
	UserId   int    `json:"userId"`
	Time     string `json:"time"`
	Ip       string `json:"ip"`
	Location string `json:"location"`
}
