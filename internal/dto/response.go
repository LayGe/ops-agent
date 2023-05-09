package dto

type Response struct {
	Code    int         `json:"code"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}
