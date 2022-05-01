package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/Makoto87/tv-comment-app/go-app/server/dbcontrol"
	"github.com/Makoto87/tv-comment-app/go-app/server/graph/generated"
	"github.com/Makoto87/tv-comment-app/go-app/server/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (string, error) {
	err := dbcontrol.CreateComment(ctx, input.EpisodeID, input.UserID, input.Comment)
	if err != nil {
		log.Printf("CreateComment of mutationResolver %v", err)
		return "", gqlerror.Errorf("Server error: failed to create comment")
	}
	return "Success to create comment", nil
}

func (r *mutationResolver) PushLike(ctx context.Context, commentID int) (int, error) {
	likes, err := dbcontrol.UpdateCommentLikes(ctx, commentID)
	if err != nil {
		log.Printf("PushLike of mutationResolver %v", err)
		return likes, gqlerror.Errorf("Server error: failed to increase likes")
	}
	return likes, nil
}

func (r *queryResolver) Programs(ctx context.Context, search string) ([]*model.Program, error) {

	dbPrograms, err := dbcontrol.GetPrograms(ctx, search)
	if err != nil {
		log.Printf("Programs of queryResolver %v", err)
		return nil, gqlerror.Errorf("Server error: failed to get programs")
	}

	programs := make([]*model.Program, 0, len(dbPrograms))
	for _, p := range dbPrograms {
		programs = append(programs, &model.Program{ID: p.ID, Name: p.Name})
	}

	return programs, nil
}

func (r *queryResolver) Episodes(ctx context.Context, input model.QueryEpisodesInput) ([]*model.Episode, error) {

	dbEpisodes, err := dbcontrol.GetEpisodes(ctx, input.ProgramID, input.FromDate, input.ToDate)
	if err != nil {
		log.Printf("Episodes of queryResolver %v", err)
		return nil, gqlerror.Errorf("Server error: failed to get episodes")
	}

	episodes := make([]*model.Episode, 0, len(dbEpisodes))
	for _, e := range dbEpisodes {
		episodes = append(episodes, &model.Episode{ID: e.ID, Date: e.Date})
	}

	return episodes, nil
}

func (r *queryResolver) Comments(ctx context.Context, episodeID int) ([]*model.Comment, error) {

	dbComments, err := dbcontrol.GetComments(ctx, episodeID)
	if err != nil {
		log.Printf("Comments of queryResolver %v", err)
		return nil, gqlerror.Errorf("Server error: failed to get comments")
	}

	comments := make([]*model.Comment, 0, len(dbComments))
	for _, c := range dbComments {
		u := model.User{ID: c.User.ID, Name: c.User.Name}
		comments = append(comments, &model.Comment{ID: c.ID, Comment: c.Comment, Likes: c.Likes, User: &u, PostDate: c.PostDate})
	}

	return comments, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
