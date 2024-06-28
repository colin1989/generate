//go:generate go run tobson-gen-go.go

package main_test

type User struct {
	ID    int64 `bson:"_id"`
	Level int64 `bson:"lv"`
}

type Task struct {
	ID    int64 `bson:"_id"`
	Level int64 `bson:"lv"`
	Mtx   int64 `bson:"-"`
	Empty string
}
