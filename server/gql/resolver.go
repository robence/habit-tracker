package gql

import (
	"context"
	"github.com/robence/habit-tracker/db"
	"github.com/robence/habit-tracker/gql/gen"
	"github.com/robence/habit-tracker/model"
)

type Resolver struct {
	DB db.DB
}

func (r *Resolver) Query() gen.QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

// func (r *queryResolver) Programmers(ctx context.Context, skill string) ([]*model.Programmer, error) {
func (r *queryResolver) Habits(ctx context.Context, name string) ([]*model.Habit, error) {
	// return r.DB.GetProgrammers(skill)
	return r.DB.GetHabits()
}
