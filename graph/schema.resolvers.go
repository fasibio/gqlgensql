package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fasibio/gqlgensql/graph/generated"
	"github.com/fasibio/gqlgensql/graph/model"
)

// AddFood is the resolver for the addFood field.
func (r *mutationResolver) AddFood(ctx context.Context, name string, price int) (*model.CatFood, error) {
	panic(fmt.Errorf("not implemented: AddFood - addFood"))
}

// Catfoods is the resolver for the catfoods field.
func (r *queryResolver) Catfoods(ctx context.Context) ([]*model.CatFood, error) {
	panic(fmt.Errorf("not implemented: Catfoods - catfoods"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
