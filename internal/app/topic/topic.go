package topic

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Topic represents a topic of interest
// for users to follow and post their
// stories on.
type Topic struct {
	ID          uuid.UUID
	Name        string
	Slug        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// New creates a new Topic instance.
func New(name, description string) (*Topic, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New("failed to generate topic id")
	}

	return &Topic{
		ID:          id,
		Name:        name,
		Slug:        generateSlug(name),
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

// generateSlug generates a slug from the topic name.
func generateSlug(name string) string {
	slug := strings.ToLower(name)
	slug = strings.ReplaceAll(slug, " ", "-")
	// Remove non-alphanumeric characters except hyphens
	re := regexp.MustCompile("[^a-z0-9-]+")
	slug = re.ReplaceAllString(slug, "")
	return slug
}

// Update modifies the topic's name and description.
func (t *Topic) Update(name, description string) {
	if name != "" {
		t.Name = name
		t.Slug = generateSlug(name)
	}
	if description != "" {
		t.Description = description
	}
	t.UpdatedAt = time.Now()
}
