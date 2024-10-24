package user

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// usernameValidation validates username accounts.
func usernameValidation(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	// letters, numbers, periods, underscores, and hyphens
	return regexp.MustCompile(`^[a-zA-Z0-9._-]+$`).MatchString(username)
}

// roleValidation validates user role accounts.
func roleValidation(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	switch role {
	case "PATIENT", "EXPERT", "ADMIN":
		return true
	default:
		return false
	}
}

// avatarValidation validates user avatar accounts.
func avatarValidation(fl validator.FieldLevel) bool {
	avatar := fl.Field().Interface().(Avatar)
	switch avatar {
	case "NONE", "OLD_FEMALE", "OLD_MALE", "YOUNG_FEMALE", "YOUNG_MALE":
		return true
	default:
		return false
	}
}

// genderValidation validates user role accounts.
func genderValidation(fl validator.FieldLevel) bool {
	gender := fl.Field().Interface().(Gender)
	switch gender {
	case "PATIENT", "EXPERT", "ADMIN":
		return true
	default:
		return false
	}
}
