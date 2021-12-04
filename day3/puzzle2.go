package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var numBits = 12

func p2() {
	rows := getRows()
	o := oxygen(rows)
	c := carbon(rows)
	fmt.Println(o)
	fmt.Println(c)
	fmt.Println(c * o)
}

func mcb(rows []int64, bp int) bool {
	var sum int64
	for _, row := range rows {
		sum += (row & (1 << bp)) >> bp
	}
	return sum*2 >= int64(len(rows))
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
	return filter(rows, true)
}

func carbon(rows []int64) int64 {
	return filter(rows, false)
}

func filter(rows []int64, useMCB bool) int64 {
	r2 := clone(rows)
	for bp := numBits - 1; bp >= 0; bp-- {
		next := make([]int64, 0)
		var matchVal int
		mostCommon := mcb(r2, bp)
		switch {
		case (mostCommon && useMCB) || (!mostCommon && !useMCB):
			matchVal = 1
		case (mostCommon && !useMCB) || (!mostCommon && useMCB):
			matchVal = 0
		}
		for _, r := range r2 {
			if match(r, bp, matchVal) {
				next = append(next, r)
			}
		}
		r2 = next
		if len(next) == 1 {
			break
		}
	}
	return r2[0]
}

func clone(rows []int64) []int64 {
	r2 := make([]int64, len(rows))
	for i, r := range rows {
		r2[i] = r
	}
	return r2
}

func printBin(rows []int64) {
	fmt.Print("[ ")
	for _, row := range rows {
		fmt.Printf("%05b ", row)
	}
	fmt.Println("]")
}
