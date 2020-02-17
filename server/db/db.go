package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/robence/habit-tracker/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	GetHabits(name string) ([]*model.Habit, error)
	CreateHabit(name string) (*model.Habit, error)
}

type MongoDB struct {
	habitModel *mongo.Collection
	// habits []*model.Habit
}

func New(client *mongo.Client) *MongoDB {
	habits := client.Database("habits").Collection("habits")
	return &MongoDB{
		habitModel: habits,
	}
}

func (db MongoDB) GetHabits(name string) ([]*model.Habit, error) {
	res, err := db.habitModel.Find(context.TODO(), db.filter())
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
	fmt.Printf("queried habits with name %s\n", name)
	return p, nil
}

func (db MongoDB) CreateHabit(name string) (*model.Habit, error) {
	createdAt := time.Now().Format(time.RFC3339)
	id := primitive.NewObjectID().Hex()
	habit := model.Habit{id, name, createdAt}

	res, err := db.habitModel.InsertOne(context.TODO(), habit)

	if err != nil {
		log.Printf("Error while adding habit: %s", err.Error())
		return nil, err
	}

	fmt.Printf("inserted habit with ID %v\n", res.InsertedID)
	var h *model.Habit
	filter := bson.D{{"name", name}}
	if err := db.habitModel.FindOne(context.TODO(), filter).Decode(&h); err != nil {
		log.Fatal(err)
	}
	return h, nil
}

func (db MongoDB) filter() bson.D {
	return bson.D{}
}

// func (db MongoDB) filterByName(name string) bson.D {
// 	return bson.D{{Name: name}}
// }
