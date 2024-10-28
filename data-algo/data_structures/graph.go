package data_structures

import "slices"

type Graph[T comparable] struct {
	v []*Vertex[T]
}

type Vertex[T comparable] struct {
	value T
	adj   []Edge[T]
}

type Edge[T comparable] struct {
	to     *Vertex[T]
	weight int
}

func (e *Vertex[T]) empty() bool {
	return e == nil ||
		(e.value == *new(T) && e.adj == nil)
}

func (g *Graph[T]) AddVertex(value T) *Vertex[T] {
	vertex := &Vertex[T]{value: value}
	g.v = append(g.v, vertex)
	return vertex
}

func (g *Graph[T]) AddEdge(from, to *Vertex[T], weight int) {
	from.adj = append(from.adj, Edge[T]{to: to, weight: weight})
}

func (g *Graph[T]) vertexIdx(vertex *Vertex[T]) int {
	for i := range g.v {
		if g.v[i] == vertex {
			return i
		}
	}

	return -1
}

// DFS performs a depth-first search from the start vertex and returns a path to the target vertex
func (g *Graph[T]) DFS(start, target *Vertex[T]) []*Vertex[T] {
	path := Stack[*Vertex[T]]{}
	seen := make([]bool, len(g.v))

	g.walk(g.vertexIdx(start), target, seen, &path)

	res := make([]*Vertex[T], 0)
	for {
		Vertex := path.Pop()
		if Vertex.empty() {
			break
		}

		res = append(res, Vertex)
	}

	slices.Reverse(res)

	return res
}

func (g *Graph[T]) walk(current int, target *Vertex[T], seen []bool, path *Stack[*Vertex[T]]) bool {
	// base
	if seen[current] {
		return false
	}

	if g.v[current].value == target.value {
		path.Push(g.v[current])
		return true
	}

	// pre
	seen[current] = true

	path.Push(g.v[current])

	// recursion
	for _, edge := range g.v[current].adj {
		if g.walk(g.vertexIdx(edge.to), target, seen, path) {
			return true
		}
	}

	// post
	path.Pop()

	return false
}
