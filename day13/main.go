package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	coordRegex = regexp.MustCompile("[0-9]+,[0-9]+")
	foldRegex  = regexp.MustCompile("fold along .+")
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	data := string(b)

	rows := strings.Split(data, "\n")

	var coords []*Coord
	var folds []*Fold
	for _, row := range rows {
		if coordRegex.MatchString(row) {
			c := strings.Split(row, ",")
			coord := &Coord{
				X: mustInt(c[0]),
				Y: mustInt(c[1]),
			}
			coords = append(coords, coord)
		} else if foldRegex.MatchString(row) {
			s := strings.TrimPrefix(row, "fold along ")
			f := strings.Split(s, "=")
			var dir int
			if f[0] == "y" {
				dir = up
			} else {
				dir = left
			}
			folds = append(folds, &Fold{
				Dir:      dir,
				Location: mustInt(f[1]),
			})
		}
	}

	for _, fold := range folds {
		for _, coord := range coords {
			coord.Fold(fold)
		}
	}

	maxY, maxX := 0, 0
	for _, coord := range coords {
		if coord.X > maxX {
			maxX = coord.X
		}
		if coord.Y > maxY {
			maxY = coord.Y
		}
	}

	img := image.NewGray(image.Rect(0, 0, maxX+1, maxY+1))
	for _, coord := range coords {
		img.Set(coord.X, coord.Y, color.White)
	}
	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	png.Encode(f, img)
}

type Coord struct {
	X int
	Y int
}

func (c *Coord) Fold(f *Fold) {
	switch f.Dir {
	case up:
		if c.Y < f.Location {
			return
		}
		c.Y = f.Location - (c.Y - f.Location)
	case left:
		if c.X < f.Location {
			return
		}
		c.X = f.Location - (c.X - f.Location)
	}
}

const (
	up = iota
	left
)

type Fold struct {
	Dir      int
	Location int
}

func mustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
