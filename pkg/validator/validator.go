package validator

import (
	"strings"

	"github.com/Ajulll22/belajar-microservice/pkg/formatter"
	"github.com/go-playground/validator/v10"
)

type ErrorValidator struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "can not be empty!"
	case "required_without":
		return "can not be empty if " + formatter.ToSnakeCase(fe.Param()) + " is not present!"
	case "max":
		return "length maximum " + formatter.ToSnakeCase(fe.Param()) + "!"
	case "min":
		return "length minimum " + formatter.ToSnakeCase(fe.Param()) + "!"
	case "gte":
		return "should be greater than or equal to " + formatter.ToSnakeCase(fe.Param()) + "!"
	case "gt":
		return "should be greater than " + formatter.ToSnakeCase(fe.Param()) + "!"
	case "lte":
		return "should be less than or equal to " + formatter.ToSnakeCase(fe.Param()) + "!"
	case "email":
		return "must be a valid email address!"
	case "eqfield":
		return "does not match with " + formatter.ToSnakeCase(fe.Param()) + "!"
	case "ltfield":
		return "must be less than " + formatter.ToSnakeCase(fe.Param()) + " field!"
	case "gtfield":
		return "must be greater than " + formatter.ToSnakeCase(fe.Param()) + " field!"
	case "alpha":
		return "must be entirely alphabetic characters!"
	case "alphanum":
		return "must be entirely alpha-numeric characters!"
	case "numeric":
		return "must be an integer!"
	case "oneof":
		return "must be one of " + strings.Replace(fe.Param(), " ", ", ", -1)
	case "len":
		return "must have a length of " + formatter.ToSnakeCase(fe.Param()) + "!"
	case "filesize":
		return "file size must be less than " + formatter.ToSnakeCase(fe.Param()) + "MB !"
	case "filetype":
		return "file type does not match with " + formatter.ToSnakeCase(fe.Param()) + " type!"
	}
	return "something is wrong with this field!"
}

func FormatValidation(ve validator.ValidationErrors) []ErrorValidator {
	errList := []ErrorValidator{}

	for _, fe := range ve {
		errList = append(errList, ErrorValidator{
			Key:     formatter.ToSnakeCase(fe.Field()),
			Message: getErrorMsg(fe),
		})
	}

	return errList
}
