package handler

import (
	"net/http"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/topic"
)

// FollowTopic is the handler for adding a topic to the authenticated user's followed topics.
func (h *Handler) FollowTopic() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")

		if err := h.topicSvc.FollowTopic(r.Context(), id); err != nil {
			return handleAccountTopicError(err)
		}

		return writeJSON(w, http.StatusOK, nil, nil)
	}
}

// UnfollowTopic is the handler for removing a topic from the authenticated user's followed topics.
func (h *Handler) UnfollowTopic() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.PathValue("id")

		if err := h.topicSvc.UnfollowTopic(r.Context(), id); err != nil {
			return handleAccountTopicError(err)
		}

		return writeJSON(w, http.StatusOK, nil, nil)
	}
}

// listFollowedTopicsResponse holds the response data for the list followed topics handler.
type listFollowedTopicsResponse struct {
	Limit  int32                  `json:"limit"`
	Page   int32                  `json:"page"`
	Count  int64                  `json:"count"`
	Topics []*topic.FollowedTopic `json:"topics"`
}

// ListFollowedTopics is the handler for fetching the authenticated user's followed topics.
func (h *Handler) ListFollowedTopics() APIFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		limit, page := getLimitPage(r)

		topics, count, err := h.topicSvc.ListFollowedTopics(r.Context(), limit, page)
		if err != nil {
			return handleAccountTopicError(err)
		}

		res := &listFollowedTopicsResponse{
			Limit:  limit,
			Page:   page,
			Count:  count,
			Topics: topics,
		}

		return writeJSON(w, http.StatusOK, res, nil)
	}
}

// handleAccountTopicError determines the appropriate HTTP status code for the accounttopic service error.
func handleAccountTopicError(err error) APIError {
	switch err {
	case topic.ErrTopicNotFound:
		return NotFoundRequest(err)
	case topic.ErrAccountAlreadyUnfollowedTopic, topic.ErrAccountTopicsAlreadyExists:
		return ConflictRequest(err)
	case topic.ErrAccountTopicsInvalid:
		return UnprocessableRequest(err)
	default:
		return BadRequest(err)
	}
}
