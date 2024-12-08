package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

var locA []int
var locB []int

func distance(a int, b int) int {
	return int(math.Abs(float64(b) - float64(a)))
}

func readInput() {

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("Can't open file: " + err.Error())
	}

	s := bufio.NewScanner(file)

	s.Split(bufio.ScanLines)

	var a, b int

	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d   %d", &a, &b)
		// fmt.Scanln(s.Text(), "%d %d", &a, &b)
		locA = append(locA, a)
		locB = append(locB, b)
	}
}

func day01() int {
	slices.Sort(locA)
	slices.Sort(locB)

	var total = 0

	for i := range locA {
		// fmt.Println(locA[i], locB[i], distance(locA[i], locB[i]))
		total += distance(locA[i], locB[i])
	}

	return total
}

func main() {
	readInput()

	total := day01()

	fmt.Println("Day01 total: ", total)
}
