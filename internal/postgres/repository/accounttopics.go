package repository

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/topic"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// RelateAccountToTopic inserts account topic data to postgres database.
func (r *PostgresRepository) RelateAccountToTopic(ctx context.Context, accountTopic *topic.AccountTopic) error {
	arg := postgres.InsertAccountTopicParams{
		AccountID: accountTopic.AccountID,
		TopicID:   accountTopic.TopicID,
		CreatedAt: accountTopic.CreatedAt,
	}

	if err := r.db.InsertAccountTopic(ctx, arg); err != nil {
		return handleAccountTopicError(err)
	}

	return nil
}

// ListFollowedTopics retrieves all topics data related by an account from postgres database.
func (r *PostgresRepository) ListFollowedTopics(
	ctx context.Context,
	accountID uuid.UUID,
	limit, page int32,
) ([]*topic.FollowedTopic, int64, error) {
	var results []postgres.Topic
	var err error

	if limit == 0 {
		results, err = r.db.SelectTopicsByAccountID(ctx, accountID)
	} else {
		results, err = r.db.SelectTopicsByAccountIDPaginated(
			ctx, postgres.SelectTopicsByAccountIDPaginatedParams{
				AccountID: accountID,
				Limit:     limit,
				Offset:    (page - 1) * limit,
			})
	}
	if err != nil {
		return nil, 0, err
	}

	count, err := r.db.CountTopicsByAccountID(ctx, accountID)
	if err != nil {
		return nil, 0, err
	}

	topics := make([]*topic.FollowedTopic, len(results))
	for i, result := range results {
		topics[i] = &topic.FollowedTopic{Topic: result.ToDomain()}
	}

	return topics, count, err
}

// UnrelateAccountFromTopic deletes account topic data from postgres database.
func (r *PostgresRepository) UnrelateAccountFromTopic(ctx context.Context, accountTopic *topic.AccountTopic) error {
	arg := postgres.DeleteAccountTopicParams{
		AccountID: accountTopic.AccountID,
		TopicID:   accountTopic.TopicID,
	}

	count, err := r.db.DeleteAccountTopic(ctx, arg)
	if err != nil {
		return err
	}

	if count == 0 {
		return topic.ErrAccountAlreadyUnfollowedTopic
	}

	return nil
}

// handleAccountTopicError handles account topics postgres repository errors and returns domain errors.
func handleAccountTopicError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch {
		case pgerrcode.IsDataException(pgErr.Code):
			return topic.ErrAccountTopicsInvalid
		case pgerrcode.IsIntegrityConstraintViolation(pgErr.Code):
			switch pgErr.ConstraintName {
			case "account_topics_topic_id_fkey":
				return topic.ErrTopicNotFound
			default:
				return topic.ErrAccountTopicsAlreadyExists
			}
		default:
			return pgErr
		}
	}
	return err
}
