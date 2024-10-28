package comment

// CommentError is a custom error for comment domain logic.
type CommentError struct {
	message string
}

// Error returns the error message for the CommentError type.
func (te CommentError) Error() string {
	return te.message
}

var (
	// list of errors for comment repository.
	ErrCommentInvalid         = CommentError{"comment data is invalid"}
	ErrCommentActionForbidden = CommentError{"comment action is forbidden"}
	ErrCommentNotFound        = CommentError{"comment not found"}
)
