package handlers

import (
	"financial/internal/errors"

	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func validateRequest(data any) *errors.ServiceError {
	validate := newValidator()
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	serviceErr := errors.BadRequestError()

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		addFieldErrors(serviceErr, validationErrors)
	} else {
		return errors.InternalServerError().WithDetails(err.Error())
	}

	return serviceErr
}

func addFieldErrors(serviceErr *errors.ServiceError, validationErrors validator.ValidationErrors) {
	for _, fieldErr := range validationErrors {
		fieldName := fieldErr.Field()
		tag := fieldErr.Tag()
		param := fieldErr.Param()

		switch tag {
		case "required":
			serviceErr.AddFieldError(
				fieldName,
				errors.RequiredField(fieldName).Message,
			)
		case "min":
			serviceErr.AddFieldError(
				fieldName,
				fmt.Sprintf("%s must be at least %s", fieldName, param),
			)
		case "max":
			serviceErr.AddFieldError(
				fieldName,
				fmt.Sprintf("%s must be at most %s", fieldName, param),
			)
		case "email":
			serviceErr.AddFieldError(
				fieldName,
				errors.InvalidFormat(fieldName).Message,
			)
		default:
			serviceErr.AddFieldError(
				fieldName,
				fmt.Sprintf("%s is invalid", fieldName),
			)
		}
	}
}

func newValidator() *validator.Validate {
	validate := validator.New()
	// register validator to use field as it is in json
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("json")
		// Handle cases where json tag might have options like ",omitempty"
		if commaIdx := strings.Index(name, ","); commaIdx != -1 {
			name = name[:commaIdx]
		}
		if name == "" {
			name = fld.Name
		}
		return name
	})

	return validate
}
