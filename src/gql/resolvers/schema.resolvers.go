package resolvers

import (
	"_/src/gql"
	"_/src/gql/models"
	"context"
	"fmt"
)

// CreateTodo is the resolver for the CreateTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - CreateTodo"))
}

// TodosGet is the resolver for the TodosGet field.
func (r *queryResolver) TodosGet(ctx context.Context) ([]models.Todo, error) {
	todos := []models.Todo{
		{
			ID:   "1",
			Text: "words",
			Done: false,
			User: &models.User{
				ID:   "1",
				Name: "Person",
			},
		},
	}

	return todos, nil
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
