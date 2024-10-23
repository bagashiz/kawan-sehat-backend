package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Role represents the role of the user's account.
type Role string

const (
	Patient Role = "patient"
	Expert  Role = "expert"
	Admin   Role = "admin"
)

// Avatar represents the avatar of the user's account.
type Avatar string

const (
	OldFemale   Avatar = "old_female"
	OldMale     Avatar = "old_male"
	YoungFemale Avatar = "young_female"
	YoungMale   Avatar = "young_male"
)

// Gender represents the gender of the user's account.
type Gender string

const (
	Female      Gender = "female"
	Male        Gender = "male"
	Unspecified Gender = "unspecified"
)

// Account represents the user's account.
type Account struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Name           string
	NIK            string
	Email          string
	Password       string
	Gender         Gender
	Role           Role
	Avatar         Avatar
	IllnessHistory string
	ID             uuid.UUID
}

// New creates a new Account instance.
func New(
	name, nik, email, password, gender, role, avatar, illnessHistory string,
) (*Account, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New("failed to generate account id")
	}

	account := &Account{
		ID:             id,
		Name:           name,
		NIK:            nik,
		Email:          email,
		Password:       password,
		Gender:         Gender(gender),
		Role:           Role(role),
		Avatar:         Avatar(avatar),
		IllnessHistory: illnessHistory,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return account, nil
}
