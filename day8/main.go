package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(p2("input.txt"))
}

// not sure if this is the right approach here. moving on to another puzzle... will return to it later
func p2(filename string) int {
	rows := getInput(filename)
	var sum int
	for _, row := range rows {
		sum += getOutput(row)
	}
	return sum
}
func byteSortedString(d []byte) string {
	ints := make([]int, len(d))
	for i, b := range d {
		ints[i] = int(b)
	}
	sort.Ints(ints)
	var sb strings.Builder
	for _, i := range ints {
		sb.WriteString(strconv.Itoa(i))
	}
	return sb.String()
}

func getOutput(d Data) int {
	all := append(d.Input, d.Output...)

	ones := findOnes(all)
	sevens := findSevens(all)
	sixes := findSixes(all, sevens[0])
	twos := findTwos(all, ones[0], sixes[0])
	threes := findThrees(all, ones[0])
	fours := findFours(all)
	fives := findFives(all, sixes[0])
	eights := findEights(all)
	nines := findNines(all, threes[0])
	zeros := findZeros(all, threes[0], sevens[0])

	one := byteSortedString(ones[0])
	two := byteSortedString(twos[0])
	three := byteSortedString(threes[0])
	four := byteSortedString(fours[0])
	five := byteSortedString(fives[0])
	six := byteSortedString(sixes[0])
	seven := byteSortedString(sevens[0])
	eight := byteSortedString(eights[0])
	nine := byteSortedString(nines[0])
	zero := byteSortedString(zeros[0])

	strs := make([]string, len(d.Output))
	for i, digit := range d.Output {
		str := byteSortedString(digit)
		switch str {
		case one:
			strs[i] = "1"
		case two:
			strs[i] = "2"
		case three:
			strs[i] = "3"
		case four:
			strs[i] = "4"
		case five:
			strs[i] = "5"
		case six:
			strs[i] = "6"
		case seven:
			strs[i] = "7"
		case eight:
			strs[i] = "8"
		case nine:
			strs[i] = "9"
		case zero:
			strs[i] = "0"
		}
	}
	i, err := strconv.Atoi(strings.Join(strs, ""))
	if err != nil {
		panic(err)
	}
	return i
}

func p1(filename string) int {
	rows := getInput(filename)
	var outputs [][]byte
	for _, row := range rows {
		outputs = append(outputs, row.Output...)
	}
	ones := findOnes(outputs)
	fours := findFours(outputs)
	sevens := findSevens(outputs)
	eights := findEights(outputs)
	return len(ones) + len(fours) + len(sevens) + len(eights)
}

func findByLen(digits [][]byte, l int) [][]byte {
	var d [][]byte
	for _, digit := range digits {
		if len(digit) == l {
			d = append(d, digit)
		}
	}
	return d
}

func containsDigit(d1, d2 []byte) bool {
Outer:
	for _, b := range d2 {
		for _, b2 := range d1 {
			if b2 == b {
				continue Outer
			}
		}
		return false
	}
	return true
}

func findOnes(digits [][]byte) [][]byte {
	return findByLen(digits, 2)
}

func findTwos(digits [][]byte, one, six []byte) [][]byte {
	potentials := findByLen(digits, 5)
	var twos [][]byte
	for _, potential := range potentials {
		if containsDigit(potential, one) {
			continue
		}
		if containsDigit(six, potential) {
			continue
		}
		twos = append(twos, potential)
	}
	return twos
}

func findThrees(digits [][]byte, one []byte) [][]byte {
	potentials := findByLen(digits, 5)
	var threes [][]byte
	for _, potential := range potentials {
		if containsDigit(potential, one) {
			threes = append(threes, potential)
		}
	}
	return threes
}

func findFours(digits [][]byte) [][]byte {
	return findByLen(digits, 4)
}

func findFives(digits [][]byte, six []byte) [][]byte {
	potentials := findByLen(digits, 5)
	var fives [][]byte
	for _, potential := range potentials {
		if containsDigit(six, potential) {
			fives = append(fives, potential)
		}
	}
	return fives
}

func findSixes(digits [][]byte, seven []byte) [][]byte {
	potentials := findByLen(digits, 6)
	var sixes [][]byte
	for _, potential := range potentials {
		if !containsDigit(potential, seven) {
			sixes = append(sixes, potential)
		}
	}
	return sixes
}

func findSevens(digits [][]byte) [][]byte {
	return findByLen(digits, 3)
}

func findEights(digits [][]byte) [][]byte {
	return findByLen(digits, 7)
}

func findNines(digits [][]byte, three []byte) [][]byte {
	potentials := findByLen(digits, 6)
	var nines [][]byte
	for _, potential := range potentials {
		if containsDigit(potential, three) {
			nines = append(nines, potential)
		}
	}
	return nines
}

func findZeros(digits [][]byte, three, seven []byte) [][]byte {
	potentials := findByLen(digits, 6)
	var zeros [][]byte
	for _, potential := range potentials {
		if containsDigit(potential, three) {
			// is a 9
			continue
		}
		if containsDigit(potential, seven) {
			zeros = append(zeros, potential)
		}
	}
	return zeros
}

func getInput(filename string) []Data {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	rows := bytes.Split(bytes.TrimSpace(data), []byte("\n"))
	dataRows := make([]Data, len(rows))
	for i, row := range rows {
		s := bytes.SplitN(row, []byte("|"), 2)
		dataRows[i] = Data{
			Input:  bytes.Split(bytes.TrimSpace(s[0]), []byte(" ")),
			Output: bytes.Split(bytes.TrimSpace(s[1]), []byte(" ")),
		}
	}
	return dataRows
}

type Data struct {
	Input  [][]byte
	Output [][]byte
}
