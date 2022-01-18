package models

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Qq       string `json:"qq"`
}

type UserInfo struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Role     string `json:"role"`
	Phone    string `json:"phone"`
	Token    string `json:"token"`
	Qq       string `json:"qq"`
}
