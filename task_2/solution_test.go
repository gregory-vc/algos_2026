package task2

import "testing"

func TestFindMinMaxCompact(t *testing.T) {
	if got := findMinMaxCompact([]int{1, 2, 3, 4, 5}); got != (Result{1, 5}) {
		t.Fatalf("пример из задания: получено %+v", got)
	}

	values := make([]int, 1000)
	for i := range values {
		values[i] = len(values) - i
	}

	want := Result{Min: 1, Max: 1000}
	if got := findMinMax(values); got != want {
		t.Fatalf("findMinMax: получено %+v, ожидалось %+v", got, want)
	}
	if got := findMinMaxCompact(values); got != want {
		t.Fatalf("findMinMaxCompact: получено %+v, ожидалось %+v", got, want)
	}

	originalComparisons, compactComparisons := 0, 0
	findMinMaxBy(values, func(a, b int) bool {
		originalComparisons++
		return a < b
	})
	findMinMaxCompactBy(values, func(a, b int) bool {
		compactComparisons++
		return a < b
	})
	t.Logf("сравнений: исходный — %d, улучшенный — %d (на %d меньше)",
		originalComparisons, compactComparisons, originalComparisons-compactComparisons)
}
