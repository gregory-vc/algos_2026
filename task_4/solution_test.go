package task4

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{5, 5, 5, 5}, []int{5, 5, 5, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 3, 1, 2}, []int{1, 1, 2, 2, 3, 3}},
	}

	for _, test := range tests {
		quickSort(test.input)
		if !reflect.DeepEqual(test.input, test.want) {
			t.Fatalf("получено %v, ожидалось %v", test.input, test.want)
		}
	}
}
