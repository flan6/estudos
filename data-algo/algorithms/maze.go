package algorithms

import (
	"slices"

	ds "github.com/flan6/data-algo/data_structures"
)

type Point struct {
	x int
	y int
}

func (p *Point) Empty() bool {
	return p == nil || *p == Point{}
}

var directions = [][]int{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func SolveMaze(maze []string, wall rune, start Point, end Point) []Point {
	path := ds.NewStack[Point]()
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[i]))
	}

	walk(maze, wall, start, end, seen, &path)

	res := make([]Point, 0)
	for {
		pt := path.Pop()
		if pt.Empty() {
			break
		}

		res = append(res, pt)
	}

	slices.Reverse(res)

	return res
}

func walk(maze []string, wall rune, curr, end Point, seen [][]bool, path *ds.Stack[Point]) bool {
	// base cases

	// 1. outside map
	if curr.x < 0 || curr.x >= len(maze[0]) ||
		curr.y < 0 || curr.y >= len(maze) {
		return false
	}

	// 2. is a wall
	if maze[curr.y][curr.x] == byte(wall) {
		return false
	}

	// 3. is our end goal
	if curr.x == end.x &&
		curr.y == end.y {
		path.Push(end)
		return true
	}

	// 4. already visited
	if seen[curr.y][curr.x] {
		return false
	}

	seen[curr.y][curr.x] = true
	path.Push(curr)

	// recursion
	for _, dir := range directions {
		x, y := dir[0], dir[1]
		if walk(maze, wall, Point{x: curr.x + x, y: curr.y + y}, end, seen, path) {
			return true
		}
	}

	// post
	path.Pop()

	return false
}
