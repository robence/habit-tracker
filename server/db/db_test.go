package db

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestFilter(t *testing.T) {
	mongoDB := MongoDB{}

	got := mongoDB.filter()
	want := bson.D{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("filter() got = %v, want %v", got, want)
	}
}
