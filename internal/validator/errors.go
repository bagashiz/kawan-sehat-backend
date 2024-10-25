package validator

// ValidationError is a custom error for user input validation.
type ValidationError struct {
	message string
}

// Error returns the error message for the UserError type.
func (ve *ValidationError) Error() string {
	return ve.message
}
