package db

import (
	"context"
	"github.com/robence/habit-tracker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type DB interface {
	// GetProgrammers(skill string) ([]*model.Programmer, error)
	GetHabits() ([]*model.Habit, error)
}

type MongoDB struct {
	collection *mongo.Collection
}

func New(client *mongo.Client) *MongoDB {
	// programmers := client.Database("programmers").Collection("programmers")
	habits := client.Database("habits").Collection("habits")
	return &MongoDB{
		// collection: programmers,
		collection: habits,
	}
}

// func (db MongoDB) GetProgrammers(skill string) ([]*model.Programmer, error) {
func (db MongoDB) GetHabits() ([]*model.Habit, error) {
	// res, err := db.collection.Find(context.TODO(), db.filter(skill))
	res, err := db.collection.Find(context.TODO(), db.filter())
	if err != nil {
		// log.Printf("Error while fetching programmers: %s", err.Error())
		log.Printf("Error while fetching habits: %s", err.Error())
		return nil, err
	}
	// var p []*model.Programmer
	var p []*model.Habit
	err = res.All(context.TODO(), &p)
	if err != nil {
		// log.Printf("Error while decoding programmers: %s", err.Error())
		log.Printf("Error while decoding habits: %s", err.Error())
		return nil, err
	}
	return p, nil
}

// func (db MongoDB) filter(skill string) bson.D {
// 	return bson.D{{
// 		"skills.name",
// 		bson.D{{
// 			"$regex",
// 			"^" + skill + ".*$",
// 		}, {
// 			"$options",
// 			"i",
// 		}},
// 	}}
// }

func (db MongoDB) filter() bson.D {
	return bson.D{}
}