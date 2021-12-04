package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(data, []byte("\n"))
	var (
		aim      int
		depth    int
		position int
	)
	for _, row := range rows {
		dir, mag := splitRow(row)
		//fmt.Println("aim:", aim, "depth:", depth, "position", position, "dir:", dir, "mag:", mag)
		switch dir {
		case "forward":
			depth += aim * mag
			position += mag
		case "up":
			aim -= mag
		case "down":
			aim += mag
		}
	}
	fmt.Println("depth:", depth, "position:", position)
	fmt.Println("product:", depth*position)
}

func splitRow(row []byte) (string, int) {
	data := strings.Split(string(row), " ")
	i, err := strconv.Atoi(data[1])
	if err != nil {
		panic(err)
	}
	return data[0], i
}
