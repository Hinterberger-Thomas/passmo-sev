package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Hinterberger-Thomas/passmo-sev/db"
	"github.com/Hinterberger-Thomas/passmo-sev/graph/generated"
	"github.com/Hinterberger-Thomas/passmo-sev/graph/model"
)

var dbc *db.DB = db.InitDB()

func (r *mutationResolver) CreUse(ctx context.Context, input string) (bool, error) {
	return dbc.InsNewAcc(input)
}

func (r *mutationResolver) UpdAcc(ctx context.Context, input string) (bool, error) {
	return dbc.UpdAccData(email, input)
}

func (r *queryResolver) Passwords(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
