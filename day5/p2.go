package main

import (
	"fmt"
	"os"
	"strings"
)

func p2(filename string, dim int) {
	d, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var lines []*Line
	for _, s := range strings.Split(string(d), "\n") {
		nums := re.FindAllString(s, -1)
		if len(nums) == 0 {
			continue
		}
		lines = append(lines, &Line{
			Start: &Coord{X: MustConvert(nums[0]), Y: MustConvert(nums[1])},
			End:   &Coord{X: MustConvert(nums[2]), Y: MustConvert(nums[3])},
		})
	}
	// initialize board
	board := make([][]int, dim)
	for i := 0; i < dim; i++ {
		board[i] = make([]int, dim)
	}
	for _, line := range lines {
		x := line.Start.X
		y := line.Start.Y
		var dx func(int) int
		var dy func(int) int
		if line.Start.X < line.End.X {
			dx = incr
		} else if line.Start.X > line.End.X {
			dx = dec
		} else {
			// noop
			dx = func(i int) int { return i }
		}
		if line.Start.Y < line.End.Y {
			dy = incr
		} else if line.Start.Y > line.End.Y {
			dy = dec
		} else {
			dy = func(i int) int { return i }
		}
		for x != line.End.X || y != line.End.Y {
			board[y][x]++
			x = dx(x)
			y = dy(y)
		}
		board[y][x]++
	}

	var c int
	for _, row := range board {
		for _, col := range row {
			if col > 1 {
				c++
			}
		}
	}
	fmt.Println(c)
}
