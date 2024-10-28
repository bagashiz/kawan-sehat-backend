package handler

import (
	"net/http"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
)

// postResponse holds the response data for the post object.
type postResponse struct {
	ID        string               `json:"id"`
	Account   *postAccountResponse `json:"account,omitempty"`
	Topic     *postTopicResponse   `json:"topic,omitempty"`
	Title     string               `json:"title"`
	Content   string               `json:"content"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

// postTopicResponse holds the response data for the post account object.
type postAccountResponse struct {
	ID       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

// postTopicResponse holds the response data for the post topic object.
type postTopicResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
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
			ID: post.ID.String(),
			Account: &postAccountResponse{
				ID:       post.Account.ID.String(),
				Username: post.Account.Username,
			},
			Topic: &postTopicResponse{
				ID:   post.Topic.ID.String(),
				Name: post.Topic.Name,
			},
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
			ID: post.ID.String(),
			Account: &postAccountResponse{
				ID:       post.Account.ID.String(),
				Username: post.Account.Username,
			},
			Topic: &postTopicResponse{
				ID:   post.Topic.ID.String(),
				Name: post.Topic.Name,
			},
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
	Limit int32           `json:"limit"`
	Page  int32           `json:"page"`
	Count int64           `json:"count"`
	Posts []*postResponse `json:"posts"`
}

// ListPosts is the handler for the post list route.
func (h *Handler) ListPosts() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		limit, page := getLimitPage(r)
		accountID := r.URL.Query().Get("account_id")
		topicID := r.URL.Query().Get("topic_id")

		posts, count, err := h.postSvc.ListPosts(
			r.Context(),
			post.ListPostsParams{
				AccountID: accountID,
				TopicID:   topicID,
				Limit:     limit,
				Page:      page,
			})
		if err != nil {
			return handlePostError(err)
		}

		postsRes := make([]*postResponse, len(posts))
		for i, p := range posts {
			postsRes[i] = &postResponse{
				ID: p.ID.String(),
				Account: &postAccountResponse{
					ID:       p.Account.ID.String(),
					Username: p.Account.Username,
				},
				Topic: &postTopicResponse{
					ID:   p.Topic.ID.String(),
					Name: p.Topic.Name,
				},
				Title:     p.Title,
				Content:   p.Content,
				CreatedAt: p.CreatedAt,
				UpdatedAt: p.UpdatedAt,
			}
		}

		res := listPostsResponse{
			Limit: limit,
			Page:  page,
			Count: count,
			Posts: postsRes,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// Bookmark is the handler for the post bookmark route.
func (h *Handler) Bookmark() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		err := h.postSvc.BookmarkPost(r.Context(), id)
		if err != nil {
			return handlePostError(err)
		}
		return writeJSON(w, http.StatusOK, nil, nil)
	}
}

// Unbookmark is the handler for the post unbookmark route.
func (h *Handler) Unbookmark() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")
		err := h.postSvc.UnbookmarkPost(r.Context(), id)
		if err != nil {
			return handlePostError(err)
		}
		return writeJSON(w, http.StatusOK, nil, nil)
	}
}

// ListBookmarks is the handler for the post bookmarks route.
func (h *Handler) ListBookmarks() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		limit, page := getLimitPage(r)

		posts, count, err := h.postSvc.ListBookmarks(r.Context(), limit, page)
		if err != nil {
			return handlePostError(err)
		}

		postsRes := make([]*postResponse, len(posts))
		for i, p := range posts {
			postsRes[i] = &postResponse{
				ID: p.ID.String(),
				Account: &postAccountResponse{
					ID:       p.Account.ID.String(),
					Username: p.Account.Username,
				},
				Topic: &postTopicResponse{
					ID:   p.Topic.ID.String(),
					Name: p.Topic.Name,
				},
				Title:     p.Title,
				Content:   p.Content,
				CreatedAt: p.CreatedAt,
				UpdatedAt: p.UpdatedAt,
			}
		}

		res := listPostsResponse{
			Limit: limit,
			Page:  page,
			Count: count,
			Posts: postsRes,
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
