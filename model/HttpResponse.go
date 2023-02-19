package model

type HttpResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	CSRF    interface{} `json:"_csrf"`
}