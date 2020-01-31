package gql

import (
	"context"
	"errors"
	"github.com/robence/habit-tracker/model"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

type MockDB struct {
	collection *mongo.Collection
}

// func (mockDB MockDB) GetProgrammers(string) ([]*model.Programmer, error) {
func (mockDB MockDB) GetHabits(string) ([]*model.Habit, error) {
	return []*model.Habit{{ID: "test-id"}}, errors.New("test-error")
}

// func TestProgrammers(t *testing.T) {
func TestHabits(t *testing.T) {
	r := &queryResolver{
		Resolver: &Resolver{&MockDB{}},
	}

	// programmers, err := r.Programmers(context.TODO(), "test")
	habits, err := r.Habits(context.TODO(), "test")

	// if programmers[0].ID != "test-id" {
	if habits[0].ID != "test-id" {
		// t.Errorf("GetProgrammers() got = %v, want test-id", programmers[0].ID)
		t.Errorf("GetHabits() got = %v, want test-id", habits[0].ID)
	}
	if err.Error() != "test-error" {
		// t.Errorf("GetProgrammers() got = %v, want test-error", err.Error())
		t.Errorf("GetHabits() got = %v, want test-error", err.Error())
	}
}
