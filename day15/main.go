package main

import (
	"fmt"
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
	return s[i].Cost() > s[j].Cost()
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
	data := expandData(loadFile("input.txt"), 5)
	board := loadBoard(data)
	stack := Stack([]*Node{board[0][0]})
	step := 0
	current := stack.Pop()
	current.Visited = true
	for !current.IsEnd {
		step++
		via := current.Via
		for _, n := range []*Node{current.Left, current.Right, current.Top, current.Bottom} {
			if n == nil || n == via {
				continue
			}
			totalRisk := current.TotalRisk + n.Risk
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
	// for _, nodes := range board {
	// 	for _, node := range nodes {
	// 		char := "-"
	// 		if node.Path {
	// 			char = strconv.Itoa(node.Risk)
	// 		}
	// 		fmt.Print(char)
	// 	}
	// 	fmt.Println()
	// }
}

func loadBoard(rows [][]int) [][]*Node {
	b := make([][]*Node, len(rows))
	var finalX, finalY int
	finalY = len(rows) - 1
	for y, row := range rows {
		b[y] = make([]*Node, len(row))
		finalX = len(row) - 1
		for x, r := range row {
			n := &Node{
				X:    x,
				Y:    y,
				Risk: r,
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

func loadFile(filename string) [][]int {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(data), "\n")
	b := make([][]int, len(rows))
	for y, row := range rows {
		b[y] = make([]int, len(row))
		for x, char := range row {
			r, err := strconv.Atoi(string(char))
			if err != nil {
				panic(err)
			}
			b[y][x] = r
		}
	}
	return b
}

func expandData(data [][]int, expansionFactor int) [][]int {
	rows := len(data)
	cols := len(data[0])
	d := make(board, rows*expansionFactor)

	for i := 0; i < expansionFactor; i++ {
		for j := 0; j < expansionFactor; j++ {
			for r := 0; r < rows; r++ {
				if d[r+(i*rows)] == nil {
					d[r+(i*rows)] = make([]int, cols*expansionFactor)
				}
				for c := 0; c < cols; c++ {
					v := data[r][c] + i + j
					v = (v % 10) + (v / 10)
					d[r+(i*rows)][c+(j*cols)] = v
				}
			}
		}
	}
	return d
}

type Node struct {
	X         int
	Y         int
	TotalRisk int
	Risk      int

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

func (n Node) Cost() int {
	return n.TotalRisk
}

type board [][]int

func (b board) String() string {
	var sb strings.Builder
	for _, row := range b {
		for _, col := range row {
			fmt.Fprint(&sb, col)
		}
		fmt.Fprintln(&sb)
	}
	return sb.String()
}
