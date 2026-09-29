package task12

import (
	"fmt"
	"reflect"
)

type DB struct {
	data       []string
	converters map[reflect.Type]func(string) (any, error)
}

func NewDB() *DB {
	return &DB{converters: make(map[reflect.Type]func(string) (any, error))}
}

func (db *DB) Add(value any) {
	db.data = append(db.data, fmt.Sprint(value))
}

func Register[T any](db *DB, converter func(string) (T, error)) {
	db.converters[reflect.TypeFor[T]()] = func(value string) (any, error) {
		return converter(value)
	}
}

func (db *DB) Get(index int, target any) (any, error) {
	if index < 0 || index >= len(db.data) {
		return nil, fmt.Errorf("index %d out of range", index)
	}

	targetType := reflect.TypeOf(target)
	converter, ok := db.converters[targetType]
	if !ok {
		return nil, fmt.Errorf("no converter for %v", targetType)
	}

	return converter(db.data[index])
}
