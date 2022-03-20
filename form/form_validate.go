package form

import (
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

type FormRequest interface {
	Messages() ValidatorMessages
}

type ValidatorMessages map[string]string

func ParseError(request FormRequest, err error) string {
	for _, v := range err.(validator.ValidationErrors) {
		if message, exist := request.Messages()[v.Field()+"."+v.Tag()]; exist {
			return message
		}

		return v.Error()
	}

	return "Parameter error"
}

func ParseErrors(request FormRequest, err error) map[string][]string {
	data := map[string][]string{}

	for _, v := range err.(validator.ValidationErrors) {
		errField := strcase.ToSnake(v.Field())
		errMsg := v.Error()

		if message, exist := request.Messages()[v.Field()+"."+v.Tag()]; exist {
			errMsg = message
		}

		if _, ok := data[errField]; ok {
			data[errField] = append(data[errField], errMsg)
		} else {
			data[errField] = []string{errMsg}
		}
	}

	return data
}
