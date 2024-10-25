package user

// Service provides user business logic and operations.
type Service struct {
	repo      ReadWriter
	tokenizer Tokenizer
}

// NewService creates a new user service instance.
func NewService(repo ReadWriter, tokenizer Tokenizer) *Service {
	return &Service{repo: repo, tokenizer: tokenizer}
}
