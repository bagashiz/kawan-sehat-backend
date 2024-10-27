package topic

// Service provides topic business logic and operations.
type Service struct {
	repo ReadWriter
}

// NewService creates a new topic service instance.
func NewService(repo ReadWriter) *Service {
	return &Service{repo: repo}
}
