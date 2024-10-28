package postgres

import (
	"github.com/bagashiz/kawan-sehat-backend/internal/app/comment"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/topic"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
)

// ToDomain converts generated Account struct to domain account struct.
func (a Account) ToDomain() *user.Account {
	return &user.Account{
		ID:             a.ID,
		FullName:       a.FullName.String,
		Username:       a.Username,
		NIK:            a.Nik.String,
		Email:          a.Email,
		Password:       a.Password,
		Gender:         user.Gender(a.Gender),
		Role:           user.Role(a.Role),
		Avatar:         user.Avatar(a.Avatar),
		IllnessHistory: a.IllnessHistory.String,
		CreatedAt:      a.CreatedAt,
		UpdatedAt:      a.UpdatedAt,
	}
}

// ToDomain converts generated Topic struct to domain topic struct.
func (t Topic) ToDomain() *topic.Topic {
	return &topic.Topic{
		ID:          t.ID,
		Name:        t.Name,
		Slug:        t.Slug,
		Description: t.Description,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

// ToDomain converts generated Post struct to domain post struct.
func (p SelectPostByIDRow) ToDomain() *post.Post {
	return &post.Post{
		ID: p.ID,
		Account: &post.Account{
			ID:       p.AccountID,
			Username: p.AccountUsername,
		},
		Topic: &post.Topic{
			ID:   p.TopicID,
			Name: p.TopicName,
		},
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToDomain converts generated Post struct to domain post struct.
func (p SelectAllPostsRow) ToDomain() *post.Post {
	return &post.Post{
		ID: p.ID,
		Account: &post.Account{
			ID:       p.AccountID,
			Username: p.AccountUsername,
		},
		Topic: &post.Topic{
			ID:   p.TopicID,
			Name: p.TopicName,
		},
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToDomain converts generated Post struct to domain post struct.
func (p SelectAllPostsPaginatedRow) ToDomain() *post.Post {
	return &post.Post{
		ID: p.ID,
		Account: &post.Account{
			ID:       p.AccountID,
			Username: p.AccountUsername,
		},
		Topic: &post.Topic{
			ID:   p.TopicID,
			Name: p.TopicName,
		},
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToDomain converts generated Post struct to domain post struct.
func (p SelectPostsByAccountIDRow) ToDomain() *post.Post {
	return &post.Post{
		ID: p.ID,
		Account: &post.Account{
			ID:       p.AccountID,
			Username: p.AccountUsername,
		},
		Topic: &post.Topic{
			ID:   p.TopicID,
			Name: p.TopicName,
		},
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToDomain converts generated Post struct to domain post struct.
func (p SelectPostsByAccountIDPaginatedRow) ToDomain() *post.Post {
	return &post.Post{
		ID: p.ID,
		Account: &post.Account{
			ID:       p.AccountID,
			Username: p.AccountUsername,
		},
		Topic: &post.Topic{
			ID:   p.TopicID,
			Name: p.TopicName,
		},
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToDomain converts generated Post struct to domain post struct.
func (p SelectPostsByTopicIDRow) ToDomain() *post.Post {
	return &post.Post{
		ID: p.ID,
		Account: &post.Account{
			ID:       p.AccountID,
			Username: p.AccountUsername,
		},
		Topic: &post.Topic{
			ID:   p.TopicID,
			Name: p.TopicName,
		},
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToDomain converts generated Post struct to domain post struct.
func (p SelectPostsByTopicIDPaginatedRow) ToDomain() *post.Post {
	return &post.Post{
		ID: p.ID,
		Account: &post.Account{
			ID:       p.AccountID,
			Username: p.AccountUsername,
		},
		Topic: &post.Topic{
			ID:   p.TopicID,
			Name: p.TopicName,
		},
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToDomain converts generated Post struct to domain post struct.
func (p SelectBookmarksByAccountIDRow) ToDomain() *post.Post {
	return &post.Post{
		ID: p.ID,
		Account: &post.Account{
			ID:       p.AccountID,
			Username: p.AccountUsername,
		},
		Topic: &post.Topic{
			ID:   p.TopicID,
			Name: p.TopicName,
		},
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToDomain converts generated Post struct to domain post struct.
func (p SelectBookmarksByAccountIDPaginatedRow) ToDomain() *post.Post {
	return &post.Post{
		ID: p.ID,
		Account: &post.Account{
			ID:       p.AccountID,
			Username: p.AccountUsername,
		},
		Topic: &post.Topic{
			ID:   p.TopicID,
			Name: p.TopicName,
		},
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ToDomain converts postgres repository results to domain entities.
func (c Comment) ToDomain() *comment.Comment {
	return &comment.Comment{
		ID:     c.ID,
		PostID: c.PostID,
		Account: &comment.Account{
			ID: c.AccountID,
		},
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
	}
}

// ToDomain converts postgres repository results to domain entities.
func (c SelectCommentsByPostIDRow) ToDomain() *comment.Comment {
	return &comment.Comment{
		ID:     c.ID,
		PostID: c.PostID,
		Account: &comment.Account{
			ID:       c.AccountID,
			Username: c.AccountUsername,
		},
		Content: c.Content,
		Vote: &comment.Vote{
			Total: c.TotalVotes.(int64),
			State: c.VoteState.(int32),
		},
		TotalReplies: c.TotalReplies,
		CreatedAt:    c.CreatedAt,
	}
}

// ToDomain converts postgres repository results to domain entities.
func (c SelectCommentsByPostIDPaginatedRow) ToDomain() *comment.Comment {
	return &comment.Comment{
		ID:     c.ID,
		PostID: c.PostID,
		Account: &comment.Account{
			ID:       c.AccountID,
			Username: c.AccountUsername,
		},
		Content: c.Content,
		Vote: &comment.Vote{
			Total: c.TotalVotes.(int64),
			State: c.VoteState.(int32),
		},
		TotalReplies: c.TotalReplies,
		CreatedAt:    c.CreatedAt,
	}
}
