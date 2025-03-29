package repository

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/topic"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

// AddPost inserts post data to postgres database.
func (r *PostgresRepository) AddPost(ctx context.Context, p *post.Post) error {
	arg := postgres.InsertPostParams{
		ID:        p.ID,
		AccountID: p.Account.ID,
		TopicID:   p.Topic.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
	if err := r.db.InsertPost(ctx, arg); err != nil {
		return handlePostError(err)
	}
	return nil
}

// UpdatePost updates topic data in postgres database.
func (r *PostgresRepository) UpdatePost(ctx context.Context, p *post.Post) error {
	arg := postgres.UpdatePostParams{
		ID:        p.ID,
		Title:     pgtype.Text{String: p.Title, Valid: p.Title != ""},
		Content:   pgtype.Text{String: p.Content, Valid: p.Content != ""},
		UpdatedAt: p.UpdatedAt,
	}
	if err := r.db.UpdatePost(ctx, arg); err != nil {
		return handlePostError(err)
	}
	return nil
}

// DeletePost removes post data from postgres database.
func (r *PostgresRepository) DeletePost(ctx context.Context, id uuid.UUID) error {
	count, err := r.db.DeletePost(ctx, id)
	if err != nil {
		return err
	}
	if count == 0 {
		return post.ErrPostNotFound
	}
	return nil
}

// GetPostByID retrieves post data from postgres database by ID.
func (r *PostgresRepository) GetPostByID(ctx context.Context, accountID, postID uuid.UUID) (*post.Post, error) {
	result, err := r.db.SelectPostByID(ctx, postgres.SelectPostByIDParams{
		AccountID: accountID, ID: postID,
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, post.ErrPostNotFound
		}
		return nil, err
	}
	post := result.ToDomain()
	return post, nil
}

// ListPosts retrieves all posts data from postgres database.
func (r *PostgresRepository) ListPosts(
	ctx context.Context, accountID uuid.UUID, limit, page int32,
) ([]*post.Post, int64, error) {
	offset := calculateOffset(limit, page)
	results, err := r.fetchAllPosts(ctx, accountID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	count, err := r.db.CountPosts(ctx)
	if err != nil {
		return nil, 0, err
	}
	posts, err := toPostDomain(results)
	if err != nil {
		return nil, 0, err
	}
	return posts, count, nil
}

// fetchAllPosts retrieves all posts data from postgres database.
func (r *PostgresRepository) fetchAllPosts(ctx context.Context, accountID uuid.UUID, limit, offset int32) (any, error) {
	if limit == 0 {
		return r.db.SelectAllPosts(ctx, accountID)
	}
	return r.db.SelectAllPostsPaginated(ctx, postgres.SelectAllPostsPaginatedParams{
		AccountID: accountID,
		Limit:     limit,
		Offset:    offset,
	})
}

// ListPostsByTopicID retrieves all posts data from postgres database by topic ID.
func (r *PostgresRepository) ListPostsByTopicID(
	ctx context.Context, accountID, topicID uuid.UUID, limit, page int32,
) ([]*post.Post, int64, error) {
	offset := calculateOffset(limit, page)
	results, err := r.fetchPostsByTopicID(ctx, accountID, topicID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	count, err := r.db.CountPostsByTopicID(ctx, topicID)
	if err != nil {
		return nil, 0, err
	}
	posts, err := toPostDomain(results)
	if err != nil {
		return nil, 0, err
	}
	return posts, count, err
}

// fetchPostsByTopicID retrieves all posts data from postgres database by topic ID.
func (r *PostgresRepository) fetchPostsByTopicID(
	ctx context.Context, accountID, topicID uuid.UUID, limit, offset int32,
) (any, error) {
	if limit == 0 {
		return r.db.SelectPostsByTopicID(ctx, postgres.SelectPostsByTopicIDParams{
			AccountID: accountID, TopicID: topicID,
		})
	}
	return r.db.SelectPostsByTopicIDPaginated(ctx, postgres.SelectPostsByTopicIDPaginatedParams{
		TopicID: topicID, Limit: limit, Offset: offset,
	})
}

// ListPostsByAccountID retrieves all posts data from postgres database by account ID.
func (r *PostgresRepository) ListPostsByAccountID(
	ctx context.Context, accountID uuid.UUID, limit, page int32,
) ([]*post.Post, int64, error) {
	offset := calculateOffset(limit, page)
	results, err := r.fetchPostsByAccountID(ctx, accountID, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	count, err := r.db.CountPostsByAccountID(ctx, accountID)
	if err != nil {
		return nil, 0, err
	}
	posts, err := toPostDomain(results)
	if err != nil {
		return nil, 0, err
	}
	return posts, count, err
}

// fetchPostsByAccountID retrieves all posts data from postgres database by account ID.
func (r *PostgresRepository) fetchPostsByAccountID(
	ctx context.Context, accountID uuid.UUID, limit, offset int32,
) (any, error) {
	if limit == 0 {
		return r.db.SelectPostsByAccountID(ctx, accountID)
	}
	return r.db.SelectPostsByAccountIDPaginated(ctx, postgres.SelectPostsByAccountIDPaginatedParams{
		AccountID: accountID, Limit: limit, Offset: offset,
	})
}

// handlePostError handles post postgres repository errors and returns domain errors.
func handlePostError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch {
		case pgerrcode.IsDataException(pgErr.Code):
			return post.ErrPostInvalid
		case pgerrcode.IsIntegrityConstraintViolation(pgErr.Code):
			switch pgErr.ConstraintName {
			case "posts_account_id_fkey":
				return user.ErrAccountNotFound
			case "posts_topic_id_fkey":
				return topic.ErrTopicNotFound
			default:
				return err
			}
		}
	}
	return err
}
