package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator is a struct that holds the validator instance
// to validate struct fields.
type Validator struct {
	validate *validator.Validate
}

// New creates a new Validator instance and registers custom validations.
func New() *Validator {
	v := validator.New(validator.WithRequiredStructEnabled())

	customValidations := map[string]validator.Func{
		"username": usernameValidation,
		"role":     roleValidation,
		"avatar":   avatarValidation,
		"gender":   genderValidation,
		// add more custom validations here as needed
	}

	for tag, fn := range customValidations {
		v.RegisterValidation(tag, fn)
	}

	return &Validator{validate: v}
}

// ValidateParams validates and compiles each validation error into a single error message.
// Each validation error is separated by a new line.
func (v *Validator) ValidateParams(params any) error {
	if err := v.validate.Struct(params); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			var errorMsgs []string
			for _, e := range validationErrs {
				errorMessage := fmt.Sprintf("field:'%s' invalid:'%s' tag:'%s'", e.Field(), e.Value(), e.Tag())
				errorMsgs = append(errorMsgs, errorMessage)
			}
			return &ValidationError{strings.Join(errorMsgs, "\n")}
		}
		return err
	}
	return nil
}
