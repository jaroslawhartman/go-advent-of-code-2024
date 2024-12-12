package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var areaMap [][]int
var visited [][]bool

var visitedCount = 1

type direction struct {
	x int
	y int
	s rune
}

type position struct {
	x int
	y int
	d direction
}

var guard position

func visitMap(x, y int) {
	if !visited[y][x] {
		visitedCount++
	}
	visited[y][x] = true
}

func getFromMap(x, y int) int {
	return areaMap[y][x]
}

func drawMap() {
	fmt.Print("\033[H\033[2J")
	fmt.Printf("======== Visited %d ====================================\n", visitedCount)
	for y := range areaMap {

		// Map
		for x, _ := range areaMap[y] {
			char := ""

			if x == guard.x && y == guard.y {
				char = string(guard.d.s)
			} else {
				v := getFromMap(x, y)
				if v == 0 {
					char = "."
				} else {
					char = "#"
				}
			}
			fmt.Print(char)
		}
		fmt.Print(" || ")
		// Visited
		for _, v := range visited[y] {
			char := ""
			if v {
				char = "X"
			} else {
				char = "."
			}
			fmt.Print(char)
		}

		fmt.Println()
	}
}

func turnGuardRight(g position) position {
	if g.d.s == '^' {
		g.d = direction{1, 0, '>'}
	} else if g.d.s == '>' {
		g.d = direction{0, 1, 'V'}
	} else if g.d.s == 'V' {
		g.d = direction{-1, 0, '<'}
	} else if g.d.s == '<' {
		g.d = direction{0, -1, '^'}
	}
	return g
}

func guardCharToDirection(g rune) direction {
	if g == '^' {
		return direction{0, -1, '^'}
	} else if g == '>' {
		return direction{1, 0, '>'}
	} else if g == 'v' || g == 'V' {
		return direction{0, 1, 'V'}
	} else if g == '<' {
		return direction{-1, 0, '<'}
	}

	return direction{0, 0, '!'}
}

func moveForward() {
	for {
		drawMap()
		newPos := guard
		newPos.x = guard.x + guard.d.x
		newPos.y = guard.y + guard.d.y

		if newPos.x < 0 || newPos.y < 0 || newPos.x >= len(areaMap[0]) || newPos.y >= len(areaMap) {
			fmt.Println("Leaving the map!")
			break
		}

		// Check if on an obstacle
		if getFromMap(newPos.x, newPos.y) == 0 {
			guard = newPos
			visitMap(newPos.x, newPos.y)
			continue
		}

		if getFromMap(newPos.x, newPos.y) != 0 {
			fmt.Println("On the obstacle!")
			guard = turnGuardRight(guard)
			continue
		}

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
		var v_row []bool
		for x, v := range line {
			if v == '.' {
				row = append(row, 0)
				v_row = append(v_row, false)
			} else if v == '#' {
				row = append(row, 1)
				v_row = append(v_row, false)
			} else {
				guard = position{x, y, guardCharToDirection(v)}
				row = append(row, 0)
				v_row = append(v_row, true)
			}
		}
		fmt.Println(row)
		y++
		areaMap = append(areaMap, row)
		visited = append(visited, v_row)
	}
}

func run(file string) int {
	readInput(file)

	total := 0

	moveForward()

	total = visitedCount

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
