package Serializer

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
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

// For the ErrorResponse in the, it is a function that returns a Serializer.Response struct.
func ErrorResponse(err error) Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := T(fmt.Sprintf("Field.%s", e.Field))
			tag := T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return Response{
				Code:  400,
				Msg:   fmt.Sprintf("%s%s", field, tag),
				Error: fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return Response{
			Code:  400,
			Msg:   "JSON类型不匹配",
			Error: fmt.Sprint(err),
		}
	}

	return Response{
		Code:  400,
		Msg:   "参数错误",
		Error: fmt.Sprint(err),
	}
}

var Dictionary *map[interface{}]interface{}

// T 翻译
func T(key string) string {
	dic := *Dictionary
	keys := strings.Split(key, ".")
	for index, path := range keys {
		// 如果到达了最后一层，寻找目标翻译
		if len(keys) == (index + 1) {
			for k, v := range dic {
				if k, ok := k.(string); ok {
					if k == path {
						if value, ok := v.(string); ok {
							return value
						}
					}
				}
			}
			return path
		}
		// 如果还有下一层，继续寻找
		for k, v := range dic {
			if ks, ok := k.(string); ok {
				if ks == path {
					if dic, ok = v.(map[interface{}]interface{}); ok == false {
						return path
					}
				}
			} else {
				return ""
			}
		}
	}
	return ""
}
