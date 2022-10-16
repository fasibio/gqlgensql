package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/fasibio/gqlgensql/graph/model"
)

// AddCat is the resolver for the addCat field.
func (r *mutationResolver) AddCat(ctx context.Context, input []*model.AddCatInput) (*model.AddCatPayload, error) {
	panic(fmt.Errorf("not implemented: AddCat - addCat"))
}

// UpdateCat is the resolver for the updateCat field.
func (r *mutationResolver) UpdateCat(ctx context.Context, input model.UpdateCatInput) (*model.UpdateCatPayload, error) {
	panic(fmt.Errorf("not implemented: UpdateCat - updateCat"))
}

// DeleteCat is the resolver for the deleteCat field.
func (r *mutationResolver) DeleteCat(ctx context.Context, filter model.CatFilter) (*model.DeleteCatPayload, error) {
	panic(fmt.Errorf("not implemented: DeleteCat - deleteCat"))
}

// AddTodo is the resolver for the addTodo field.
func (r *mutationResolver) AddTodo(ctx context.Context, input []*model.AddTodoInput) (*model.AddTodoPayload, error) {
	panic(fmt.Errorf("not implemented: AddTodo - addTodo"))
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodoInput) (*model.UpdateTodoPayload, error) {
	panic(fmt.Errorf("not implemented: UpdateTodo - updateTodo"))
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, filter model.TodoFilter) (*model.DeleteTodoPayload, error) {
	panic(fmt.Errorf("not implemented: DeleteTodo - deleteTodo"))
}

// AddUser is the resolver for the addUser field.
func (r *mutationResolver) AddUser(ctx context.Context, input []*model.AddUserInput) (*model.AddUserPayload, error) {
	panic(fmt.Errorf("not implemented: AddUser - addUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.UpdateUserPayload, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, filter model.UserFilter) (*model.DeleteUserPayload, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
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

// GetTodo is the resolver for the getTodo field.
func (r *queryResolver) GetTodo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: GetTodo - getTodo"))
}

// QueryTodo is the resolver for the queryTodo field.
func (r *queryResolver) QueryTodo(ctx context.Context, filter *model.TodoFilter, order *model.TodoOrder, first *int, offset *int) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: QueryTodo - queryTodo"))
}

// AggregateTodo is the resolver for the aggregateTodo field.
func (r *queryResolver) AggregateTodo(ctx context.Context, filter *model.TodoFilter) (*model.TodoAggregateResult, error) {
	panic(fmt.Errorf("not implemented: AggregateTodo - aggregateTodo"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// QueryUser is the resolver for the queryUser field.
func (r *queryResolver) QueryUser(ctx context.Context, filter *model.UserFilter, order *model.UserOrder, first *int, offset *int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: QueryUser - queryUser"))
}

// AggregateUser is the resolver for the aggregateUser field.
func (r *queryResolver) AggregateUser(ctx context.Context, filter *model.UserFilter) (*model.UserAggregateResult, error) {
	panic(fmt.Errorf("not implemented: AggregateUser - aggregateUser"))
}
