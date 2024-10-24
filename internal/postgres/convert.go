package postgres

import "github.com/bagashiz/kawan-sehat-backend/internal/app/user"

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
