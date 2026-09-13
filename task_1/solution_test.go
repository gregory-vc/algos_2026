package task1

import "testing"

func TestMatrixNet(_ *testing.T) {
	net := matrixNet([]string{"A", "B", "C", "D", "E", "F"}, []MatrixEdge{
		{0, 1, Value{1500, .90}},
		{0, 2, Value{2000, .10}},
		{0, 3, Value{1000, .50}},
		{1, 5, Value{1500, .60}},
		{2, 4, Value{900, .05}},
		{2, 5, Value{500, .20}},
		{3, 4, Value{2500, .01}},
		{4, 5, Value{300, .85}},
	})
	net.Print()
}

func TestNodesNet(_ *testing.T) {
	a, b, c := &Node{Name: "A"}, &Node{Name: "B"}, &Node{Name: "C"}
	d, e, f := &Node{Name: "D"}, &Node{Name: "E"}, &Node{Name: "F"}
	root := nodesNet(a, []NodeEdge{
		{a, b, Value{1500, .90}},
		{a, c, Value{2000, .10}},
		{a, d, Value{1000, .50}},
		{b, f, Value{1500, .60}},
		{c, e, Value{900, .05}},
		{c, f, Value{500, .20}},
		{d, e, Value{2500, .01}},
		{e, f, Value{300, .85}},
	})
	root.Print()
}
