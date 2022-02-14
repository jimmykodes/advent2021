package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	b, err := os.ReadFile("input.txt")
	checkErr(err)
	rows := strings.Split(string(b), "\n")
	pattern := rows[0]
	data := rows[2:]

	pairs := make(map[string]string)
	for _, datum := range data {
		s := strings.Split(datum, "->")
		pair, insertion := s[0], s[1]
		pair = strings.TrimSpace(pair)
		insertion = strings.TrimSpace(insertion) + string(pair[1])

		pairs[pair] = insertion
	}

	steps := 40
	current := pattern
	for i := 0; i < steps; i++ {
		current = makeNew(current, pairs, i)
	}
	dat := make(map[rune]int)
	for _, s := range current {
		dat[s]++
	}
	greatest := 0
	least := math.MaxInt64
	for _, c := range dat {
		if c > greatest {
			greatest = c
		}
		if c < least {
			least = c
		}
	}
	fmt.Println(greatest - least)
}

func makeNew(current string, pairs map[string]string, i int) string {
	start := time.Now()
	newPat := string(current[0])
	for j := 1; j < len(current); j++ {
		p := string(current[j-1]) + string(current[j])
		newPat += pairs[p]
	}
	fmt.Println("finished step", i, "duration:", time.Since(start))
	return newPat
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
