package topic

// TopicError is a custom error for topic domain logic.
type TopicError struct {
	message string
}

// Error returns the error message for the TopicError type.
func (te TopicError) Error() string {
	return te.message
}

var (
	// list of errors for user repository.
	ErrTopicInvalid                  = TopicError{"topic data is invalid"}
	ErrTopicDuplicateName            = TopicError{"topic with this name already exists"}
	ErrTopicNotFound                 = TopicError{"topic not found"}
	ErrAccountTopicsInvalid          = TopicError{"account topics data is invalid"}
	ErrAccountTopicsAlreadyExists    = TopicError{"topic already followed by this account"}
	ErrAccountAlreadyUnfollowedTopic = TopicError{"topic already unfollowed or does not exist"}
)
