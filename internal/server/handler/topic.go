package handler

import (
	"net/http"
	"time"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/topic"
)

// topicResponse holds the response data for the topic object.
type topicResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// createTopicRequest holds the request data for the create topic handler.
type createTopicRequest struct {
	Name        string `json:"name" validate:"required,lte=50"`
	Description string `json:"description" validate:"required,lte=500"`
}

// CreateTopic is the handler for the topic creation route.
func (h *Handler) CreateTopic() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req createTopicRequest
		if err := h.decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}

		topic, err := h.topicSvc.AddTopic(
			r.Context(),
			topic.CreateTopicParams{
				Name:        req.Name,
				Description: req.Description,
			})
		if err != nil {
			return handleTopicError(err)
		}

		res := &topicResponse{
			ID:          topic.ID.String(),
			Name:        topic.Name,
			Slug:        topic.Slug,
			Description: topic.Description,
			CreatedAt:   topic.CreatedAt,
			UpdatedAt:   topic.UpdatedAt,
		}

		return writeJSON(w, http.StatusCreated, res, nil)
	}
}

type updateTopicRequest struct {
	Name        string `json:"name" validate:"lte=50"`
	Description string `json:"description" validate:"lte=500"`
}

// UpdateTopic is the handler for the topic update route.
func (h *Handler) UpdateTopic() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		var req updateTopicRequest
		if err := h.decodeAndValidateJSONRequest(r, &req); err != nil {
			return err
		}

		id := r.PathValue("id")

		topic, err := h.topicSvc.UpdateTopic(
			r.Context(),
			topic.UpdateTopicParams{
				ID:          id,
				Name:        req.Name,
				Description: req.Description,
			})
		if err != nil {
			return handleTopicError(err)
		}

		res := &topicResponse{
			ID:          topic.ID.String(),
			Name:        topic.Name,
			Slug:        topic.Slug,
			Description: topic.Description,
			CreatedAt:   topic.CreatedAt,
			UpdatedAt:   topic.UpdatedAt,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// DeleteTopic is the handler for the topic deletion route.
func (h *Handler) DeleteTopic() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")

		err := h.topicSvc.DeleteTopic(r.Context(), id)
		if err != nil {
			return handleTopicError(err)
		}

		return writeJSON(w, http.StatusOK, nil, nil)
	}
}

// GetTopicByID is the handler for the topic retrieval route.
func (h *Handler) GetTopicByID() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")

		topic, err := h.topicSvc.GetTopicByID(r.Context(), id)
		if err != nil {
			return handleTopicError(err)
		}

		res := &topicResponse{
			ID:          topic.ID.String(),
			Name:        topic.Name,
			Slug:        topic.Slug,
			Description: topic.Description,
			CreatedAt:   topic.CreatedAt,
			UpdatedAt:   topic.UpdatedAt,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// listTopicsResponse holds the response data for the list topics handler.
type listTopicsResponse struct {
	Limit  int32           `json:"limit"`
	Offset int32           `json:"offset"`
	Count  int             `json:"count"`
	Topics []topicResponse `json:"topics"`
}

// ListTopics is the handler for the topic list route.
func (h *Handler) ListTopics() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		limit, offset := getLimitOffset(r)

		topics, err := h.topicSvc.ListTopics(r.Context(), limit, offset)
		if err != nil {
			return handleTopicError(err)
		}

		topicRes := make([]topicResponse, len(topics))
		for i, t := range topics {
			topicRes[i] = topicResponse{
				ID:          t.ID.String(),
				Name:        t.Name,
				Slug:        t.Slug,
				Description: t.Description,
				CreatedAt:   t.CreatedAt,
				UpdatedAt:   t.UpdatedAt,
			}
		}

		res := &listTopicsResponse{
			Limit:  limit,
			Offset: offset,
			Count:  len(topics),
			Topics: topicRes,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// handleTopicError determines the appropriate HTTP status code for the topic service error.
func handleTopicError(err error) APIError {
	switch err {
	case topic.ErrTopicNotFound:
		return NotFoundRequest(err)
	case topic.ErrTopicDuplicateName:
		return ConflictRequest(err)
	case topic.ErrTopicInvalid:
		return UnprocessableRequest(err)
	default:
		return BadRequest(err)
	}
}
