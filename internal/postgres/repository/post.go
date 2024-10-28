package repository

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
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
		AccountID: p.AccountID,
		TopicID:   p.TopicID,
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
func (r *PostgresRepository) GetPostByID(ctx context.Context, id uuid.UUID) (*post.Post, error) {
	result, err := r.db.SelectPostByID(ctx, id)
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
func (r *PostgresRepository) ListPosts(ctx context.Context, limit, page int32) ([]*post.Post, int64, error) {
	var results []postgres.Post
	var err error

	offset := (page - 1) * limit
	if limit == 0 {
		results, err = r.db.SelectAllPosts(ctx)
	} else {
		results, err = r.db.SelectAllPostsPaginated(ctx, postgres.SelectAllPostsPaginatedParams{
			Limit: limit, Offset: offset,
		})
	}
	if err != nil {
		return nil, 0, err
	}

	count, err := r.db.CountPosts(ctx)
	if err != nil {
		return nil, 0, err
	}

	posts := make([]*post.Post, len(results))
	for i, result := range results {
		posts[i] = result.ToDomain()
	}

	return posts, count, nil
}

// ListPostsByTopicID retrieves all posts data from postgres database by topic ID.
func (r *PostgresRepository) ListPostsByTopicID(
	ctx context.Context, topicID uuid.UUID, limit, page int32,
) ([]*post.Post, int64, error) {
	var results []postgres.Post
	var err error

	offset := (page - 1) * limit
	if limit == 0 {
		results, err = r.db.SelectPostsByTopicID(ctx, topicID)
	} else {
		results, err = r.db.SelectPostsByTopicIDPaginated(ctx, postgres.SelectPostsByTopicIDPaginatedParams{
			TopicID: topicID, Limit: limit, Offset: offset,
		})
	}
	if err != nil {
		return nil, 0, err
	}

	count, err := r.db.CountPostsByTopicID(ctx, topicID)
	if err != nil {
		return nil, 0, err
	}

	posts := make([]*post.Post, len(results))
	for i, result := range results {
		posts[i] = result.ToDomain()
	}
	return posts, count, err
}

// ListPostsByAccountID retrieves all posts data from postgres database by account ID.
func (r *PostgresRepository) ListPostsByAccountID(
	ctx context.Context, accountID uuid.UUID, limit, page int32,
) ([]*post.Post, int64, error) {
	var results []postgres.Post
	var err error

	offset := (page - 1) * limit
	if limit == 0 {
		results, err = r.db.SelectPostsByAccountID(ctx, accountID)
	} else {
		results, err = r.db.SelectPostsByAccountIDPaginated(ctx, postgres.SelectPostsByAccountIDPaginatedParams{
			AccountID: accountID, Limit: limit, Offset: offset,
		})
	}
	if err != nil {
		return nil, 0, err
	}

	count, err := r.db.CountPostsByAccountID(ctx, accountID)
	if err != nil {
		return nil, 0, err
	}

	posts := make([]*post.Post, len(results))
	for i, result := range results {
		posts[i] = result.ToDomain()
	}

	return posts, count, nil
}

// handlePostError handles post postgres repository errors and returns domain errors.
func handlePostError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch {
		case pgerrcode.IsDataException(pgErr.Code):
			return post.ErrPostInvalid
		case pgerrcode.IsIntegrityConstraintViolation(pgErr.Code):
			return err
		default:
			return err
		}
	}
	return err
}
