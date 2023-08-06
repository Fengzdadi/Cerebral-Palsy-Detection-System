package WsApi

import (
	"Cerebral-Palsy-Detection-System/WS/Conf"
	_ "Cerebral-Palsy-Detection-System/WS/Conf"
	"Cerebral-Palsy-Detection-System/WS/Serializer"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}

func MyErrorResponse(message string) error {
	return &MyError{Message: message}
}

// For the ErrorResponse in the WsApi\user.go, it is a function that returns a Serializer.Response struct.
func ErrorResponse(err error) Serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := Conf.T(fmt.Sprintf("Field.%s", e.Field))
			tag := Conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return Serializer.Response{
				Code:  400,
				Msg:   fmt.Sprintf("%s%s", field, tag),
				Error: fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return Serializer.Response{
			Code:  400,
			Msg:   "JSON类型不匹配",
			Error: fmt.Sprint(err),
		}
	}

	return Serializer.Response{
		Code:  400,
		Msg:   "参数错误",
		Error: fmt.Sprint(err),
	}
}
