package repository

import (
	"context"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/user"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

// BookmarkPost relates an account to a post in postgres database.
func (r *PostgresRepository) BookmarkPost(ctx context.Context, bookmark *post.Bookmark) error {
	arg := postgres.InsertBookmarkParams{
		AccountID: bookmark.AccountID,
		PostID:    bookmark.PostID,
		CreatedAt: bookmark.CreatedAt,
	}

	if err := r.db.InsertBookmark(ctx, arg); err != nil {
		return handleBookmarkError(err)
	}

	return nil
}

// ListAccountBookmarks retrieves all bookmarked posts data related by an account from postgres database.
func (r *PostgresRepository) ListAccountBookmarks(
	ctx context.Context,
	accountID uuid.UUID,
	limit, page int32,
) ([]*post.Post, int64, error) {
	var results []postgres.Post
	var err error

	if limit == 0 {
		results, err = r.db.SelectBookmarksByAccountID(ctx, accountID)
	} else {
		results, err = r.db.SelectBookmarksByAccountIDPaginated(
			ctx, postgres.SelectBookmarksByAccountIDPaginatedParams{
				AccountID: accountID,
				Limit:     limit,
				Offset:    (page - 1) * limit,
			})
	}
	if err != nil {
		return nil, 0, err
	}

	count, err := r.db.CountAccountBookmarks(ctx, accountID)
	if err != nil {
		return nil, 0, err
	}

	bookmarks := make([]*post.Post, len(results))
	for i, result := range results {
		bookmarks[i] = result.ToDomain()
	}

	return bookmarks, count, err
}

// UnbookmarkPost removes a bookmarked post from an account in postgres database.
func (r *PostgresRepository) UnbookmarkPost(ctx context.Context, bookmark *post.Bookmark) error {
	arg := postgres.DeleteBookmarkParams{
		AccountID: bookmark.AccountID,
		PostID:    bookmark.PostID,
	}

	count, err := r.db.DeleteBookmark(ctx, arg)
	if err != nil {
		return err
	}

	if count == 0 {
		return post.ErrBookmarkAlreadyUnmarked
	}

	return nil
}

// handleBookmarkError handles bookmarks postgres repository errors and returns domain errors.
func handleBookmarkError(err error) error {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch {
		case pgerrcode.IsDataException(pgErr.Code):
			return post.ErrBookmarkInvalid
		case pgerrcode.IsIntegrityConstraintViolation(pgErr.Code):
			switch pgErr.ConstraintName {
			case "bookmarks_post_id_fkey":
				return post.ErrPostNotFound
			case "bookmarks_account_id_fkey":
				return user.ErrAccountNotFound
			default:
				return post.ErrBookmarkExists
			}
		default:
			return pgErr
		}
	}
	return err
}
