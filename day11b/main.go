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

var blink int

func Atoi(n string) int {
	if i, err := strconv.Atoi(n); err != nil {
		return 0
	} else {
		return i
	}
}

func displayMap() {
	fmt.Printf("[%d] (len: %d) %v\n", blink, len(stones), stones)
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

func doBlink() {
	newStones := []int{}

	for _, v := range stones {
		strV := fmt.Sprintf("%d", v)
		if v == 0 {
			newStones = append(newStones, 1)
		} else if len(strV)%2 == 0 {

			newStones = append(newStones, Atoi(strV[:len(strV)/2]))
			newStones = append(newStones, Atoi(strV[len(strV)/2:]))
		} else {
			newStones = append(newStones, v*2024)
		}
	}

	stones = newStones
}

func run(file string) int {
	readInput(file)
	displayMap()

	for b := range 25 {
		blink = b + 1
		doBlink()
		displayMap()
	}

	total := len(stones)

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
