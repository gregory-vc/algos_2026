package task12

import (
	"fmt"
	"reflect"
)

type DB struct {
	data       []string
	converters map[reflect.Type]any
}

func NewDB() *DB {
	return &DB{converters: make(map[reflect.Type]any)}
}

func (db *DB) Add(value any) {
	db.data = append(db.data, fmt.Sprint(value))
}

func Register[T any](db *DB, converter func(string) (T, error)) {
	db.converters[reflect.TypeFor[T]()] = converter
}

func Get[T any](db *DB, index int) (T, error) {
	var zero T
	if index < 0 || index >= len(db.data) {
		return zero, fmt.Errorf("index %d out of range", index)
	}

	target := reflect.TypeFor[T]()
	converter, ok := db.converters[target].(func(string) (T, error))
	if !ok {
		return zero, fmt.Errorf("no converter for %v", target)
	}

	return converter(db.data[index])
}
