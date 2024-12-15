package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type position struct {
	x int
	y int
}

type antenna struct {
	freq      rune
	positions []position
}

func (e antenna) String() string {
	return fmt.Sprintf("%s: %v\n", string(e.freq), e.positions)
}

var antennas []antenna
var antennasMap [][]rune
var antinodes [][]bool

var totalAntinodes int

func drawMap() {
	// fmt.Print("\033[H\033[2J")
	// fmt.Printf("======== Visited %d ======= No new %d ==== Obstacles %d ====\n", visitedCount, movesWithoutNewVisited, obstaclesFound)
	for y := range antennasMap {

		// Map
		for _, v := range antennasMap[y] {
			fmt.Print(string(v))
		}
		fmt.Print(" || ")
		// antinodes
		for _, v := range antinodes[y] {
			char := ""
			if v {
				char = "#"
			} else {
				char = "."
			}
			fmt.Print(char)
		}

		fmt.Println()
	}
}

func addAntenna(p position, freq rune) {
	i := slices.IndexFunc(antennas, func(a antenna) bool {
		return a.freq == freq
	})

	if i != -1 {
		antennas[i].positions = append(antennas[i].positions, p)
		return
	}

	a := antenna{
		freq:      freq,
		positions: nil,
	}

	a.positions = append(a.positions, p)
	antennas = append(antennas, a)
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

		var row []rune   // antennasMap
		var a_row []bool // antinodes
		for x, v := range line {
			if v == '.' {
				row = append(row, v)
			} else {
				row = append(row, v)
				addAntenna(position{x, y}, v)
			}
			a_row = append(a_row, false)
		}
		// fmt.Println(row)
		y++
		antennasMap = append(antennasMap, row)
		antinodes = append(antinodes, a_row)
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

func setAntinode(p1, p2 position) {
	vx := p2.x - p1.x
	vy := p2.y - p1.y

	ax := p2.x + vx
	ay := p2.y + vy

	fmt.Printf("p1: %v, p2: %v, v: %d,%d, a: %d,%d\n", p1, p2, vx, vy, ax, ay)

	if ax >= 0 && ay >= 0 && ax < len(antinodes[0]) && ay < len(antinodes) {
		if !antinodes[ay][ax] {
			antinodes[ay][ax] = true
			totalAntinodes++
		}
	}

}

func checkAntenna(a antenna) {
	fmt.Println("Checking freq", string(a.freq))
	for _, p1 := range a.positions {
		for _, p2 := range a.positions {
			if p1 == p2 {
				continue
			}
			setAntinode(p1, p2)
			setAntinode(p2, p1)
		}
	}
}

func run(file string) int {
	readInput(file)
	fmt.Println(antennas)

	for _, a := range antennas {
		checkAntenna(a)
	}
	drawMap()

	total := totalAntinodes

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
