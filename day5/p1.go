package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

type Line struct {
	Start *Coord
	End   *Coord
}

func (l Line) IsHorizontal() bool {
	return l.Start.Y == l.End.Y
}

func (l Line) IsVertical() bool {
	return l.Start.X == l.End.X
}

func (l Line) String() string {
	return fmt.Sprintf("%s -> %s\n", l.Start.String(), l.End.String())
}

var re = regexp.MustCompile("[0-9]+")

func p1(filename string, dim int) {
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
		l := &Line{
			Start: &Coord{X: MustConvert(nums[0]), Y: MustConvert(nums[1])},
			End:   &Coord{X: MustConvert(nums[2]), Y: MustConvert(nums[3])},
		}
		if l.IsHorizontal() || l.IsVertical() {
			lines = append(lines, l)
		}
	}
	// initialize board
	board := make([][]int, dim)
	for i := 0; i < dim; i++ {
		board[i] = make([]int, dim)
	}
	for _, line := range lines {
		if line.IsHorizontal() {
			x := line.Start.X
			y := line.Start.Y
			var f func(int) int
			if line.Start.X < line.End.X {
				f = incr
			} else {
				f = dec
			}
			for x != line.End.X {
				board[y][x]++
				x = f(x)
			}
		} else {
			// is vertical
			x := line.Start.X
			y := line.Start.Y
			var f func(int) int
			if line.Start.Y < line.End.Y {
				f = incr
			} else {
				f = dec
			}
			for y != line.End.Y {
				board[y][x]++
				y = f(y)
			}
		}
		board[line.End.Y][line.End.X]++
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

func incr(n int) int {
	return n + 1
}

func dec(n int) int {
	return n - 1
}

func MustConvert(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
