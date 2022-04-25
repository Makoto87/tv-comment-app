package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Makoto87/tv-comment-app/go-app/server/graph/generated"
	"github.com/Makoto87/tv-comment-app/go-app/server/graph/model"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PushLike(ctx context.Context, commentID int) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Programs(ctx context.Context, search string) ([]*model.Program, error) {
	var programs []*model.Program
	dummyLink := model.Program{
		ID:   1,
		Name: "test_name",
	}
	programs = append(programs, &dummyLink)
	return programs, nil
}

func (r *queryResolver) Episodes(ctx context.Context, input model.QueryEpisodesInput) ([]*model.Episode, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comments(ctx context.Context, episodeID int) ([]*model.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
