package validators

import (
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

func ValidateTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}
