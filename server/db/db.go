package db

import (
	"context"
	"fmt"
	"log"

	"github.com/robence/habit-tracker/model"
	"go.mongodb.org/mongo-driver/bson"
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
	fmt.Println("Get Habits", name)
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
	return p, nil
}

func (db MongoDB) CreateHabit(name string) (*model.Habit, error) {
	// func (db MongoDB) CreateHabit(name string) (interface{}, err) {
	fmt.Println("Create Habit", name)

	habit := model.Habit{}
	// var p *model.User

	// habit.ID = primitive.NewObjectID()
	habit.Name = name
	// habit.CreatedAt = time.Now()
	// habit.UpdatedAt = time.Now()
	res, err := db.habitModel.InsertOne(context.TODO(), habit)
	fmt.Printf("inserted document with ID %v\n", res.InsertedID)

	if err != nil {
		log.Printf("Error while adding habit: %s", err.Error())
		return nil, err
	}

	var insertedHabit *model.Habit
	filter := bson.D{{"name", name}}
	documentReturned := db.habitModel.FindOne(context.TODO(), filter)
	if documentReturned == nil {
		log.Printf("Error while finding habit: %s")
		return nil, err
	}
	documentReturned.Decode(&insertedHabit)
	return insertedHabit, nil
}

func (db MongoDB) filter() bson.D {
	return bson.D{}
}

// func (db MongoDB) filterByName(name string) bson.D {
// 	return bson.D{{Name: name}}
// }
