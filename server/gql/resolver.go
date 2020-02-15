package gql

import (
	"context"

	"github.com/robence/habit-tracker/db"
	"github.com/robence/habit-tracker/gql/gen"
	"github.com/robence/habit-tracker/model"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	DB db.DB
}

func (r *Resolver) Mutation() gen.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateHabit(ctx context.Context, name string) (*model.Habit, error) {
	return r.DB.CreateHabit(name)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Habits(ctx context.Context, name string) ([]*model.Habit, error) {
	return r.DB.GetHabits(name)
}
