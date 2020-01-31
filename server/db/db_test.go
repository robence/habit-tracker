package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	// mongoDB := MongoDB{}

	got := mongoDB.filter()
	// got := "a"
	// want := "b"
	want := bson.D{}

	if !reflect.DeepEqual(got, want) {
	// if "Hello Dude!" != "Hello Dude!" {
		t.Errorf("filter() got = %v, want %v", got, want)
	}
}
