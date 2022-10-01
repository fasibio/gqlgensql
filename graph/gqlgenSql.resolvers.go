package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fasibio/gqlgensql/graph/model"
)

// Empty is the resolver for the empty field.
func (r *queryResolver) Empty(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented: Empty - empty"))
}

// GetCat is the resolver for the getCat field.
func (r *queryResolver) GetCat(ctx context.Context, id string) (*model.Cat, error) {
	panic(fmt.Errorf("not implemented: GetCat - getCat"))
}

// QueryCat is the resolver for the queryCat field.
func (r *queryResolver) QueryCat(ctx context.Context, filter *model.CatFilter, order *model.CatOrder, first *int, offset *int) ([]*model.Cat, error) {
	panic(fmt.Errorf("not implemented: QueryCat - queryCat"))
}

// AggregateCat is the resolver for the aggregateCat field.
func (r *queryResolver) AggregateCat(ctx context.Context, filter *model.CatFilter) (*model.CatAggregateResult, error) {
	panic(fmt.Errorf("not implemented: AggregateCat - aggregateCat"))
}
