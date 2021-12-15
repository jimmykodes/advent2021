package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

func main() {
	fmt.Println(p1("input.txt"))
}

func p1(filename string) int {
	d := getData(filename)
	var risk int
	for y, row := range d {
		for x, col := range row {
			var (
				left, right, top, bottom bool
			)
			if x > 0 {
				// not first column, look left
				left = d[y][x-1] <= col
			}
			if x < len(d[0])-1 {
				// not last column, look right
				right = d[y][x+1] <= col
			}
			if y > 0 {
				// not first row, look up
				top = d[y-1][x] <= col
			}
			if y < len(d)-1 {
				// not last row, look down
				bottom = d[y+1][x] <= col
			}
			if !(left || right || top || bottom) {
				// nothing is lower, increment risk
				risk += col + 1
			}
		}
	}
	return risk
}

func getData(filename string) [][]int {
	d, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(d, []byte("\n"))
	data := make([][]int, len(rows))
	for i, row := range rows {
		row = bytes.TrimSpace(row)
		for _, b := range row {
			parsed, err := strconv.Atoi(string(b))
			if err != nil {
				panic(err)
			}
			data[i] = append(data[i], parsed)
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
