package repository

import (
	"fmt"

	"github.com/bagashiz/kawan-sehat-backend/internal/app/comment"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/post"
	"github.com/bagashiz/kawan-sehat-backend/internal/app/reply"
	"github.com/bagashiz/kawan-sehat-backend/internal/postgres"
)

// calculateOffset sets the offset for the query.
func calculateOffset(limit, page int32) int32 {
	if limit == 0 && page == 0 {
		return 0
	}
	return (page - 1) * limit
}

// toPostDomain converts postgres repository results to domain entities.
func toPostDomain(results any) ([]*post.Post, error) {
	var posts []*post.Post
	switch res := results.(type) {
	case []postgres.SelectAllPostsRow:
		posts = make([]*post.Post, len(res))
		for i, result := range res {
			posts[i] = result.ToDomain()
		}
	case []postgres.SelectAllPostsPaginatedRow:
		posts = make([]*post.Post, len(res))
		for i, result := range res {
			posts[i] = result.ToDomain()
		}
	case []postgres.SelectPostsByAccountIDRow:
		posts = make([]*post.Post, len(res))
		for i, result := range res {
			posts[i] = result.ToDomain()
		}
	case []postgres.SelectPostsByAccountIDPaginatedRow:
		posts = make([]*post.Post, len(res))
		for i, result := range res {
			posts[i] = result.ToDomain()
		}
	case []postgres.SelectPostsByTopicIDRow:
		posts = make([]*post.Post, len(res))
		for i, result := range res {
			posts[i] = result.ToDomain()
		}
	case []postgres.SelectPostsByTopicIDPaginatedRow:
		posts = make([]*post.Post, len(res))
		for i, result := range res {
			posts[i] = result.ToDomain()
		}
	case []postgres.SelectBookmarksByAccountIDRow:
		posts = make([]*post.Post, len(res))
		for i, result := range res {
			posts[i] = result.ToDomain()
		}
	case []postgres.SelectBookmarksByAccountIDPaginatedRow:
		posts = make([]*post.Post, len(res))
		for i, result := range res {
			posts[i] = result.ToDomain()
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", results)
	}
	return posts, nil
}

// toCommentDomain converts postgres repository results to domain entities.
func toCommentDomain(results any) ([]*comment.Comment, error) {
	var comments []*comment.Comment
	switch res := results.(type) {
	case []postgres.SelectCommentsByPostIDRow:
		comments = make([]*comment.Comment, len(res))
		for i, result := range res {
			comments[i] = result.ToDomain()
		}
	case []postgres.SelectCommentsByPostIDPaginatedRow:
		comments = make([]*comment.Comment, len(res))
		for i, result := range res {
			comments[i] = result.ToDomain()
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", results)
	}
	return comments, nil
}

// toReplyDomain converts postgres repository results to domain entities.
func toReplyDomain(results any) ([]*reply.Reply, error) {
	var replies []*reply.Reply
	switch res := results.(type) {
	case []postgres.SelectRepliesByCommentIDRow:
		replies = make([]*reply.Reply, len(res))
		for i, result := range res {
			replies[i] = result.ToDomain()
		}
	case []postgres.SelectRepliesByCommentIDPaginatedRow:
		replies = make([]*reply.Reply, len(res))
		for i, result := range res {
			replies[i] = result.ToDomain()
		}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", results)
	}
	return replies, nil
}
