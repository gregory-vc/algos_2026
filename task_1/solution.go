package task1

import "fmt"

type Value struct {
	Speed int
	Loss  float64
}

type MatrixEdge struct {
	A, B int
	Value
}

type Matrix struct {
	Names  []string
	Values [][]*Value
}

func matrixNet(names []string, edges []MatrixEdge) *Matrix {
	net := &Matrix{Names: names, Values: make([][]*Value, len(names))}
	for i := range net.Values {
		net.Values[i] = make([]*Value, len(names))
	}
	for _, edge := range edges {
		net.Values[edge.A][edge.B] = &Value{edge.Speed, edge.Loss}
		net.Values[edge.B][edge.A] = &Value{edge.Speed, edge.Loss}
	}
	return net
}

func (net *Matrix) Print() {
	fmt.Println("Сеть:")
	for i := range net.Names {
		for j := i + 1; j < len(net.Names); j++ {
			value := net.Values[i][j]
			if value != nil {
				printEdge(net.Names[i], net.Names[j], *value)
			}
		}
	}
}

func printEdge(a, b string, value Value) {
	if a > b {
		a, b = b, a
	}
	fmt.Printf("  %s ──[%d МБ, потери %.0f%%]── %s\n", a, value.Speed, value.Loss*100, b)
}

type Edge struct {
	To *Node
	Value
}

type Node struct {
	Name  string
	Edges []Edge
}

func (n *Node) connect(other *Node, speed int, loss float64) {
	v := Value{speed, loss}
	n.Edges = append(n.Edges, Edge{other, v})
	other.Edges = append(other.Edges, Edge{n, v})
}

func (n *Node) Print() {
	fmt.Println("Сеть:")
	visited := map[*Node]bool{}
	var printNode func(*Node)
	printNode = func(node *Node) {
		if visited[node] {
			return
		}
		visited[node] = true
		for _, edge := range node.Edges {
			if !visited[edge.To] {
				printEdge(node.Name, edge.To.Name, edge.Value)
			}
		}
		for _, edge := range node.Edges {
			printNode(edge.To)
		}
	}
	printNode(n)
}

type NodeEdge struct {
	A, B *Node
	Value
}

func nodesNet(root *Node, edges []NodeEdge) *Node {
	for _, edge := range edges {
		edge.A.connect(edge.B, edge.Speed, edge.Loss)
	}
	return root
}
