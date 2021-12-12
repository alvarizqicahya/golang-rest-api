package helper

import (
	"strings"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Status:  true,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err string) Response {
	splitError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splitError,
	}
	return res
}
