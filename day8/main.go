package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println(p1("input.txt"))
}

// not sure if this is the right approach here. moving on to another puzzle... will return to it later
func p2(filename string) int {
	rows := getInput(filename)
	var sum int
	for _, row := range rows {
		conn := make([]map[rune]int, 7)
		for i := range conn {
			conn[i] = make(map[rune]int)
		}
		// indexes make a seven segment display like so:
		//  1111
		// 0    2
		// 0    2
		//  3333
		// 4    6
		// 4    6
		//  5555
		for _, input := range row.Input {
			switch len(input) {
			case 2:
				// this is a 1
				for _, b := range input {
					conn[2][rune(b)]++
					conn[6][rune(b)]++
				}
			case 3:
				// this is a 7
				for _, b := range input {
					conn[1][rune(b)]++
					conn[2][rune(b)]++
					conn[6][rune(b)]++
				}
			case 4:
				// this is a 4
				for _, b := range input {
					conn[0][rune(b)]++
					conn[2][rune(b)]++
					conn[3][rune(b)]++
					conn[6][rune(b)]++
				}
			case 7:
				// this is an 8
				for _, b := range input {
					conn[0][rune(b)]++
					conn[1][rune(b)]++
					conn[2][rune(b)]++
					conn[3][rune(b)]++
					conn[4][rune(b)]++
					conn[5][rune(b)]++
					conn[6][rune(b)]++
				}
			}
		}
	}
	return sum
}

func p1(filename string) int {
	var c int
	rows := getInput(filename)
	for _, row := range rows {
		for _, b := range row.Output {
			switch len(b) {
			case 2, 3, 4, 7:
				c++
			}
		}
	}
	return c
}

func getInput(filename string) []Data {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	dataRows := make([]Data, len(rows))
	for i, row := range rows {
		s := bytes.SplitN(row, []byte("|"), 2)
		dataRows[i] = Data{
			Input:  bytes.Split(bytes.TrimSpace(s[0]), []byte(" ")),
			Output: bytes.Split(bytes.TrimSpace(s[1]), []byte(" ")),
		}
	}
	return dataRows
}

type Data struct {
	Input  [][]byte
	Output [][]byte
}
