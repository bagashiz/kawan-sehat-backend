package postgres

import (
	"github.com/bagashiz/kawan-sehat-backend/internal/app/topic"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
)

// ToDomain converts generated Account struct to domain account struct
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

// ToDomain converts generated Topic struct to domain account struct
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
