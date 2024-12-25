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
	done      bool
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
	return fmt.Sprintf("[area: %d, perimeter %d, done: %v] -- [%v]\n", e.area, e.perimeter, e.done, e.positions)
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
		fmt.Printf("[field: %s][%d] = %v", r.plant, i, r)
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
	for i := range regions {
		if regions[i].plant == p && !regions[i].done {
			regions[i].area += 1
			regions[i].positions = append(regions[i].positions, position{x, y})
			return
		}
	}
	// if we're here - no plant on the list

	r := region{
		plant:     p,
		area:      1,
		perimeter: 0,
		positions: []position{},
	}

	r.positions = append(r.positions, position{x, y})

	regions = append(regions, r)
}

func findField(x, y int, i int) bool {
	if x < 0 || x >= len(farm[0]) || y < 0 || y > len(farm) {
		return false
	}

	v := regions[i]

	for _, pos := range v.positions {
		newPos := position{x: x, y: y}
		if pos == newPos {
			return true
		}
	}

	return false
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

	// 4.
	addField(x, y, p)

	for _, d := range directions {
		a, c := scanRegion(x+d[0], y+d[1], p)
		aCount += a
		pCount += c
	}

	return aCount, pCount
}

func countEdges() int {
	fmt.Println("==========================")

	totalCount := 0

	for i, r := range regions {
		fmt.Printf("[field: %s][%d] = %v", r.plant, i, r)

		found := false

		perimeter := 0

		//       |
		// Edges V
		for y := 0; y < len(farm); y++ {
			for x := 0; x < len(farm[y]); x++ {
				if findField(x, y, i) && !findField(x, y-1, i) {
					if !found {
						perimeter += 1
						found = true
					}
				} else {
					found = false
				}

				// if findField(x, y, i) {
				// 	fmt.Printf(" ... [%s %d,%d] - Field: %v, Perim: %d, Found %v ", r.plant, x, y, findField(x, y, i), perimeter, found)
				// }
			}
			// fmt.Println()
		}
		// Edges  ^
		//        |
		for y := 0; y < len(farm); y++ {
			for x := 0; x < len(farm[y]); x++ {
				if findField(x, y, i) && !findField(x, y+1, i) {
					if !found {
						perimeter += 1
						found = true
					}
				} else {
					found = false
				}

				// if findField(x, y, i) {
				// 	fmt.Printf(" ... [%s %d,%d] - Field: %v, Perim: %d, Found %v ", r.plant, x, y, findField(x, y, i), perimeter, found)
				// }
			}
			// fmt.Println()
		}

		// Edges  ->
		//
		for x := 0; x < len(farm[0]); x++ {
			for y := 0; y < len(farm); y++ {
				if findField(x, y, i) && !findField(x-1, y, i) {
					if !found {
						perimeter += 1
						found = true
					}
				} else {
					found = false
				}

				// if findField(x, y, i) {
				// 	fmt.Printf(" ... [%s %d,%d] - Field: %v, Perim: %d, Found %v ", r.plant, x, y, findField(x, y, i), perimeter, found)
				// }
			}
			// fmt.Println()
		}

		// Edges  <--
		//
		for x := 0; x < len(farm[0]); x++ {
			for y := 0; y < len(farm); y++ {
				if findField(x, y, i) && !findField(x+1, y, i) {
					if !found {
						perimeter += 1
						found = true
					}
				} else {
					found = false
				}

				// if findField(x, y, i) {
				// 	fmt.Printf(" ... [%s %d,%d] - Field: %v, Perim: %d, Found %v ", r.plant, x, y, findField(x, y, i), perimeter, found)
				// }
			}
			// fmt.Println()
		}
		regions[i].perimeter += perimeter
		fmt.Printf(" ... [%s] - Area %d,  Perimeter %d, Total %d \n", r.plant, r.area, regions[i].perimeter, r.area*regions[i].perimeter)

		totalCount += (perimeter * r.area)
	}

	return totalCount
}

func run(file string) int {
	readInput(file)
	drawMap()

	total := 0

	for y := range farm {
		for x := range farm[y] {
			aC, pC := scanRegion(x, y, getMap(x, y))
			if aC > 0 {
				totalCount += aC * pC
				fmt.Println("==== Plant: ", getMap(x, y), "Total:", totalCount, " Delta area:", aC, " Delta per", pC, "====================")
				regions[len(regions)-1].done = true
				// drawMap()
			}
		}
	}

	total = countEdges()

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
