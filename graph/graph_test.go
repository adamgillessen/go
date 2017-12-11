package graph

import (
	"reflect"
	"testing"
)

func TestAddNode(t *testing.T) {
	want := &Graph{
		graph: map[Node]map[Node]int{
			Node(420): make(map[Node]int),
		},
	}

	got := New()
	if err := got.AddNode(Node(420)); err != nil {
		t.Errorf("Failed to add edge: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddNode -> %v, want %v", got, want)
	}
}

func TestAddEdge(t *testing.T) {
	want := &Graph{
		graph: map[Node]map[Node]int{
			Node(420): map[Node]int{
				Node(421): 1,
			},
			Node(421): make(map[Node]int),
		},
	}

	got := New()
	got.AddNode(Node(420))
	got.AddNode(Node(421))
	if err := got.AddEdge(Node(420), Node(421), 1); err != nil {
		t.Errorf("Failed to add edge: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddEdge -> %v, want %v", got, want)
	}
}

func TestNeighbours(t *testing.T) {
	want := []Node{Node(421), Node(422)}

	g := &Graph{
		graph: map[Node]map[Node]int{
			Node(420): map[Node]int{
				Node(421): 1,
				Node(422): 1,
			},
			Node(421): map[Node]int{
				Node(420): 1,
				Node(422): 1,
			},
			Node(422): map[Node]int{
				Node(421): 1,
				Node(420): 1,
			},
		},
	}
	got, err := g.Neighbours(Node(420))
	if err != nil {
		t.Errorf("Failed to get neighbours: %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("Neighbours(Node(420)) = %v, want %v", got, want)
	}

	for i := 0; i < len(got); i++ {
		found := false
		for j := 0; j < len(want); j++ {
			if got[i] == want[j] {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Neighbours(Node(420)) = %v, want %v", got, want)
		}
	}
}

func TestNodes(t *testing.T) {
	want := []Node{Node(420), Node(421), Node(422)}

	g := &Graph{
		graph: map[Node]map[Node]int{
			Node(420): make(map[Node]int),
			Node(421): make(map[Node]int),
			Node(422): make(map[Node]int),
		},
	}
	got := g.Nodes()

	if len(got) != len(want) {
		t.Fatalf("Nodes() = %v, want %v", got, want)
	}

	for i := 0; i < len(got); i++ {
		found := false
		for j := 0; j < len(want); j++ {
			if got[i] == want[j] {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Nodes() = %v, want %v", got, want)
		}
	}
}

func TestBFS(t *testing.T) {
	want := map[Node]int{
		Node(0): 0,
		Node(1): 1,
		Node(2): 2,
		Node(3): 2,
		Node(4): -1,
	}

	g := New()
	for n := 0; n < 5; n++ {
		g.AddNode(Node(n))
	}

	g.AddEdge(Node(0), Node(1), 1)
	g.AddEdge(Node(1), Node(2), 1)
	g.AddEdge(Node(1), Node(3), 1)

	got, err := g.BFS(Node(0))
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("BFS() = %v, %v want %v, <nil>", got, err, want)
	}
}

func TestBFSCycleGraph(t *testing.T) {
	want := map[Node]int{
		Node(0): 0,
		Node(1): 1,
		Node(2): 1,
	}

	g := New()
	for n := 0; n < 3; n++ {
		g.AddNode(Node(n))
	}

	g.AddEdge(Node(0), Node(1), 1)
	g.AddEdge(Node(1), Node(0), 1)

	g.AddEdge(Node(1), Node(2), 1)
	g.AddEdge(Node(2), Node(1), 1)

	g.AddEdge(Node(2), Node(0), 1)
	g.AddEdge(Node(0), Node(2), 1)

	got, err := g.BFS(Node(0))
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Errorf("BFS() = %v, %v want %v, <nil>", got, err, want)
	}
}
