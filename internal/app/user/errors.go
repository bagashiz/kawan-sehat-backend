package user

// UserError is a custom error for user domain logic.
type UserError struct {
	message string
}

// Error returns the error message for the UserError type.
func (ue *UserError) Error() string {
	return ue.message
}

var (
	// list of errors for user repository
	ErrAccountInvalid           = &UserError{"account data is invalid"}
	ErrAccountDuplicateEmail    = &UserError{"account with this email already exists"}
	ErrAccountDuplicateUsername = &UserError{"account with this username already exists"}
	ErrAccountNotFound          = &UserError{"account not found"}
)
