package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

var input []string
var found [][]string

const pattern = "XMAS"

func getChar(x int, y int) string {

	maxY := len(input[0])
	maxX := len(input)

	if x >= maxX || y >= maxY || x < 0 || y < 0 {
		return "@"
	}

	x, y = normalizeXY(x, y)

	return string(input[y][x])
}

func normalizeX(v int) int {
	max := len(input[0])

	output := v % max

	if output < 0 {
		output = max + output
	}
	return output
}

func normalizeY(v int) int {
	max := len(input)

	output := v % max

	if output < 0 {
		output = max + output
	}
	return output
}

func normalizeXY(x int, y int) (int, int) {
	x = normalizeX(x)
	y = normalizeX(y)

	return x, y
}

func readInput(file string) {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Can't open file: " + err.Error())
	}

	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	for s.Scan() {
		fmt.Println(s.Text())

		input = append(input, s.Text())
		// fmt.Println(input)
	}
}

func findXMAS(x int, y int, moveX int, moveY int) bool {

	for i := range pattern {
		if getChar(x+moveX*i, y+moveY*i) != string(pattern[i]) {
			return false
		}
	}

	for i := range pattern {
		found[normalizeY(x+moveX*i)][normalizeY(y+moveY*i)] = getChar(x+moveX*i, y+moveY*i)
	}

	return true
}

func initFound() {
	for y := 0; y < len(input); y++ {
		var line []string
		for x := 0; x < len(input[0]); x++ {
			line = append(line, ".")
		}

		found = append(found, []string(line))
	}
}

func printSummary(m int, n int) {

	// red := color.New(color.FgRed, color.Bold)

	fmt.Println("---------- SUMMARY ----------")

	for y := 0; y < len(found); y++ {
		for x := 0; x < len(found[0]); x++ {
			if x == m && y == n {
				color.Set(color.FgRed, color.Bold)
				fmt.Printf("%s", string(found[x][y]))
				color.Unset()
			} else {
				fmt.Printf("%s", string(found[x][y]))
			}
		}
		fmt.Println()
	}
}

func run(file string) int {

	directions := [][]int{{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	readInput(file)
	initFound()

	total := 0

	// for y := range input {
	// 	for x := range input[y] {
	// 		fmt.Printf("%s ", getChar(x, y))
	// 	}
	// 	fmt.Println()
	// }

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {

			for _, v := range directions {
				if findXMAS(x, y, v[0], v[1]) {
					total += 1
				}
			}
			// printSummary(x, y)
		}
		fmt.Println()
	}

	return total
}

func main() {

	total := run("input.txt")

	fmt.Println("Total: ", total)
}
