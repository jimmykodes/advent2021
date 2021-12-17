package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println(p2("input.txt"))
}

func p1(filename string) int {
	data := getData(filename)
	var score int
RowLoop:
	for _, row := range data {
		var stack []byte
		for _, b := range row {
			if isOpen(b) {
				stack = append(stack, b)
			} else {
				if correctCloser(stack[len(stack)-1], b) {
					stack = stack[:len(stack)-1]
				} else {
					switch rune(b) {
					case '}':
						score += 1197
					case ')':
						score += 3
					case ']':
						score += 57
					case '>':
						score += 25137
					}
					continue RowLoop
				}
			}
		}
	}
	return score
}

func p2(filename string) int {
	data := getData(filename)
	var scores []int
RowLoop:
	for _, row := range data {
		var stack []byte
		for _, b := range row {
			if isOpen(b) {
				stack = append(stack, b)
			} else {
				if correctCloser(stack[len(stack)-1], b) {
					stack = stack[:len(stack)-1]
				} else {
					// invalid, ignore row
					continue RowLoop
				}
			}
		}
		var localScore int
		for i := len(stack) - 1; i >= 0; i-- {
			localScore *= 5
			switch rune(stack[i]) {
			case '{':
				localScore += 3
			case '(':
				localScore += 1
			case '[':
				localScore += 2
			case '<':
				localScore += 4
			}
		}
		scores = append(scores, localScore)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func getData(filename string) [][]byte {
	d, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes.Split(d, []byte("\n"))
}

func correctCloser(open, close byte) bool {
	switch rune(open) {
	case '{':
		return rune(close) == '}'
	case '(':
		return rune(close) == ')'
	case '[':
		return rune(close) == ']'
	case '<':
		return rune(close) == '>'
	default:
		return false
	}
}

func isOpen(b byte) bool {
	switch rune(b) {
	case '{', '(', '[', '<':
		return true
	default:
		return false
	}
}
