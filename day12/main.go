package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var farm [][]string

type position struct {
	x int
	y int
}

type region struct {
	plant     string
	area      int
	perimeter int
	positions []position
}

var totalCount int

var regions = []region{}

func getMap(x, y int) string {
	if x < 0 || y < 0 || x >= len(farm[0]) || y >= len(farm) {
		return ""
	}

	return farm[y][x]
}

func (e region) String() string {
	return fmt.Sprintf("[area: %d, perimeter %d] -- [%v]\n", e.area, e.perimeter, e.positions)
}

func Atoi(n string) int {
	if i, err := strconv.Atoi(n); err != nil {
		return 0
	} else {
		return i
	}
}

func drawMap() {
	for y := range farm {

		// Map
		for _, v := range farm[y] {
			var c *color.Color
			if v == strings.ToLower(v) {
				c = color.New(color.FgRed)
			} else {
				c = color.New(color.FgCyan)
			}
			c.Print(v)
		}
		fmt.Print(" || ")
		fmt.Println()
	}

	for i, r := range regions {
		fmt.Printf("[%d] = %v", i, r)
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

		var row []string
		for _, v := range line {
			f := string(v)
			row = append(row, f)
		}
		y++

		farm = append(farm, row)
	}
}
func addField(x, y int, p string) {
	for _, v := range regions {
		if v.plant == p {
			v.area += 1
			v.positions = append(v.positions, position{x, y})
			return
		}
	}
	// if we're here - no plant on the list

}

func scanRegion(x, y int, p string) (int, int) {
	var aCount int
	var pCount int

	directions := [][]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
	}

	// Out of boundaries
	if x < 0 || y < 0 || x >= len(farm[0]) || y >= len(farm) {
		return 0, 0
	}

	// Was here already
	if getMap(x, y) == strings.ToLower(getMap(x, y)) {
		return 0, 0
	}

	// Out of plant region
	if getMap(x, y) != p {
		return 0, 0
	}

	// Here means we found a new fragment
	// 1. Turn the field to lowercase
	farm[y][x] = strings.ToLower(getMap(x, y))
	// 2. Increase the count
	aCount = 1
	// 3. Count perimeters
	for _, d := range directions {
		lP := getMap(x+d[0], y+d[1])
		if lP == "" || !strings.EqualFold(lP, p) {
			pCount += 1
		}
	}

	for _, d := range directions {
		a, c := scanRegion(x+d[0], y+d[1], p)
		aCount += a
		pCount += c
	}

	return aCount, pCount
}

func run(file string) int {
	readInput(file)
	drawMap()

	// total := 0

	for y := range farm {
		for x := range farm[y] {
			aC, pC := scanRegion(x, y, getMap(x, y))
			if aC > 0 {
				totalCount += aC * pC
				fmt.Println("==== Plant: ", getMap(x, y), "Total:", totalCount, " Delta area:", aC, " Delta per", pC, "====================")
				// drawMap()
			}
		}
	}

	return totalCount
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
