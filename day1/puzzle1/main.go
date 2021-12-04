package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(data, []byte("\n"))
	var (
		last  int
		count int
	)
	for i, row := range rows {
		current, err := strconv.Atoi(string(row))
		if err != nil {
			panic(err)
		}
		if i == 0 {
			last = current
			continue
		}
		if current > last {
			count++
		}
		last = current
	}
	fmt.Println(count)
}
