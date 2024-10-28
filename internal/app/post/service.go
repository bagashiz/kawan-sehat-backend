package post

// Service provides post business logic and operations.
type Service struct {
	repo ReadWriter
}

// NewService creates a new post service instance.
func NewService(repo ReadWriter) *Service {
	return &Service{repo: repo}
}
