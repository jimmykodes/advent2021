package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numRegex = regexp.MustCompile("[0-9]+")

type tile struct {
	val    int
	called bool
}

func (t tile) String() string {
	if t.called {
		return fmt.Sprintf("(%2d)", t.val)
	}
	return fmt.Sprintf(" %2d ", t.val)
}

type board struct {
	tiles      [][]*tile
	lastCalled int
}

func (b board) String() string {
	var sb strings.Builder
	for _, tiles := range b.tiles {
		for _, t := range tiles {
			sb.WriteString(t.String() + " ")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (b board) score() int {
	var sum int
	for _, tiles := range b.tiles {
		for _, t := range tiles {
			if !t.called {
				sum += t.val
			}
		}
	}
	return sum * b.lastCalled
}

func (b *board) update(drawn int) {
TileLoop:
	for _, tiles := range b.tiles {
		for _, t := range tiles {
			if t.val == drawn {
				t.called = true
				b.lastCalled = drawn
				// we'll only have one instance of a number per board, so exit once matched
				break TileLoop
			}
		}
	}
}

func (b board) won() bool {
	cols := []bool{true, true, true, true, true}
	for _, tiles := range b.tiles {
		row := true
		for i, t := range tiles {
			if !t.called {
				cols[i] = false
				row = false
			}
		}
		if row {
			return true
		}
	}
	for _, col := range cols {
		if col {
			return true
		}
	}
	return false
}

type data struct {
	round  int
	draws  []int
	boards []*board
}

// DoRound will update all boards with the number at draws[round] and then increment round
// returns the first board to have a bingo, or nil of no board wins
func (d *data) DoRound() ([]int, []*board) {
	if d.round > len(d.draws) {
		panic("exceeded rounds")
	}
	drawn := d.draws[d.round]
	d.round++
	var (
		winners       []*board
		winnerIndexes []int
	)
	for i, b := range d.boards {
		b.update(drawn)
		if b.won() {
			winners = append(winners, b)
			winnerIndexes = append(winnerIndexes, i)
		}
	}
	return winnerIndexes, winners
}

func (d *data) removeBoard(at int) {
	d.boards = append(d.boards[:at], d.boards[at+1:]...)
}

func (d data) String() string {
	var sb strings.Builder
	_, _ = fmt.Fprintf(&sb, "%v\n", d.draws)
	for _, b := range d.boards {
		sb.WriteString(b.String() + "\n")
	}
	return sb.String()
}

func getData(filename string) *data {
	fData, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(fData), "\n")
	d := &data{}
	var b *board
	for i, row := range rows {
		if i == 0 {
			parts := strings.Split(row, ",")
			d.draws = make([]int, len(parts))
			for i2, part := range parts {
				d.draws[i2], err = strconv.Atoi(part)
				if err != nil {
					panic(err)
				}
			}
			continue
		}
		if row == "" {
			// empty line, start a new board
			if b != nil {
				d.boards = append(d.boards, b)
			}
			b = &board{}
			continue
		}
		rawNums := numRegex.FindAllString(row, -1)
		tiles := make([]*tile, len(rawNums))
		for i2, num := range rawNums {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			tiles[i2] = &tile{val: n}
		}
		b.tiles = append(b.tiles, tiles)
	}
	return d
}
