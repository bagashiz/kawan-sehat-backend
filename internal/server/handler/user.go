package handler

import (
	"net/http"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
)

// accountResponse holds the response data for the account object.
type accountResponse struct {
	ID             string    `json:"id"`
	FullName       string    `json:"full_name"`
	Username       string    `json:"username"`
	NIK            string    `json:"nik"`
	Email          string    `json:"email"`
	Gender         string    `json:"gender"`
	Role           string    `json:"role"`
	Avatar         string    `json:"avatar"`
	IllnessHistory string    `json:"illness_history"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// registerRequest holds the request data for the register handler.
type registerRequest struct {
	Username string `json:"username" validate:"required,username,lte=30"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// RegisterAccount is the handler for the account registration route.
func (h *Handler) RegisterAccount() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req registerRequest
		if err := h.decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}

		account, err := h.userSvc.RegisterAccount(
			r.Context(),
			user.RegisterAccountParams{
				Username: req.Username,
				Email:    req.Email,
				Password: req.Password,
			})
		if err != nil {
			return handleUserError(err)
		}

		res := &accountResponse{
			ID:             account.ID.String(),
			FullName:       account.FullName,
			Username:       account.Username,
			NIK:            account.NIK,
			Email:          account.Email,
			Gender:         string(account.Gender),
			Role:           string(account.Role),
			Avatar:         string(account.Avatar),
			IllnessHistory: account.IllnessHistory,
			CreatedAt:      account.CreatedAt,
			UpdatedAt:      account.UpdatedAt,
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
func (h *Handler) LoginAccount() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req loginRequest
		if err := h.decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}

		account, token, err := h.userSvc.LoginAccount(
			r.Context(),
			user.LoginAccountParams{
				Username: req.Username,
				Password: req.Password,
			})
		if err != nil {
			return handleUserError(err)
		}

		res := &loginResponse{
			Account: accountResponse{
				ID:             account.ID.String(),
				FullName:       account.FullName,
				Username:       account.Username,
				NIK:            account.NIK,
				Email:          account.Email,
				Gender:         string(account.Gender),
				Role:           string(account.Role),
				Avatar:         string(account.Avatar),
				IllnessHistory: account.IllnessHistory,
				CreatedAt:      account.CreatedAt,
				UpdatedAt:      account.UpdatedAt,
			},
			Token: token,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// updateAccountRequest holds the request data for the update handler.
type updateAccountRequest struct {
	FullName       string `json:"full_name" validate:"omitempty,lte=50"`
	Username       string `json:"username" validate:"omitempty,username,lte=30"`
	NIK            string `json:"nik" validate:"omitempty,len=16"`
	Email          string `json:"email" validate:"omitempty,email"`
	Password       string `json:"password" validate:"omitempty,min=8"`
	Gender         string `json:"gender" validate:"omitempty,gender"`
	Role           string `json:"role" validate:"omitempty,role"`
	Avatar         string `json:"avatar" validate:"omitempty,avatar"`
	IllnessHistory string `json:"illness_history" validate:"omitempty,lte=255"`
}

// UpdateAccount is the handler for the account update route.
func (h *Handler) UpdateAccount() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req updateAccountRequest
		if err := h.decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}

		account, err := h.userSvc.UpdateAccount(
			r.Context(),
			user.UpdateAccountParams{
				FullName:       req.FullName,
				Username:       req.Username,
				NIK:            req.NIK,
				Email:          req.Email,
				Password:       req.Password,
				Gender:         req.Gender,
				Role:           req.Role,
				Avatar:         req.Avatar,
				IllnessHistory: req.IllnessHistory,
			})
		if err != nil {
			return handleUserError(err)
		}

		res := &accountResponse{
			ID:             account.ID.String(),
			FullName:       account.FullName,
			Username:       account.Username,
			NIK:            account.NIK,
			Email:          account.Email,
			Gender:         string(account.Gender),
			Role:           string(account.Role),
			Avatar:         string(account.Avatar),
			IllnessHistory: account.IllnessHistory,
			CreatedAt:      account.CreatedAt,
			UpdatedAt:      account.UpdatedAt,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// GetAccountByID is the handler for the account retrieval by ID route.
func (h *Handler) GetAccountByID() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")

		account, err := h.userSvc.GetAccountByID(r.Context(), id)
		if err != nil {
			return handleUserError(err)
		}

		res := &accountResponse{
			ID:        account.ID.String(),
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
	case user.ErrAccountUnauthorized:
		return UnauthorizedRequest(err)
	case user.ErrAccountForbidden:
		return ForbiddenRequest(err)
	case user.ErrAccountNotFound:
		return NotFoundRequest(err)
	case user.ErrAccountDuplicateEmail,
		user.ErrAccountDuplicateUsername,
		user.ErrAccountDuplicateNIK:
		return ConflictRequest(err)
	case user.ErrAccountInvalid:
		return UnprocessableRequest(err)
	default:
		return BadRequest(err)
	}
}
