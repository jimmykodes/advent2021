package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Stack []*Node

func (s Stack) Len() int {
	return len(s)
}

func (s Stack) Less(i, j int) bool {
	return s[i].TotalRisk > s[j].TotalRisk
}

func (s Stack) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *Stack) Push(x *Node) {
	*s = append(*s, x)
	sort.Sort(*s)
}

func (s *Stack) Pop() *Node {
	o := *s
	n := len(o) - 1
	node := o[n]
	o[n] = nil
	*s = o[0:n]
	return node
}

func (s *Stack) Peek() *Node {
	o := *s
	return o[len(*s)-1]
}

func main() {
	board := loadBoard("input.txt")
	stack := Stack([]*Node{board[0][0]})
	step := 0
	current := stack.Pop()
	current.Visited = true
	for !current.IsEnd {
		step++
		via := current.Via
		previousRisk := 0
		if via != nil {
			previousRisk = via.TotalRisk
		}
		for _, n := range []*Node{current.Left, current.Right, current.Top, current.Bottom} {
			if n == nil || n == via {
				continue
			}
			totalRisk := previousRisk + current.Risk + n.Risk
			if n.Visited {
				if totalRisk < n.TotalRisk {
					// less risky through this path, update things
					n.Via = current
					n.TotalRisk = totalRisk
					sort.Sort(stack)
				}
				continue
			}
			n.Via = current
			n.TotalRisk = totalRisk
			n.Visited = true
			stack.Push(n)
		}
		current = stack.Pop()
	}
	fmt.Println("completed in steps:", step)
	r := 0
	for !current.IsStart {
		current.Path = true
		r += current.Risk
		current = current.Via
	}
	fmt.Println("total risk:", r)
	for _, nodes := range board {
		for _, node := range nodes {
			char := "-"
			if node.Path {
				char = strconv.Itoa(node.Risk)
			}
			fmt.Print(char)
		}
		fmt.Println()
	}
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
			if y == 0 && x == 0 {
				n.IsStart = true
			}
			if y == finalY && x == finalX {
				n.IsEnd = true
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
	TotalRisk int
	Risk      int
	SquareMag float64

	Left   *Node
	Right  *Node
	Top    *Node
	Bottom *Node
	Via    *Node

	Visited bool
	IsStart bool
	IsEnd   bool
	Path    bool
}

func (n Node) String() string {
	return fmt.Sprintf("(%d, %d) %d %d", n.X, n.Y, n.Risk, n.TotalRisk)
}
