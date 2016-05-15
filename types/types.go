package types

import (
	"encoding/json"
)

type CRUDError struct {
	//ErrorCode int
	ErrorMessage string
}

func (crudError CRUDError) Error() string {
	return crudError.ErrorMessage
}

type HttpResult struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func (httpResult *HttpResult) JsonString() string {
	str, _ := json.Marshal(httpResult)
	return string(str)
}