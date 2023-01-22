package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"rs/internal/graph/generated"
	"rs/internal/graph/model"
)

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) (*model.Todo, error) {
	return &model.Todo{
		Text: "success",
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
