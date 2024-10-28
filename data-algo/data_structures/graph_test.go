package data_structures_test

import (
	"testing"

	ds "github.com/flan6/data-algo/data_structures"
)

func TestGraph(t *testing.T) {
	// setup
	graph := ds.Graph[string]{}

	bh := graph.AddVertex("bh")
	rio := graph.AddVertex("rio")
	betim := graph.AddVertex("betim")
	sp := graph.AddVertex("sp")

	graph.AddEdge(bh, rio, 498)
	graph.AddEdge(bh, betim, 35)
	graph.AddEdge(rio, sp, 529)
	graph.AddEdge(betim, sp, 647)

	solution := []*ds.Vertex[string]{bh, rio, sp}

	path := graph.DFS(bh, sp)
	if len(path) == 0 {
		t.Fatal("got empty path")
	}

	for i := range solution {
		if solution[i] != path[i] {
			t.Fatalf("wrong path, got: %v expected: %v", path, solution)
		}
	}
}
