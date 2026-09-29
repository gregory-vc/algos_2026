package task12

import (
	"fmt"
	"strconv"
	"testing"
)

type Point struct {
	X int
	Y int
}

func TestDB(t *testing.T) {
	db := NewDB()
	db.Add(17)
	db.Add("7.11")

	Register(db, func(value string) (string, error) { return value, nil })
	Register(db, strconv.Atoi)
	Register(db, func(value string) (Point, error) {
		var point Point
		_, err := fmt.Sscanf(value, "%d.%d", &point.X, &point.Y)
		return point, err
	})

	text, err := db.Get(0, "")
	if err != nil || text != "17" {
		t.Fatalf("string: got %q, error %v; want %q", text, err, "17")
	}

	number, err := db.Get(0, 0)
	if err != nil || number != 17 {
		t.Fatalf("int: got %d, error %v; want 17", number, err)
	}

	point, err := db.Get(1, Point{})
	if err != nil || point != (Point{X: 7, Y: 11}) {
		t.Fatalf("Point: got %v, error %v; want {7 11}", point, err)
	}
}
