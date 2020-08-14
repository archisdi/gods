package graph

import (
	"errors"
	"fmt"
)

// Stack ...
type Stack []*Node

func (p *Stack) push(node *Node) {
	*p = append(*p, node)
}

func (p *Stack) add(node *Node) {
	*p = append([]*Node{node}, *p...)
}

func (p *Stack) pop() {
	*p = (*p)[:len(*p)-1]
}
func (p *Stack) dequeue() *Node {
	node := (*p)[0]
	*p = (*p)[1:]
	return node
}

func (p *Stack) isEmpty() bool {
	return len(*p) == 0
}

func (p *Stack) contains(node *Node) bool {
	for _, s := range *p {
		if s.Data == node.Data {
			return true
		}
	}
	return false
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
func (g *Graph) DepthFirstSearch(source int, destination int) {
	src, _ := g.getNode(source)
	dst, _ := g.getNode(destination)
	visited := &Graph{}
	Path := &Stack{}

	g.hasDfs(src, dst, visited, Path)

	// print Stack
	for _, p := range *Path {
		fmt.Println(p.Data)
	}
}

func (g *Graph) hasDfs(source *Node, destination *Node, visited *Graph, path *Stack) bool {
	if visited.contains(source) {
		path.pop()
		return false
	}

	path.push(source)
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

// BreadthFirstSearch ...
func (g *Graph) BreadthFirstSearch(source int, destination int) bool {
	src, _ := g.getNode(source)
	dst, _ := g.getNode(destination)
	Queue := &Stack{}
	visited := &Graph{}

	Queue.push(src)
	for !Queue.isEmpty() {
		currentNode := Queue.dequeue()
		if currentNode.Data == dst.Data {
			break
		}
		visited.add(currentNode)
		for _, child := range currentNode.Connection {
			if !visited.contains(child) && !Queue.contains(child) {
				Queue.push(child)
				child.Parent = currentNode
			}
		}
	}

	paths := Stack{}
	path := dst
	for {
		paths.add(path)
		if path.Parent == nil {
			break
		}
		path = path.Parent
	}

	for _, p := range paths {
		fmt.Println(p.Data)
	}

	return dst.Parent != nil
}

// Node ...
type Node struct {
	Data       int
	Connection Graph
	Parent     *Node
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
			// uncomment for two way graph-
			// child.Connect(parent)
		}
	}

	return gr
}
