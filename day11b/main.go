package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var stones []int

var total int

var maxDepth int

func Atoi(n string) int {
	if i, err := strconv.Atoi(n); err != nil {
		return 0
	} else {
		return i
	}
}

func displayMap() {
	fmt.Printf("(len: %d) %v\n", len(stones), stones)
}

func readInput(file string) {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Can't open file: " + err.Error())
	}

	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()

		numbers := strings.Split(line, " ")

		for _, n := range numbers {
			stones = append(stones, Atoi(n))
		}
	}
}

func doBlink(stone, depth, maxDepth int) {
	if depth == maxDepth {
		return
	}

	strStone := fmt.Sprintf("%d", stone)

	if stone == 0 {
		doBlink(1, depth+1, maxDepth)
	} else if len(strStone)%2 == 0 {
		total += 1
		doBlink(Atoi(strStone[:len(strStone)/2]), depth+1, maxDepth)
		doBlink(Atoi(strStone[len(strStone)/2:]), depth+1, maxDepth)
	} else {
		doBlink(stone*2024, depth+1, maxDepth)
	}
}

func run(file string) int {
	readInput(file)
	displayMap()

	for _, stone := range stones {
		total += 1
		doBlink(stone, 0, 25)
	}

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
