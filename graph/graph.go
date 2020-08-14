package graph

import (
	"ds/list"
	"errors"
	"fmt"
)

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
	Path := &list.Stack{}

	g.hasDfs(src, dst, visited, Path)

	// print Stack
	for _, p := range *Path {
		fmt.Println(p.(int))
	}
}

func (g *Graph) hasDfs(source *Node, destination *Node, visited *Graph, path *list.Stack) bool {
	if visited.contains(source) {
		path.Pop()
		return false
	}

	path.Push(source.Data)
	visited.add(source)

	if source.Data == destination.Data {
		return true
	}

	for _, child := range source.Connection {
		if g.hasDfs(child, destination, visited, path) {
			return true
		}
	}

	path.Pop()
	return false
}

// BreadthFirstSearch ...
func (g *Graph) BreadthFirstSearch(source int, destination int) bool {
	src, _ := g.getNode(source)
	dst, _ := g.getNode(destination)
	Queue := &list.Queue{}
	visited := &Graph{}

	Queue.Enqueue(src.Data)
	for !Queue.IsEmpty() {
		nodeIdx := Queue.Dequeue()
		currentNode, _ := g.getNode(nodeIdx.(int))

		if currentNode.Data == dst.Data {
			break
		}

		visited.add(currentNode)

		for _, childIdx := range currentNode.Connection {
			child, _ := g.getNode(childIdx.Data)
			if !visited.contains(child) && !Queue.Contains(child) {
				Queue.Enqueue(child.Data)
				child.Parent = currentNode
			}
		}
	}

	paths := list.Queue{}
	path := dst
	for {
		paths.Enqueue(path)
		if path.Parent == nil {
			break
		}
		path = path.Parent
	}

	for _, p := range paths {
		fmt.Println(p.(*Node).Data)
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
