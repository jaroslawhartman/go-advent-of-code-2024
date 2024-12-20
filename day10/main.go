package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

type trailhead struct {
	x     int
	y     int
	score int
}

var peaks []position

var topoMap [][]int
var trailheads []trailhead

func getMap(x, y int) int {
	return topoMap[y][x]
}

func (e trailhead) String() string {
	return fmt.Sprintf("[%d, %d] -> %d\n", e.x, e.y, e.score)
}

func Atoi(n string) int {
	if i, err := strconv.Atoi(n); err != nil {
		return 0
	} else {
		return i
	}
}

func drawMap() {
	// fmt.Print("\033[H\033[2J")
	// fmt.Printf("======== Visited %d ======= No new %d ==== Obstacles %d ====\n", visitedCount, movesWithoutNewVisited, obstaclesFound)
	for y := range topoMap {

		// Map
		for _, v := range topoMap[y] {
			fmt.Print(v)
		}
		fmt.Print(" || ")
		// // antinodes
		// for _, v := range antinodes[y] {
		// 	char := ""
		// 	if v {
		// 		char = "#"
		// 	} else {
		// 		char = "."
		// 	}
		// 	fmt.Print(char)
		// }

		fmt.Println()
	}

	// trailheads
	for _, v := range trailheads {
		fmt.Print(v)
	}
}

func readInput(file string) {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Can't open file: " + err.Error())
	}

	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	var y = 0
	for s.Scan() {
		line := s.Text()

		var row []int
		for x, v := range line {
			height := Atoi(string(v))
			row = append(row, height)

			if height == 0 {
				th := trailhead{
					x:     x,
					y:     y,
					score: 0,
				}

				trailheads = append(trailheads, th)
			}

		}
		// fmt.Println(row)
		y++

		topoMap = append(topoMap, row)
	}
}

func deepCopyIntSlice(src [][]int) [][]int {
	cpy := make([][]int, len(src))
	for i := range src {
		cpy[i] = make([]int, len(src[i]))
		copy(cpy[i], src[i])
	}
	return cpy
}

func deepCopyBoolSlice(src [][]bool) [][]bool {
	cpy := make([][]bool, len(src))
	for i := range src {
		cpy[i] = make([]bool, len(src[i]))
		copy(cpy[i], src[i])
	}
	return cpy
}

func findPeak(x, y int) bool {
	for _, v := range peaks {
		if v.x == x && v.y == y {
			return true
		}
	}
	return false
}

func findTrails(id int, level int, x, y int, h int) int {
	var trails int

	directions := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	for _, d := range directions {
		dx := d[0]
		dy := d[1]

		if x+dx < 0 || x+dx >= len(topoMap[0]) || y+dy < 0 || y+dy >= len(topoMap) {
			fmt.Printf("%s[%d][%d,%d] -> [%d,%d] -- NOK\n", strings.Repeat(" ", level), id, x, y, x+dx, y+dy)
		} else {
			newHeight := getMap(x+dx, y+dy)
			fmt.Printf("%s[%d][%d,%d] -> [%d,%d] -- NewHeight: %d", strings.Repeat(" ", level), id, x, y, x+dx, y+dy, newHeight)

			if newHeight == h+1 && newHeight != 9 {
				fmt.Printf(" ... Continuing (%d ^ %d)\n", h, h+1)
				trails += findTrails(id, level+1, x+dx, y+dy, h+1)
			} else if newHeight == 9 && h == 8 {
				if !findPeak(x+dx, y+dy) {
					trails += 1
					peaks = append(peaks, position{x + dx, y + dy})
					fmt.Printf(" ... End! (+1)\n")
				} else {
					fmt.Printf(" ... End! (was here)\n")
				}
			} else {
				fmt.Printf("\n")
			}
		}
	}

	return trails
}

func run(file string) int {
	readInput(file)
	drawMap()

	total := 0
	// trailheads
	for id, th := range trailheads {
		peaks = []position{}
		th.score = findTrails(id+1, 0, th.x, th.y, 0)

		fmt.Print("### Finished ", id+1, th)

		total += th.score
	}

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
