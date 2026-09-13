package task5

import (
	"fmt"
	"reflect"
	"testing"
)

var matrix = [][]int{
	{0, 1, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 0},
	{0, 0, 0, 0, 0, 0, 1, 1},
	{0, 0, 0, 0, 1, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0},
}

func TestAdjacencyMatrixToEdgeList(t *testing.T) {
	want := []Edge{
		{From: 0, To: 1},
		{From: 0, To: 2},
		{From: 1, To: 3},
		{From: 2, To: 3},
		{From: 2, To: 4},
		{From: 3, To: 6},
		{From: 4, To: 6},
		{From: 4, To: 7},
		{From: 5, To: 4},
		{From: 5, To: 7},
		{From: 6, To: 7},
	}

	fmt.Println("До: матрица смежности")
	printMatrix(matrix)
	got := adjacencyMatrixToEdgeList(matrix)
	fmt.Println("После: список рёбер")
	for _, edge := range got {
		fmt.Printf("  %d -> %d\n", edge.From, edge.To)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("получено %v, ожидалось %v", got, want)
	}
}

func TestTopologicalSort(t *testing.T) {
	fmt.Println("До: матрица смежности")
	printMatrix(matrix)
	order := topologicalSort(matrix)
	fmt.Printf("После: топологический порядок %v\n", order)
	if len(order) != len(matrix) {
		t.Fatalf("получен неполный порядок: %v", order)
	}

	position := make([]int, len(order))
	for index, node := range order {
		position[node] = index
	}
	for from, row := range matrix {
		for to, edge := range row {
			if edge == 1 && position[from] >= position[to] {
				t.Fatalf("ребро %d -> %d нарушает порядок %v", from, to, order)
			}
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(" ", row)
	}
}
