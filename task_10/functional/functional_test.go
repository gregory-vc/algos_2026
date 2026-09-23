package functional_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/gregory-vc/algos_2026/task_10/functional"
)

func TestMap(t *testing.T) {
	for name, transform := range map[string]func([]int, func(int) string) []string{
		"Map":         functional.Map[int, string],
		"ParallelMap": functional.ParallelMap[int, string],
	} {
		t.Run(name, func(t *testing.T) {
			got := transform([]int{3, 1, 2, 1}, strconv.Itoa)
			want := []string{"3", "1", "2", "1"}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
			got = transform(nil, func(int) string {
				t.Error("transform called for empty input")
				return ""
			})
			if len(got) != 0 {
				t.Fatalf("expected empty result, got %v", got)
			}
		})
	}
}

func TestFilterDoesNotChangeInput(t *testing.T) {
	input := []int{1, 2, 3, 4, 2}
	got := functional.Filter(input, func(n int) bool { return n%2 == 0 })
	if !reflect.DeepEqual(got, []int{2, 4, 2}) {
		t.Fatalf("got %v, want [2 4 2]", got)
	}
	got[0] = 100
	if !reflect.DeepEqual(input, []int{1, 2, 3, 4, 2}) {
		t.Fatalf("input slice was modified: %v", input)
	}
}

func TestGroupBy(t *testing.T) {
	got := functional.GroupBy([]int{3, 2, 1, 4, 3}, func(n int) (bool, string) {
		return n%2 == 0, strconv.Itoa(n)
	})
	want := map[bool][]string{false: {"3", "1", "3"}, true: {"2", "4"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
