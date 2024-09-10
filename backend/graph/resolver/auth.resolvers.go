package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"backend/graph/generated"
	"backend/graph/graphModel"
	"context"
	"fmt"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, name string, email string, password string) (*graphModel.AuthPayload, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
