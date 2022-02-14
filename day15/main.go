package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	board := loadBoard("test.txt")
	fmt.Println(board)
}

func loadBoard(filename string) [][]*Node {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(data), "\n")
	b := make([][]*Node, len(rows))
	var finalX, finalY int
	finalY = len(rows) - 1
	for y, row := range rows {
		b[y] = make([]*Node, len(row))
		finalX = len(row) - 1
		for x, char := range row {
			r, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			n := &Node{
				X:         x,
				Y:         y,
				Risk:      r,
				SquareMag: math.Pow(float64(finalX-x), 2) + math.Pow(float64(finalY-y), 2),
			}
			b[y][x] = n
			if x > 0 {
				// connect left
				n.Left = b[y][x-1]
				b[y][x-1].Right = n
			}
			if y > 0 {
				// connect up
				n.Top = b[y-1][x]
				b[y-1][x].Bottom = n
			}
		}
	}

	return b
}

type Node struct {
	X         int
	Y         int
	Risk      int
	SquareMag float64
	Left      *Node
	Right     *Node
	Top       *Node
	Bottom    *Node
	Via       *Node
}
