package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := map[string]string{}

	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			res[v.Field()] = Translatetag(v)
		}
	}
	return res
}

func Translatetag(fd validator.FieldError) string {
	switch fd.ActualTag() {
	case "required":
		return fmt.Sprintf(" is required %s Wajib di isi", fd.StructField())
	case "main":
		return fmt.Sprintf("field %s minimal %s", fd.StructField(), fd.Param())
	case "unique":
		return fmt.Sprintf("field %s harus unik", fd.StructField())
	}
	return "validation failed"
}
