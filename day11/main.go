package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(p2("input.txt"))
}

func p1(filename string) int {
	d := getData(filename)
	var count int
	for i := 0; i < 100; i++ {
		for _, o := range d {
			o.flashed = false
		}
		for _, o := range d {
			count += o.incr()
		}
	}
	return count
}

func p2(filename string) int {
	d := getData(filename)
	var step int
	for {
		step++
		var count int
		for _, o := range d {
			o.flashed = false
		}
		for _, o := range d {
			count += o.incr()
		}
		if count == len(d) {
			return step
		}
	}
}

func getData(filename string) []*octopus {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	rows := bytes.Split(data, []byte("\n"))
	height := len(rows)
	width := len(rows[0])
	octopi := make([]*octopus, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			energyLevel, err := strconv.Atoi(string(rows[y][x]))
			if err != nil {
				panic(err)
			}
			o := &octopus{energy: energyLevel, x: x, y: y}
			octopi[(y*width)+x] = o
			ll, lr, lu := x > 0, x < width-1, y > 0

			if ll {
				octopi[(y*width)+x-1].addNeighbor(o)
				o.addNeighbor(octopi[(y*width)+x-1])
			}

			if lu {
				octopi[((y-1)*width)+x].addNeighbor(o)
				o.addNeighbor(octopi[((y-1)*width)+x])
			}

			if ll && lu {
				octopi[((y-1)*width)+x-1].addNeighbor(o)
				o.addNeighbor(octopi[((y-1)*width)+x-1])
			}

			if lr && lu {
				octopi[((y-1)*width)+x+1].addNeighbor(o)
				o.addNeighbor(octopi[((y-1)*width)+x+1])
			}

		}
	}

	return octopi
}

type octopus struct {
	energy    int
	x, y      int
	neighbors []*octopus
	flashed   bool
}

func (o *octopus) flash() int {
	o.flashed = true
	o.energy = 0
	count := 1
	for _, neighbor := range o.neighbors {
		count += neighbor.incr()
	}
	return count
}

func (o *octopus) incr() int {
	if o.flashed {
		return 0
	}
	o.energy++
	if o.energy > 9 {
		return o.flash()
	}
	return 0
}

func (o *octopus) addNeighbor(oct *octopus) {
	o.neighbors = append(o.neighbors, oct)
}
