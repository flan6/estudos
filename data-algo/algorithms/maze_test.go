package algorithms

import "testing"

func TestSolveMaze(t *testing.T) {
	maze := []string{
		"XXXXXXXX X",
		"X      X X",
		"X      X X",
		"X XXXXXX X",
		"X        X",
		"X XXXXXXXX",
	}

	start := Point{8, 0}
	end := Point{1, 5}
	solution := []Point{
		start,
		{8, 1},
		{8, 2},
		{8, 3},
		{8, 4},
		{7, 4},
		{6, 4},
		{5, 4},
		{4, 4},
		{3, 4},
		{2, 4},
		{1, 4},
		end,
	}

	res := SolveMaze(maze, rune('X'), start, end)
	if len(res) == 0 {
		t.Fatal("got empty result")
	}

	if res[0] != start {
		t.Errorf("started from wrong position, got: %v expected: %v", res[0], start)
	}

	if res[len(res)-1] != end {
		t.Errorf("wrong end, got: %v expected: %v", res[len(res)-1], end)
	}

	for i := range len(res) - 1 {
		if res[i] != solution[i] {
			t.Fatalf("wrong solution, got: %v expected: %v", res, solution)
		}
	}
}
