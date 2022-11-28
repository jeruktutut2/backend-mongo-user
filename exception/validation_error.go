package exception

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func GetValidationErrors(err error, request interface{}) error {
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		badRequestMessage := make(map[string][]string)
		for _, fieldError := range validationErrors {
			field, ok := reflect.TypeOf(request).FieldByName(fieldError.Field())
			if !ok {
				return BadRequestError{Message: "Cannot find tag " + fieldError.Field()}
			}
			fieldTagJson := field.Tag.Get("json")
			message := ""
			if fieldError.Tag() == "required" {
				message = "is required"
			} else if fieldError.Tag() == "min" {
				message = "must be more than or equal to " + fieldError.Param() + " letters"
			} else if fieldError.Tag() == "containsany" {
				message = "must contain any letter " + fieldError.Param()
			}
			badRequestMessage[fieldTagJson] = append(badRequestMessage[fieldTagJson], message)
		}
		return BadRequestError{Message: badRequestMessage}
	}
	return nil
}
