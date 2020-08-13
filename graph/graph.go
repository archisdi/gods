package graph

import (
	"errors"
	"fmt"
)

// Path ...
type Path []*Node

func (p *Path) add(node *Node) {
	*p = append(*p, node)
}

func (p *Path) pop() {
	*p = (*p)[:len(*p)-1]
}

// Graph ...
type Graph map[int]*Node

func (g *Graph) getNode(val int) (*Node, error) {
	if node, ok := (*g)[val]; ok {
		return node, nil
	}
	return nil, errors.New("node not found")
}

func (g *Graph) contains(node *Node) bool {
	_, ok := (*g)[node.Data]
	return ok
}

func (g *Graph) add(node *Node) {
	(*g)[node.Data] = node
}

// DepthFirstSearch ...
func (g *Graph) DepthFirstSearch(source int, destination int) bool {
	src, _ := g.getNode(source)
	dst, _ := g.getNode(destination)
	visited := &Graph{}
	path := &Path{}

	hasDfs := g.hasDfs(src, dst, visited, path)

	// print path
	for _, p := range *path {
		fmt.Println(p.Data)
	}

	return hasDfs
}

func (g *Graph) hasDfs(source *Node, destination *Node, visited *Graph, path *Path) bool {
	if visited.contains(source) {
		return false
	}

	path.add(source)
	visited.add(source)

	if source.Data == destination.Data {
		return true
	}

	for _, child := range source.Connection {
		if g.hasDfs(child, destination, visited, path) {
			return true
		}
	}

	path.pop()
	return false
}

// Node ...
type Node struct {
	Data       int
	Connection Graph
}

// Edge ...
type Edge struct {
	X int
	Y int
}

// HasConnection ...
func (n *Node) HasConnection(value int) bool {
	_, ok := n.Connection[value]
	return ok
}

// Connect ...
func (n *Node) Connect(node *Node) {
	if !n.HasConnection(n.Data) {
		n.Connection[node.Data] = node
	}
}

// New ...
func New(nodes []int, Connection []Edge) Graph {
	gr := Graph{}

	for _, node := range nodes {
		gr[node] = &Node{
			Data:       node,
			Connection: Graph{},
		}
	}

	for _, edge := range Connection {
		parent, pOk := gr[edge.X]
		child, cOk := gr[edge.Y]
		if pOk && cOk {
			parent.Connect(child)
		}
	}

	return gr
}
