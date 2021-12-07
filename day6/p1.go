package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Fish struct {
	days int
}

func (f *Fish) Incr() bool {
	if f.days == 0 {
		f.days = 6
		return true
	}
	f.days--
	return false
}

func (f Fish) String() string {
	return fmt.Sprintf("%d", f.days)
}

func main() {
	// fish := []*Fish{
	// 	{days: 3},
	// 	{days: 4},
	// 	{days: 3},
	// 	{days: 1},
	// 	{days: 2},
	// }
	fish := getFish()
	for day := 0; day < 80; day++ {
		newFish := make([]*Fish, 0)
		for _, f := range fish {
			if makeNewFish := f.Incr(); makeNewFish {
				newFish = append(newFish, &Fish{days: 8})
			}
		}
		fish = append(fish, newFish...)
	}
	fmt.Println(len(fish))
}

func getFish() []*Fish {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(string(bytes.TrimSpace(data)), -1)
	fish := make([]*Fish, len(nums))
	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		fish[i] = &Fish{days: n}
	}
	return fish
}
