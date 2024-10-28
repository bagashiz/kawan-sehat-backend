package handler

import (
	"net/http"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/comment"
)

// commentResponse holds the response data for the comment object.
type commentResponse struct {
	ID           string                  `json:"id"`
	PostID       string                  `json:"post_id"`
	Account      *commentAccountResponse `json:"account,omitempty"`
	Vote         *commentVoteResponse    `json:"vote,omitempty"`
	TotalReplies int64                   `json:"total_replies"`
	Content      string                  `json:"content"`
	CreatedAt    time.Time               `json:"created_at"`
}

// commentAccountResponse holds the response data for the comment account object.
type commentAccountResponse struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

// commentVoteResponse holds the response data for the comment vote object.
type commentVoteResponse struct {
	Total int64 `json:"total"`
	State int32 `json:"state"`
}

// createCommentRequest holds the request data for the create comment handler.
type createCommentRequest struct {
	PostID  string `json:"post_id" validate:"required,uuid"`
	Content string `json:"content" validate:"required,lte=1000"`
}

// CreateComment is the handler for the comment creation route.
func (h *Handler) CreateComment() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req createCommentRequest
		if err := h.decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}

		comment, err := h.commentSvc.AddComment(
			r.Context(),
			comment.CreateCommentParams{
				PostID:  req.PostID,
				Content: req.Content,
			})
		if err != nil {
			return handleCommentError(err)
		}

		res := &commentResponse{
			ID:     comment.ID.String(),
			PostID: comment.PostID.String(),
			Account: &commentAccountResponse{
				ID: comment.Account.ID.String(),
			},
			Vote: &commentVoteResponse{
				Total: comment.Vote.Total,
				State: comment.Vote.State,
			},
			TotalReplies: comment.TotalReplies,
			Content:      comment.Content,
			CreatedAt:    comment.CreatedAt,
		}
		return writeJSON(w, http.StatusCreated, res, nil)
	}
}

// DeleteComment is the handler for the comment deletion route.
func (h *Handler) DeleteComment() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		if err := h.commentSvc.DeleteComment(r.Context(), id); err != nil {
			return handleCommentError(err)
		}
		return writeJSON(w, http.StatusOK, nil, nil)
	}
}

// UpvoteComment is the handler for the comment voting route.
func (h *Handler) UpvoteComment() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		count, err := h.commentSvc.UpvoteComment(r.Context(), id)
		if err != nil {
			return handlePostError(err)
		}
		return writeJSON(w, http.StatusOK, count, nil)
	}
}

// DownvoteComment is the handler for the comment voting route.
func (h *Handler) DownvoteComment() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		count, err := h.commentSvc.DownvoteComment(r.Context(), id)
		if err != nil {
			return handlePostError(err)
		}
		return writeJSON(w, http.StatusOK, count, nil)
	}
}

// listCommentsByPostIDResponse holds the response data for the list comments by post ID handler.
type listCommentsByPostIDResponse struct {
	Limit    int32              `json:"limit" validate:"omitempty,gte=1,lte=100"`
	Page     int32              `json:"page" validate:"omitempty,gte=1"`
	Count    int64              `json:"count"`
	Comments []*commentResponse `json:"comments"`
}

// ListCommentsByPostID is the handler for the comment of a post listing route.
func (h *Handler) ListCommentsByPostID() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		limit, page := getLimitPage(r)
		postID := r.PathValue("id")
		comments, count, err := h.commentSvc.ListCommentsByPostID(r.Context(), postID, limit, page)
		if err != nil {
			return handleCommentError(err)
		}
		commentsRes := make([]*commentResponse, len(comments))
		for i, c := range comments {
			commentsRes[i] = &commentResponse{
				ID:     c.ID.String(),
				PostID: c.PostID.String(),
				Account: &commentAccountResponse{
					ID:       c.Account.ID.String(),
					Username: c.Account.Username,
				},
				Vote: &commentVoteResponse{
					Total: c.Vote.Total,
					State: c.Vote.State,
				},
				TotalReplies: c.TotalReplies,
				Content:      c.Content,
				CreatedAt:    c.CreatedAt,
			}
		}
		res := &listCommentsByPostIDResponse{
			Limit:    limit,
			Page:     page,
			Count:    count,
			Comments: commentsRes,
		}
		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// handleCommentError converts the given error into an appropriate response.
func handleCommentError(err error) APIError {
	switch err {
	case comment.ErrCommentActionForbidden:
		return ForbiddenRequest(err)
	case comment.ErrCommentNotFound:
		return NotFoundRequest(err)
	case comment.ErrCommentInvalid:
		return UnprocessableRequest(err)
	default:
		return BadRequest(err)
	}
}
