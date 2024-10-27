package repository

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/topic"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

// AddTopic inserts topic data to postgres database.
func (r *PostgresRepository) AddTopic(ctx context.Context, t *topic.Topic) error {
	arg := postgres.InsertTopicParams{
		ID:          t.ID,
		Name:        t.Name,
		Slug:        t.Slug,
		Description: t.Description,
	}

	if err := r.db.InsertTopic(ctx, arg); err != nil {
		return handleTopicError(err)
	}

	return nil
}

// GetTopicByID retrieves user account data from postgres database by ID.
func (r *PostgresRepository) GetTopicByID(ctx context.Context, id uuid.UUID) (*topic.Topic, error) {
	result, err := r.db.SelectTopicByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, topic.ErrTopicNotFound
		}
		return nil, err
	}

	topic := result.ToDomain()

	return topic, nil
}

// ListTopics retrieves all topics data from postgres database.
func (r *PostgresRepository) ListTopics(ctx context.Context, limit, offset int32) ([]*topic.Topic, error) {
	var results []postgres.Topic
	var topics []*topic.Topic
	var err error

	if limit == 0 && offset == 0 {
		results, err = r.db.SelectAllTopics(ctx)
	} else {
		results, err = r.db.SelectAllTopicsPaginated(ctx, postgres.SelectAllTopicsPaginatedParams{
			Limit:  limit,
			Offset: offset,
		})
	}
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		topic := result.ToDomain()
		topics = append(topics, topic)
	}
	return topics, nil
}

// UpdateTopic updates topic data in postgres database.
func (r *PostgresRepository) UpdateTopic(ctx context.Context, t *topic.Topic) error {
	arg := postgres.UpdateTopicParams{
		ID:          t.ID,
		Name:        pgtype.Text{String: t.Name, Valid: t.Name != ""},
		Slug:        pgtype.Text{String: t.Slug, Valid: t.Slug != ""},
		Description: pgtype.Text{String: t.Description, Valid: t.Description != ""},
	}
	if err := r.db.UpdateTopic(ctx, arg); err != nil {
		return handleTopicError(err)
	}
	return nil
}

// DeleteTopic removes topic data from postgres database.
func (r *PostgresRepository) DeleteTopic(ctx context.Context, id uuid.UUID) error {
	count, err := r.db.DeleteTopic(ctx, id)
	if err != nil {
		return err
	}

	if count == 0 {
		return topic.ErrTopicNotFound
	}

	return nil
}

// handleTopicError handles topic postgres repository errors and returns domain errors.
func handleTopicError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch {
		case pgerrcode.IsDataException(pgErr.Code):
			return topic.ErrTopicInvalid
		case pgerrcode.IsIntegrityConstraintViolation(pgErr.Code):
			return topic.ErrTopicDuplicateName
		default:
			return err
		}
	}
	return err
}
