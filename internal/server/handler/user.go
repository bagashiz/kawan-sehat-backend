package handler

import (
	"net/http"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// accountResponse holds the response data for the account object.
type accountResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Role      string    `json:"role"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// registerRequest holds the request data for the register handler.
type registerRequest struct {
	Username string `json:"username" validate:"required,username,lte=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// RegisterAccount is the handler for the account registration route.
func RegisterAccount(userSvc *user.Service) APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req registerRequest
		if err := decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}

		account, err := userSvc.RegisterAccount(
			r.Context(),
			user.RegisterAccountParams{
				Username: req.Username,
				Email:    req.Email,
				Password: req.Password,
			})
		if err != nil {
			handleUserError(err)
		}

		res := &accountResponse{
			ID:        account.ID,
			Username:  account.Username,
			Email:     account.Email,
			Gender:    string(account.Gender),
			Role:      string(account.Role),
			Avatar:    string(account.Avatar),
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		}

		return writeJSON(w, http.StatusCreated, res, nil)
	}
}

// loginRequest holds the request data for the register handler.
type loginRequest struct {
	Username string `json:"username" validate:"required,username,lte=30"`
	Password string `json:"password" validate:"required,min=8"`
}

// loginResponse holds the response data for the register handler.
type loginResponse struct {
	Account accountResponse `json:"account"`
	Token   string          `json:"token"`
}

// LoginAccount is the handler for the account login route.
func LoginAccount(userSvc *user.Service) APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req loginRequest
		if err := decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}

		account, token, err := userSvc.LoginAccount(
			r.Context(),
			user.LoginAccountParams{
				Username: req.Username,
				Password: req.Password,
			})
		if err != nil {
			handleUserError(err)
		}

		res := &loginResponse{
			Account: accountResponse{
				ID:        account.ID,
				Username:  account.Username,
				Email:     account.Email,
				Gender:    string(account.Gender),
				Role:      string(account.Role),
				Avatar:    string(account.Avatar),
				CreatedAt: account.CreatedAt,
				UpdatedAt: account.UpdatedAt,
			},
			Token: token,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// GetAccountByID is the handler for the account retrieval by ID route.
func GetAccountByID(userSvc *user.Service) APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")

		account, err := userSvc.GetAccountByID(r.Context(), id)
		if err != nil {
			handleUserError(err)
		}
		res := &accountResponse{
			ID:        account.ID,
			Username:  account.Username,
			Email:     account.Email,
			Gender:    string(account.Gender),
			Role:      string(account.Role),
			Avatar:    string(account.Avatar),
			CreatedAt: account.CreatedAt,
			UpdatedAt: account.UpdatedAt,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// handleUserError determines the appropriate HTTP status code for the user service error.
func handleUserError(err error) APIError {
	switch err {
	case user.ErrAccountNotFound:
		return NotFoundRequest(err)
	case user.ErrAccountUnauthorized:
		return UnauthorizedRequest(err)
	case user.ErrAccountForbidden:
		return ForbiddenRequest(err)
	case user.ErrAccountDuplicateEmail, user.ErrAccountDuplicateUsername:
		return ConflictRequest(err)
	default:
		return BadRequest(err)
	}
}
