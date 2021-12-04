package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var numBits = 12

func main() {
	rows := getRows()
	o := oxygen(rows)
	c := carbon(rows)
	fmt.Println(o)
	fmt.Println(c)
	fmt.Println(c * o)
}

func mcb(rows []int64, bp int) bool {
	sums := make([]float64, numBits)
	for _, row := range rows {
		for i := 0; i < numBits; i++ {
			sums[i] += float64((row & (1 << i)) >> i)
		}
	}
	return sums[bp] >= float64(len(rows))/2.0
}

func getRows() []int64 {
	data, err := os.ReadFile("./real.txt")
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(bytes.Trim(data, "\n"), []byte("\n"))
	conv := make([]int64, len(rows))
	for i, row := range rows {
		num, err := strconv.ParseInt(string(row), 2, 64)
		if err != nil {
			panic(err)
		}
		conv[i] = num
	}
	return conv
}

func match(num int64, bp, val int) bool {
	return (num&(1<<bp))>>bp == int64(val)
}

func oxygen(rows []int64) int64 {
	r2 := make([]int64, len(rows))
	for i, r := range rows {
		r2[i] = r
	}
	for bp := numBits - 1; bp >= 0; bp-- {
		next := make([]int64, 0)
		if mcb(r2, bp) {
			// most common bit is a 1
			for _, r := range r2 {
				if match(r, bp, 1) {
					next = append(next, r)
				}
			}
		} else {
			for _, r := range r2 {
				if match(r, bp, 0) {
					next = append(next, r)
				}
			}
		}
		if len(next) == 1 {
			return next[0]
		}
		r2 = next
	}
	panic("how did i get here")
}

func carbon(rows []int64) int64 {
	r2 := make([]int64, len(rows))
	for i, r := range rows {
		r2[i] = r
	}
	for bp := numBits - 1; bp >= 0; bp-- {
		next := make([]int64, 0)
		if mcb(r2, bp) {
			// most common bit is a 1
			for _, r := range r2 {
				if match(r, bp, 0) {
					next = append(next, r)
				}
			}
		} else {
			for _, r := range r2 {
				if match(r, bp, 1) {
					next = append(next, r)
				}
			}
		}
		if len(next) == 1 {
			return next[0]
		}
		r2 = next
	}
	panic("how did i get here")
}

func printBin(rows []int64) {
	fmt.Print("[ ")
	for _, row := range rows {
		fmt.Printf("%05b ", row)
	}
	fmt.Println("]")
}
