package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Role represents the role of the user's account.
type Role string

const (
	Patient Role = "PATIENT"
	Expert  Role = "EXPERT"
	Admin   Role = "ADMIN"
)

// Avatar represents the avatar of the user's account.
type Avatar string

const (
	None        Avatar = "NONE"
	OldFemale   Avatar = "OLD_FEMALE"
	OldMale     Avatar = "OLD_MALE"
	YoungFemale Avatar = "YOUNG_FEMALE"
	YoungMale   Avatar = "YOUNG_MALE"
)

// Gender represents the gender of the user's account.
type Gender string

const (
	Female      Gender = "FEMALE"
	Male        Gender = "MALE"
	Unspecified Gender = "UNSPECIFIED"
)

// Account represents the user's account.
type Account struct {
	ID        uuid.UUID
	FullName  string
	Username  string
	NIK       string
	Email     string
	Password  string
	Gender    Gender
	Role      Role
	Avatar    Avatar
	CreatedAt time.Time
	UpdatedAt time.Time
}

// IllnessHistory represents the user's illness history.
type IllnessHistory struct {
	AccountID uuid.UUID
	Illness   string
	Date      time.Time
}

// New creates a new Account instance.
func New(userName, email, password string) (*Account, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, errors.New("failed to generate account id")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash account password")
	}

	account := &Account{
		ID:        id,
		Username:  userName,
		Email:     email,
		Password:  string(hashedPassword),
		Gender:    Unspecified,
		Role:      Patient,
		Avatar:    None,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return account, nil
}

// ComparePassword compare the account hashed password with the given password.
func (a *Account) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}

// Update updates the account with the given parameters.
func (a *Account) Update(
	fullName, userName, nik, email, password, gender, role, avatar string,
) error {
	if fullName != "" {
		a.FullName = fullName
	}
	if userName != "" {
		a.Username = userName
	}
	if nik != "" && len(nik) == 16 {
		a.NIK = nik
	}
	if email != "" {
		a.Email = email
	}
	if gender != "" {
		a.Gender = Gender(gender)
	}
	if role != "" {
		if a.Role != Admin {
			return ErrAccountForbidden
		}
		a.Role = Role(role)
	}
	if avatar != "" {
		a.Avatar = Avatar(avatar)
	}
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("failed to hash account password")
		}
		a.Password = string(hashedPassword)
	}
	a.UpdatedAt = time.Now()
	return nil
}
