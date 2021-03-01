package errors

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

// ResponseError ...
type ResponseError struct {
	error
	Code      int64                  `json:"code,omitempty"`
	Message   string                 `json:"message,omitempty"`
	MessageCn string                 `json:"message_cn,omitempty"`
	Errors    []model.ApiErrorStruct `json:"errors,omitempty"`
}

// Error ...
func (e ResponseError) Error() string {
	return e.Message
}

// NewError ...
func NewError(code *int64, message *string, messageCn *string, errors []model.ApiErrorStruct) ResponseError {
	var codeValue int64
	var messageValue string
	var messageCnValue string
	if code != nil {
		codeValue = *code
	}
	if message != nil {
		messageValue = *message
	}
	if messageCn != nil {
		messageCnValue = *messageCn
	}

	return ResponseError{
		Code:      codeValue,
		Message:   messageValue,
		MessageCn: messageCnValue,
		Errors:    errors,
	}
}
