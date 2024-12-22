package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type entry struct {
	stone int
	depth int
}

var cache map[entry]int

var stones []int

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

func doBlink(stone, depth, maxDepth, count int) int {
	localCount := 0

	e := entry{
		stone: stone, depth: depth,
	}

	v, ok := cache[e]

	if ok {
		// fmt.Printf("Found in cache: %d = %v\n", v, e)
		return v
	}

	if depth == maxDepth {
		return 1
	}

	if stone == 0 {
		localCount += doBlink(1, depth+1, maxDepth, count)
	} else if numDigits := int(math.Log10(float64(stone))) + 1; numDigits%2 == 0 {

		divisor := int(math.Pow10(numDigits / 2))

		localCount += doBlink(stone/divisor, depth+1, maxDepth, count)
		localCount += doBlink(stone%divisor, depth+1, maxDepth, count)
	} else {
		localCount += doBlink(stone*2024, depth+1, maxDepth, count)
	}

	cache[e] = localCount

	return localCount
}

func run(file string) int {
	var total int
	readInput(file)
	displayMap()

	cache = make(map[entry]int)

	for _, stone := range stones {
		total += doBlink(stone, 0, 75, 1)
	}

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
