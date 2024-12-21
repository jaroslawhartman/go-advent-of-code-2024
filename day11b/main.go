package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

var stones []int

type atom struct {
	value atomic.Int64
}

var total atom = atom{}

// func incTotal(v int) {
// 	total.mu.Lock()
// 	defer total.mu.Unlock()

// 	total.total += v
// }

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

	var wg sync.WaitGroup

	if depth == maxDepth {
		return
	}

	current := total.value.Load()

	if current%10000000 == 0 {
		fmt.Printf("Depth: %d,  Total: %d\n", depth, current)
	}

	if stone == 0 {
		wg.Add(1)
		go func() {
			doBlink(1, depth+1, maxDepth)
			defer wg.Done()
		}()
	} else if numDigits := int(math.Log10(float64(stone))) + 1; numDigits%2 == 0 {
		total.value.Add(1)

		wg.Add(1)
		go func() {
			// Calculate the divisor to separate the halves
			divisor := int(math.Pow10(numDigits / 2))

			doBlink(stone/divisor, depth+1, maxDepth)
			doBlink(stone%divisor, depth+1, maxDepth)
			defer wg.Done()
		}()
	} else {
		wg.Add(1)
		go func() {
			doBlink(stone*2024, depth+1, maxDepth)
			defer wg.Done()
		}()
	}

	wg.Wait()
}

func run(file string) int {
	readInput(file)
	displayMap()

	var wg sync.WaitGroup

	for _, stone := range stones {
		total.value.Add(1)
		wg.Add(1)
		go func() {
			doBlink(stone, 0, 25)
			defer wg.Done()
		}()
	}

	wg.Wait()

	return int(total.value.Load())
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
