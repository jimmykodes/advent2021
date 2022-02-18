package main

import (
	"fmt"
	"os"
	"strings"
)

type image [][]int

func (img image) String() string {
	var sb strings.Builder
	for _, row := range img {
		for _, col := range row {
			sb.WriteString(fmt.Sprintf("%d", col))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func main() {
	img, enhancementStr := readImage("input.txt")
	for i := 0; i < 50; i++ {
		var filler int
		if enhancementStr[0] == '#' && enhancementStr[len(enhancementStr)-1] == '.' {
			filler = i % 2
		}
		img = enhanceImage(img, enhancementStr, filler)
	}
	count := 0
	for _, row := range img {
		for _, val := range row {
			count += val
		}
	}
	fmt.Println(count)
}

func enhanceImage(img image, str string, filler int) image {
	img = padImage(img, filler)
	out := make(image, len(img))
	for y := 0; y < len(img); y++ {
		out[y] = make([]int, len(img[0]))
		for x := 0; x < len(img[y]); x++ {
			var index int
			for l, y1 := range []int{-1, 0, 1} {
				for k, x1 := range []int{-1, 0, 1} {
					p := getPoint(x+x1, y+y1, img, filler)
					index |= p << (8 - ((l * 3) + k))
				}
			}
			col := str[index]
			v := 0
			if col == '#' {
				v = 1
			}
			out[y][x] = v
		}
	}
	return out
}

func padImage(i image, filler int) image {
	for y, row := range i {
		i[y] = append([]int{filler}, append(row, filler)...)
	}
	rowLen := len(i[0])
	return append([][]int{makeRow(rowLen, filler)}, append(i, makeRow(rowLen, filler))...)
}

func makeRow(length, filler int) []int {
	row := make([]int, length)
	if filler == 0 {
		return row
	}
	for i := 0; i < length; i++ {
		row[i] = filler
	}
	return row
}

func getPoint(x, y int, i image, filler int) int {
	if x < 0 || y < 0 {
		// above/left of image
		return filler
	}
	if x >= len(i[0]) || y >= len(i) {
		// below/right of image
		return filler
	}
	return i[y][x]
}

func readImage(filename string) (image, string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(data), "\n\n")
	enhancementStr := rows[0]
	rows = strings.Split(rows[1], "\n")
	out := make([][]int, len(rows))
	for y := 0; y < len(rows); y++ {
		out[y] = make([]int, len(rows[y]))
		for x := 0; x < len(rows[y]); x++ {
			var v int
			if rows[y][x] == '#' {
				v = 1
			}
			out[y][x] = v
		}
	}
	return out, enhancementStr
}
