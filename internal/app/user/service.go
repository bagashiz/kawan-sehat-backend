package user

import "github.com/go-playground/validator/v10"

// Service provides user business logic and operations.
type Service struct {
	repo      ReadWriter
	tokenizer Tokenizer
	validate  *validator.Validate
}

// NewService creates a new user service instance.
func NewService(repo ReadWriter, tokenizer Tokenizer, validate *validator.Validate) *Service {
	return &Service{repo: repo, tokenizer: tokenizer}
}
