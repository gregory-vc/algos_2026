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
	db.Add(4)
	db.Add("2,3")

	Register(db, func(value string) (string, error) { return value, nil })
	Register(db, strconv.Atoi)
	Register(db, func(value string) (Point, error) {
		var point Point
		_, err := fmt.Sscanf(value, "%d,%d", &point.X, &point.Y)
		return point, err
	})

	text, err := Get[string](db, 0)
	if err != nil || text != "4" {
		t.Fatalf("string: got %q, error %v; want %q", text, err, "4")
	}

	number, err := Get[int](db, 0)
	if err != nil || number != 4 {
		t.Fatalf("int: got %d, error %v; want 4", number, err)
	}

	point, err := Get[Point](db, 1)
	if err != nil || point != (Point{X: 2, Y: 3}) {
		t.Fatalf("Point: got %v, error %v; want {2 3}", point, err)
	}
}
