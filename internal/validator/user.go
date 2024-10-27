package validator

import (
	"regexp"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
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
	case string(user.Patient), string(user.Expert), string(user.Admin):
		return true
	default:
		return false
	}
}

// avatarValidation validates user avatar accounts.
func avatarValidation(fl validator.FieldLevel) bool {
	avatar := fl.Field().String()
	switch avatar {
	case string(user.None), string(user.OldFemale), string(user.OldMale),
		string(user.YoungFemale), string(user.YoungMale):
		return true
	default:
		return false
	}
}

// genderValidation validates user gender accounts.
func genderValidation(fl validator.FieldLevel) bool {
	gender := fl.Field().String()
	switch gender {
	case string(user.Female), string(user.Male), string(user.Unspecified):
		return true
	default:
		return false
	}
}
