package core

import (
	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("latitude", ValidateLatitude)
	validate.RegisterValidation("longitude", ValidateLongitude)
	return validate
}

func ValidateLatitude(fl validator.FieldLevel) bool {
	lat := fl.Field().Float()
	return lat >= -90 && lat <= 90
}
func ValidateLongitude(fl validator.FieldLevel) bool {
	lon := fl.Field().Float()
	return lon >= -180 && lon <= 180
}
