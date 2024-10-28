package handler

import (
	"net/http"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
)

// postResponse holds the response data for the post object.
type postResponse struct {
	ID        string    `json:"id"`
	AccountID string    `json:"account_id"`
	TopicID   string    `json:"topic_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// createPostRequest holds the request data for the create post handler.
type createPostRequest struct {
	TopicID string `json:"topic_id" validate:"required,uuid"`
	Title   string `json:"title" validate:"required,lte=100"`
	Content string `json:"content" validate:"required,lte=1000"`
}

// CreatePost is the handler for the post creation route.
func (h *Handler) CreatePost() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req createPostRequest
		if err := h.decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}

		post, err := h.postSvc.AddPost(
			r.Context(),
			post.CreatePostParams{
				TopicID: req.TopicID,
				Title:   req.Title,
				Content: req.Content,
			})
		if err != nil {
			return handlePostError(err)
		}

		res := &postResponse{
			ID:        post.ID.String(),
			AccountID: post.AccountID.String(),
			TopicID:   post.TopicID.String(),
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}
		return writeJSON(w, http.StatusCreated, res, nil)
	}
}

// updatePostRequest holds the request data for the update post handler.
type updatePostRequest struct {
	Title   string `json:"title" validate:"lte=100"`
	Content string `json:"content" validate:"lte=1000"`
}

// UpdatePost is the handler for the post update route.
func (h *Handler) UpdatePost() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req updatePostRequest
		if err := h.decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}
		id := r.PathValue("id")

		post, err := h.postSvc.UpdatePost(
			r.Context(),
			post.UpdatePostParams{
				ID:      id,
				Title:   req.Title,
				Content: req.Content,
			})
		if err != nil {
			return handlePostError(err)
		}

		res := &postResponse{
			ID:        post.ID.String(),
			AccountID: post.AccountID.String(),
			TopicID:   post.TopicID.String(),
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// DeletePost is the handler for the post deletion route.
func (h *Handler) DeletePost() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")

		err := h.postSvc.DeletePost(r.Context(), id)
		if err != nil {
			return handlePostError(err)
		}

		return writeJSON(w, http.StatusOK, nil, nil)
	}
}

// GetPostByID is the handler for the post retrieval route.
func (h *Handler) GetPostByID() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")

		post, err := h.postSvc.GetPost(r.Context(), id)
		if err != nil {
			return handleTopicError(err)
		}

		res := &postResponse{
			ID:        post.ID.String(),
			AccountID: post.AccountID.String(),
			TopicID:   post.TopicID.String(),
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// listPostsResponse holds the response data for the list posts handler.
type listPostsResponse struct {
	Limit  int32          `json:"limit"`
	Offset int32          `json:"offset"`
	Count  int            `json:"count"`
	Posts  []postResponse `json:"posts"`
}

// ListPosts is the handler for the post list route.
func (h *Handler) ListPosts() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		limit, offset := getLimitOffset(r)
		accountID := r.URL.Query().Get("account_id")
		topicID := r.URL.Query().Get("topic_id")

		posts, err := h.postSvc.ListPosts(
			r.Context(),
			post.ListPostsParams{
				AccountID: accountID,
				TopicID:   topicID,
				Limit:     limit,
				Offset:    offset,
			})
		if err != nil {
			return handlePostError(err)
		}

		postsRes := make([]postResponse, len(posts))
		for i, p := range posts {
			postsRes[i] = postResponse{
				ID:        p.ID.String(),
				AccountID: p.AccountID.String(),
				TopicID:   p.TopicID.String(),
				Title:     p.Title,
				Content:   p.Content,
				CreatedAt: p.CreatedAt,
				UpdatedAt: p.UpdatedAt,
			}
		}

		res := listPostsResponse{
			Limit:  limit,
			Offset: offset,
			Count:  len(posts),
			Posts:  postsRes,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// handlePostError converts the given error into an appropriate response.
func handlePostError(err error) APIError {
	switch err {
	case post.ErrPostActionForbidden:
		return ForbiddenRequest(err)
	case post.ErrPostNotFound,
		post.ErrPostAccountNotFound,
		post.ErrPostTopicNotFound:
		return NotFoundRequest(err)
	case post.ErrPostInvalid:
		return UnprocessableRequest(err)
	default:
		return BadRequest(err)
	}
}
