package task11

import (
	"container/list"
	"reflect"
	"testing"

	"github.com/gregory-vc/algos_2026/task_11/magic"
)

func TestMagic(t *testing.T) {
	original := newList(3, 4, 5)
	counted := NewCountingList(original)

	magic.Test(counted)
	t.Logf("Add calls: %d", counted.AddCount())

	if got := counted.AddCount(); got != 3 {
		t.Fatalf("got %d add calls, want 3", got)
	}
	want := []int{3, 4, 5, 20}
	if got := listValues(original); !reflect.DeepEqual(got, want) {
		t.Fatalf("original list: got %v, want %v", got, want)
	}
}

func newList(values ...int) *list.List {
	result := list.New()
	for _, value := range values {
		result.PushBack(value)
	}
	return result
}

func listValues(values *list.List) []int {
	result := make([]int, 0, values.Len())
	for element := values.Front(); element != nil; element = element.Next() {
		result = append(result, element.Value.(int))
	}
	return result
}
