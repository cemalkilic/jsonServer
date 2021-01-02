package validator

import (
    "github.com/go-playground/validator/v10"
    "net/http"
)

func NewValidator() *validator.Validate {
    v := validator.New()
    _ = v.RegisterValidation("httpStatus", httpStatusValidator)
    return v
}

func httpStatusValidator(fl validator.FieldLevel) bool {
    return http.StatusText(int(fl.Field().Int())) != ""
}
