package errors

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type ResponseError struct {
	error
	Code      int64                  `json:"code,omitempty"`
	Message   string                 `json:"message,omitempty"`
	MessageCn string                 `json:"message_cn,omitempty"`
	Errors    []model.ApiErrorStruct `json:"errors,omitempty"`
}

func (e ResponseError) Error() string {
	return e.Message
}

func NewError(code int64, message string, messageCn string, errors []model.ApiErrorStruct) ResponseError {
	return ResponseError{
		Code:      code,
		Message:   message,
		MessageCn: messageCn,
		Errors:    errors,
	}
}
