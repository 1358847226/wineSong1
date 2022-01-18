package entity

type Result struct {
	Code    int         `json:"code"`
	Count   int         `json:"count"`
	Message string      `json:"message"`
	Error   error       `json:"error"`
	Data    interface{} `json:"data"`
}
