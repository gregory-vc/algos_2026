package task5

type Edge struct {
	From int
	To   int
}

func adjacencyMatrixToEdgeList(matrix [][]int) []Edge {
	edges := []Edge{}
	for from, row := range matrix {
		for to, value := range row {
			if value == 1 {
				edges = append(edges, Edge{From: from, To: to})
			}
		}
	}
	return edges
}

func topologicalSort(matrix [][]int) []int {
	inDegree := make([]int, len(matrix))
	for _, row := range matrix {
		for to, value := range row {
			if value == 1 {
				inDegree[to]++
			}
		}
	}

	queue := []int{}
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	result := make([]int, 0, len(matrix))
	for first := 0; first < len(queue); first++ {
		node := queue[first]
		result = append(result, node)
		for to, value := range matrix[node] {
			if value == 1 {
				inDegree[to]--
				if inDegree[to] == 0 {
					queue = append(queue, to)
				}
			}
		}
	}

	if len(result) != len(matrix) {
		return nil
	}
	return result
}
