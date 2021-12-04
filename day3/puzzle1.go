package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func p1() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	numBits := 12
	sums := make([]int64, numBits)
	rows := bytes.Split(bytes.Trim(data, "\n"), []byte("\n"))
	for _, row := range rows {
		num, err := strconv.ParseInt(string(row), 2, 64)
		if err != nil {
			panic(err)
		}
		for i := 0; i < numBits; i++ {
			sums[i] += (num & (1 << i)) >> i
		}
	}

	fmt.Println(sums)
	var (
		gamma   int
		epsilon int
	)
	half := int64(len(rows) / 2)
	for i, s := range sums {
		if s > half {
			gamma |= 1 << i
		}
	}

	epsilon = gamma ^ ((1 << numBits) - 1)
	fmt.Println(gamma)
	fmt.Println(epsilon)
	fmt.Println(gamma * epsilon)
}
