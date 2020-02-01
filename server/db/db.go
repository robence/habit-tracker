package db

import (
	"context"
	"github.com/robence/habit-tracker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type DB interface {
	GetHabits() ([]*model.Habit, error)
}

type MongoDB struct {
	collection *mongo.Collection
}

func New(client *mongo.Client) *MongoDB {
	habits := client.Database("habits").Collection("habits")
	return &MongoDB{
		// collection: programmers,
		collection: habits,
	}
}

func (db MongoDB) GetHabits() ([]*model.Habit, error) {
	res, err := db.collection.Find(context.TODO(), db.filter())
	if err != nil {
		log.Printf("Error while fetching habits: %s", err.Error())
		return nil, err
	}
	var p []*model.Habit
	err = res.All(context.TODO(), &p)
	if err != nil {
		log.Printf("Error while decoding habits: %s", err.Error())
		return nil, err
	}
	return p, nil
}

func (db MongoDB) filter() bson.D {
	return bson.D{}
}