package task2

type Result struct {
	Min int
	Max int
}

func less(a, b int) bool {
	return a < b
}

func findMinMax(values []int) Result {
	return findMinMaxBy(values, less)
}

func findMinMaxBy(values []int, less func(int, int) bool) Result {
	if len(values) == 0 {
		panic("empty slice")
	}
	min, max := values[0], values[0]
	for _, value := range values[1:] {
		if less(value, min) {
			min = value
		}
		if less(max, value) {
			max = value
		}
	}
	return Result{min, max}
}

func findMinMaxCompact(values []int) Result {
	return findMinMaxCompactBy(values, less)
}

func findMinMaxCompactBy(values []int, less func(int, int) bool) Result {
	if len(values) == 0 {
		panic("empty slice")
	}
	min, max, start := values[0], values[0], 1
	if len(values)%2 == 0 {
		if less(values[0], values[1]) {
			min, max = values[0], values[1]
		} else {
			min, max = values[1], values[0]
		}
		start = 2
	}
	for i := start; i < len(values); i += 2 {
		a, b := values[i], values[i+1]
		if less(b, a) {
			a, b = b, a
		}
		if less(a, min) {
			min = a
		}
		if less(max, b) {
			max = b
		}
	}
	return Result{min, max}
}
