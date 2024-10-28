package post

// PostError is a custom error for post domain logic.
type PostError struct {
	message string
}

// Error returns the error message for the PostError type.
func (te PostError) Error() string {
	return te.message
}

var (
	ErrPostVoteInvalid      = PostError{"post vote data is invalid"}
	ErrPostVoteNotFound     = PostError{"post vote does not exists"}
	ErrPostVoteAlreadyVoted = PostError{"post vote already exists"}
	// list of errors for post repository.
	ErrPostInvalid             = PostError{"post data is invalid"}
	ErrPostTopicNotFound       = PostError{"post topic not found"}
	ErrPostAccountNotFound     = PostError{"post account not found"}
	ErrPostActionForbidden     = PostError{"post action is forbidden"}
	ErrPostNotFound            = PostError{"post not found"}
	ErrBookmarkInvalid         = PostError{"bookmark data is invalid"}
	ErrBookmarkExists          = PostError{"bookmark already exists"}
	ErrBookmarkAlreadyUnmarked = PostError{"bookmark is already unmarked"}
)
