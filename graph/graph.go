package graph

import (
	"ds/list"
	"errors"
	"fmt"
)

// Graph represent the whole node interconnections
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

// DepthFirstSearch will search a path with child end priority
func (g *Graph) DepthFirstSearch(source int, destination int) {
	src, _ := g.getNode(source)
	dst, _ := g.getNode(destination)
	visited := &Graph{}
	Path := &list.Stack{}

	// recursively search within a complete branch then moves on to next branch
	g.hasDfs(src, dst, visited, Path)

	// print Stack
	for _, p := range *Path {
		fmt.Println(p.(int))
	}
}

func (g *Graph) hasDfs(source *Node, destination *Node, visited *Graph, path *list.Stack) bool {
	// if node already visited, terminate branch
	if visited.contains(source) {
		path.Pop()
		return false
	}

	// mark currently visited node
	path.Push(source.Data)
	visited.add(source)

	// terminate if currently visited node is destination
	if source.Data == destination.Data {
		return true
	}

	// recursively call every connection in current node
	for _, child := range source.Connection {
		if g.hasDfs(child, destination, visited, path) {
			return true
		}
	}

	// if end of the line a.k.a no more path can be generated, terminate the function
	path.Pop()
	return false
}

// BreadthFirstSearch will search a path within every level
func (g *Graph) BreadthFirstSearch(source int, destination int) bool {
	src, _ := g.getNode(source)
	dst, _ := g.getNode(destination)
	Queue := &list.Queue{}
	visited := &Graph{}

	// insert source node to queue
	Queue.Enqueue(src.Data)

	// terminate loop if queue is empty
	for !Queue.IsEmpty() {
		// get currently visited node
		nodeIdx := Queue.Dequeue()
		currentNode, _ := g.getNode(nodeIdx.(int))

		// if current visited node is destination, terminate function
		if currentNode.Data == dst.Data {
			break
		}

		// mark current node as visited
		visited.add(currentNode)

		// for every connection in currently visited node, add to queue if not yet visited
		for _, childIdx := range currentNode.Connection {
			child, _ := g.getNode(childIdx.Data)
			if !visited.contains(child) && !Queue.Contains(child) {
				Queue.Enqueue(child.Data)
				child.Parent = currentNode
			}
		}
	}

	// reconstruct path to get correct order
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

// Node represent an element of graph
type Node struct {
	Data       int
	Connection Graph
	Parent     *Node
}

// Edge represent a connection between a node
type Edge struct {
	X int
	Y int
}

// Connect two node
func (n *Node) Connect(node *Node) {
	if _, ok := n.Connection[n.Data]; !ok {
		n.Connection[node.Data] = node
	}
}

// New returns a new instance of graph
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
