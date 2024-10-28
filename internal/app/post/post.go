package post

import (
	"time"

	"github.com/google/uuid"
)

// Post represents a user's story post on a topic.
type Post struct {
	ID        uuid.UUID
	Account   *Account
	Topic     *Topic
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Account represents a user account the post belongs to.
type Account struct {
	ID       uuid.UUID
	Username string
}

// Topic represents a topic the post belongs to.
type Topic struct {
	ID   uuid.UUID
	Name string
}

// New creates a new Post instance.
func New(accountID, topicID, title, content string) (*Post, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	accountUUID, err := uuid.Parse(accountID)
	if err != nil {
		return nil, err
	}

	topicUUID, err := uuid.Parse(topicID)
	if err != nil {
		return nil, err
	}

	return &Post{
		ID: id,
		Account: &Account{
			ID: accountUUID,
		},
		Topic: &Topic{
			ID: topicUUID,
		},
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// Update modifies title and content of the post.
func (p *Post) Update(title, content string) {
	if title != "" {
		p.Title = title
	}
	if content != "" {
		p.Content = content
	}
	p.UpdatedAt = time.Now()
}
