package task4

type interval struct {
	left  int
	right int
}

func quickSort(values []int) {
	if len(values) < 2 {
		return
	}

	stack := []interval{{left: 0, right: len(values) - 1}}
	for len(stack) > 0 {
		last := len(stack) - 1
		left, right := stack[last].left, stack[last].right
		stack = stack[:last]

		i, j := left, right
		pivot := values[left+(right-left)/2]
		for i <= j {
			for values[i] < pivot {
				i++
			}
			for values[j] > pivot {
				j--
			}
			if i <= j {
				values[i], values[j] = values[j], values[i]
				i++
				j--
			}
		}

		if left < j {
			stack = append(stack, interval{left: left, right: j})
		}
		if i < right {
			stack = append(stack, interval{left: i, right: right})
		}
	}
}
