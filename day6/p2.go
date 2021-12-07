package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(string(bytes.TrimSpace(data)), -1)
	fish := make([]int, len(nums))
	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		fish[i] = n
	}
	var total int
	c := NewCounter()
	for _, f := range fish {
		total += c.count(f, 256)
	}
	fmt.Println(total)
}

type Counter struct {
	memo map[string]int
}

func NewCounter() *Counter {
	return &Counter{memo: make(map[string]int)}
}

func (c *Counter) count(x, days int) int {
	key := fmt.Sprintf("%d_%d", x, days)
	if v, ok := c.memo[key]; ok {
		return v
	}
	if days == 0 {
		return 1
	}
	var v int
	if x == 0 {
		v = c.count(6, days-1) + c.count(8, days-1)
	} else {
		v = c.count(x-1, days-1)
	}
	c.memo[key] = v
	return v
}
