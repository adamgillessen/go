// Package graph contains functions that allow you to represent
// graphs and run common algorithms on them.
//
// To reference a node in the graph, you should wrap the identifier
// for that node with the Node type.
//
// Example:
//
// g.AddNode(graph.Node(1))
package graph

import "fmt"

// Node represents a unique node in a graph
type Node interface{}

// Graph represents a weighted directed graph
type Graph struct {
	graph map[Node]map[Node]int
}

// New returns a weighted, directed graph with
// no Nodes or edges.
func New() *Graph {
	return &Graph{
		graph: make(map[Node]map[Node]int),
	}
}

// AddNode adds a Node to the graph. An error is returned if a Node
// with the given identifier already exists.
func (g *Graph) AddNode(x Node) error {
	if _, ok := g.graph[x]; ok {
		return fmt.Errorf("node %v already exists", x)
	}
	g.graph[x] = make(map[Node]int)
	return nil
}

// AddEdge adds a directed edge between u and v for a given cost
func (g *Graph) AddEdge(u, v Node, cost int) error {
	if _, ok := g.graph[u]; !ok {
		return fmt.Errorf("node %v doesn't exist", u)
	}
	if _, ok := g.graph[v]; !ok {
		return fmt.Errorf("node %v doesn't exist", v)
	}
	g.graph[u][v] = cost
	return nil
}

// Neighbours returns the nodes which share an edge with u
func (g *Graph) Neighbours(u Node) ([]Node, error) {
	neghbours, ok := g.graph[u]
	if !ok {
		return nil, fmt.Errorf("node %v doesn't exist", u)
	}
	var vs []Node
	for v := range neghbours {
		vs = append(vs, v)
	}
	return vs, nil
}

// Nodes returns all of the nodes in the graph
func (g *Graph) Nodes() []Node {
	var nodes []Node
	for n := range g.graph {
		nodes = append(nodes, n)
	}
	return nodes
}

// BFS runs the breadth first search algorithm and returns a map
// of the *length of the paths* from the given source to that
// destination. The cost is -1 if there is no path to that destination.
func (g *Graph) BFS(src Node) (map[Node]int, error) {
	costs := make(map[Node]int)
	for _, n := range g.Nodes() {
		costs[n] = -1
	}
	if _, ok := g.graph[src]; !ok {
		return nil, fmt.Errorf("source node %v not in graph", src)
	}
	costs[src] = 0

	level := []Node{src}
	for jumps := 1; len(level) > 0; jumps++ {
		var nextLevel []Node
		for _, u := range level {
			vs, _ := g.Neighbours(u)
			for _, v := range vs {
				if costs[v] != -1 {
					continue
				}
				costs[v] = jumps
				nextLevel = append(nextLevel, v)
			}
		}
		level = nextLevel
	}
	return costs, nil
}
