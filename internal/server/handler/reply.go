package handler

import (
	"net/http"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/reply"
)

// replyResponse holds the response data for the reply object.
type replyResponse struct {
	ID        string                `json:"id"`
	CommentID string                `json:"comment_id"`
	Account   *replyAccountResponse `json:"account,omitempty"`
	Vote      *replyVoteResponse    `json:"vote,omitempty"`
	Content   string                `json:"content"`
	CreatedAt time.Time             `json:"created_at"`
}

// replyAccountResponse holds the response data for the reply account object.
type replyAccountResponse struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Role     string `json:"role,omitempty"`
}

// replyVoteResponse holds the response data for the reply vote object.
type replyVoteResponse struct {
	Total int64 `json:"total"`
	State int32 `json:"state"`
}

// createReplyRequest holds the request data for the create reply handler.
type createReplyRequest struct {
	CommentID string `json:"comment_id" validate:"required,uuid"`
	Content   string `json:"content" validate:"required,lte=1000"`
}

// CreateReply is the handler for the reply creation route.
func (h *Handler) CreateReply() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req createReplyRequest
		if err := h.decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}

		reply, err := h.replySvc.AddReply(
			r.Context(),
			reply.CreateReplyParams{
				CommentID: req.CommentID,
				Content:   req.Content,
			})
		if err != nil {
			return handleReplyError(err)
		}

		res := &replyResponse{
			ID:        reply.ID.String(),
			CommentID: reply.CommentID.String(),
			Account: &replyAccountResponse{
				ID:       reply.Account.ID.String(),
				Username: reply.Account.Username,
				Avatar:   string(reply.Account.Avatar),
				Role:     string(reply.Account.Role),
			},
			Vote: &replyVoteResponse{
				Total: reply.Vote.Total,
				State: reply.Vote.State,
			},
			Content:   reply.Content,
			CreatedAt: reply.CreatedAt,
		}
		return writeJSON(w, http.StatusCreated, res, nil)
	}
}

// DeleteReply is the handler for the reply deletion route.
func (h *Handler) DeleteReply() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		if err := h.replySvc.DeleteReply(r.Context(), id); err != nil {
			return handleReplyError(err)
		}
		return writeJSON(w, http.StatusOK, nil, nil)
	}
}

// UpvoteReply is the handler for the reply voting route.
func (h *Handler) UpvoteReply() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		count, err := h.replySvc.UpvoteReply(r.Context(), id)
		if err != nil {
			return handlePostError(err)
		}
		return writeJSON(w, http.StatusOK, count, nil)
	}
}

// DownvoteReply is the handler for the reply voting route.
func (h *Handler) DownvoteReply() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		count, err := h.replySvc.DownvoteReply(r.Context(), id)
		if err != nil {
			return handlePostError(err)
		}
		return writeJSON(w, http.StatusOK, count, nil)
	}
}

// listRepliesByCommentIDResponse holds the response data for the list replys by post ID handler.
type listRepliesByCommentIDResponse struct {
	Limit   int32            `json:"limit" validate:"omitempty,gte=1,lte=100"`
	Page    int32            `json:"page" validate:"omitempty,gte=1"`
	Count   int64            `json:"count"`
	Replies []*replyResponse `json:"replys"`
}

// ListRepliesByCommentID is the handler for the reply of a post listing route.
func (h *Handler) ListRepliesByCommentID() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		limit, page := getLimitPage(r)
		postID := r.PathValue("id")
		replys, count, err := h.replySvc.ListRepliesByCommentID(r.Context(), postID, limit, page)
		if err != nil {
			return handleReplyError(err)
		}
		replysRes := make([]*replyResponse, len(replys))
		for i, r := range replys {
			replysRes[i] = &replyResponse{
				ID:        r.ID.String(),
				CommentID: r.CommentID.String(),
				Account: &replyAccountResponse{
					ID:       r.Account.ID.String(),
					Username: r.Account.Username,
					Avatar:   string(r.Account.Avatar),
					Role:     string(r.Account.Role),
				},
				Vote: &replyVoteResponse{
					Total: r.Vote.Total,
					State: r.Vote.State,
				},
				Content:   r.Content,
				CreatedAt: r.CreatedAt,
			}
		}
		res := &listRepliesByCommentIDResponse{
			Limit:   limit,
			Page:    page,
			Count:   count,
			Replies: replysRes,
		}
		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// handleReplyError converts the given error into an appropriate response.
func handleReplyError(err error) APIError {
	switch err {
	case reply.ErrReplyActionForbidden:
		return ForbiddenRequest(err)
	case reply.ErrReplyNotFound:
		return NotFoundRequest(err)
	case reply.ErrReplyInvalid:
		return UnprocessableRequest(err)
	default:
		return BadRequest(err)
	}
}
