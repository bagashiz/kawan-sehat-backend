package reply

// ReplyError is a custom error for reply domain logic.
type ReplyError struct {
	message string
}

// Error returns the error message for the ReplyError type.
func (te ReplyError) Error() string {
	return te.message
}

var (
	ErrReplyVoteInvalid      = ReplyError{"reply vote is invalid"}
	ErrReplyVoteNotFound     = ReplyError{"reply vote does not exists"}
	ErrReplyVoteAlreadyVoted = ReplyError{"reply vote already voted"}
	// list of errors for reply repository.
	ErrReplyInvalid         = ReplyError{"reply data is invalid"}
	ErrReplyActionForbidden = ReplyError{"reply action is forbidden"}
	ErrReplyNotFound        = ReplyError{"reply not found"}
)
