package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fasibio/gqlgensql/graph/generated"
)

// B is the resolver for the b field.
func (r *mutationResolver) B(ctx context.Context) (*int, error) {
	panic(fmt.Errorf("not implemented: B - b"))
}

// A is the resolver for the a field.
func (r *queryResolver) A(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: A - a"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
