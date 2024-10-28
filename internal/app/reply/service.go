package reply

// Service provides comment business logic and operations.
type Service struct {
	repo ReadWriter
}

// NewService creates a new comment service instance.
func NewService(repo ReadWriter) *Service {
	return &Service{repo: repo}
}
