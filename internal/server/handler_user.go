package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/google/uuid"
)

// registerRequest holds the request data for the register handler.
type registerRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// registerResponse holds the response data for the register handler.
type registerResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// registerAccount is the handler for the account registration route.
func registerAccount(userSvc *user.Service) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req registerRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			return &handlerError{"invalid request body", http.StatusBadRequest}
		}

		account, err := userSvc.RegisterAccount(
			r.Context(),
			user.RegisterAccountParams{
				Username: req.Username,
				Email:    req.Email,
				Password: req.Password,
			})
		if err != nil {
			if userErr, ok := err.(*user.UserError); ok {
				return &handlerError{userErr.Error(), http.StatusBadRequest}
			}
			return err
		}

		res := registerResponse{
			ID:        account.ID,
			Username:  account.Username,
			Email:     account.Email,
			CreatedAt: account.CreatedAt,
		}

		return sendJSONResponse(w, http.StatusCreated, res, nil)
	}
}
