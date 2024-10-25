package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// validate is a singleton instance of the validator v10 for validation.
var validate = validator.New(validator.WithRequiredStructEnabled())

func init() {
	customValidations := map[string]validator.Func{
		"username": usernameValidation,
		"role":     roleValidation,
		"avatar":   avatarValidation,
		"gender":   genderValidation,
		// add more custom validations here as needed
	}

	for tag, fn := range customValidations {
		validate.RegisterValidation(tag, fn)
	}
}

// ValidateParams validates and compiles each validation errors into a single error message.
// each validation  errors separated by new line.
func ValidateParams(params any) error {
	if err := validate.Struct(params); err != nil {
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
