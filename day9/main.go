package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(p2("input.txt"))
}

func p1(filename string) int {
	d := getData(filename)
	var risk int
	lp := getLowPoints(d)
	for _, p := range lp {
		risk += p.val + 1
	}
	return risk
}

func p2(filename string) int {
	d := getData(filename)
	lowPoints := getLowPoints(d)
	basins := make([]int, len(lowPoints))
	for i, p := range lowPoints {
		stack := []*point{p}
		var c *point
		for {
			if len(stack) == 0 {
				break
			} else if len(stack) == 1 {
				c, stack = stack[0], []*point{}
			} else {
				c, stack = stack[0], stack[1:]
			}

			// if we've already visited this point, continue
			if d[c.y][c.x].visited {
				continue
			}
			// make the current point visited
			d[c.y][c.x].visited = true
			// increment the basin
			basins[i]++

			lp, rp, tp, bp := getNeighbors(d, c.x, c.y)
			if lp != nil && !lp.visited && lp.val < 9 {
				stack = append(stack, lp)
			}
			if rp != nil && !rp.visited && rp.val < 9 {
				stack = append(stack, rp)
			}
			if tp != nil && !tp.visited && tp.val < 9 {
				stack = append(stack, tp)
			}
			if bp != nil && !bp.visited && bp.val < 9 {
				stack = append(stack, bp)
			}
		}

	}
	sort.Ints(basins)
	// initialize at 1, since we are multiplying, not adding.
	total := 1
	for i := 1; i <= 3; i++ {
		total *= basins[len(basins)-i]
	}
	return total
}

type point struct {
	x, y    int
	val     int
	visited bool
}

func getLowPoints(data [][]*point) []*point {
	var lowPoints []*point
	for _, row := range data {
		for _, p := range row {
			var (
				left, right, top, bottom bool
			)
			lp, rp, tp, bp := getNeighbors(data, p.x, p.y)

			left = lp != nil && lp.val <= p.val
			right = rp != nil && rp.val <= p.val
			top = tp != nil && tp.val <= p.val
			bottom = bp != nil && bp.val <= p.val
			if !(left || right || top || bottom) {
				// nothing is lower, increment risk
				lowPoints = append(lowPoints, p)
			}
		}
	}
	return lowPoints
}

func getNeighbors(data [][]*point, x, y int) (left, right, top, bottom *point) {
	if x > 0 {
		// not first column, look left
		left = data[y][x-1]
	}
	if x < len(data[0])-1 {
		// not last column, look right
		right = data[y][x+1]
	}
	if y > 0 {
		// not first row, look up
		top = data[y-1][x]
	}
	if y < len(data)-1 {
		// not last row, look down
		bottom = data[y+1][x]
	}
	return
}

func getData(filename string) [][]*point {
	d, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(d, []byte("\n"))
	data := make([][]*point, len(rows))
	for y, row := range rows {
		row = bytes.TrimSpace(row)
		for x, b := range row {
			parsed, err := strconv.Atoi(string(b))
			if err != nil {
				panic(err)
			}
			data[y] = append(data[y], &point{
				x:   x,
				y:   y,
				val: parsed,
			})
		}
	}
	return data
}

func makeImage(data [][]int) {
	img := image.NewGray(image.Rect(0, 0, len(data[0]), len(data)))
	for y, row := range data {
		for x, col := range row {
			img.Set(x, y, color.Gray{Y: uint8(25 * col)})
		}
	}
	imgFile, err := os.Create("img.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(imgFile, img)
	if err != nil {
		panic(err)
	}
}
