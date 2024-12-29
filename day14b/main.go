package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/fatih/color"
)

var totalCost int

type robot struct {
	px int
	py int
	vx int
	vy int
}

var robots []robot

var maxX, maxY int

func (r robot) String() string {
	return fmt.Sprintf("Pos [%d,%d] V: [%d,%d]", r.px, r.py, r.vx, r.vy)
}

func Atoi(n string) int {
	if i, err := strconv.Atoi(n); err != nil {
		return 0
	} else {
		return i
	}
}

func findRobotCount(x, y int) int {
	count := 0
	for _, r := range robots {
		if r.px == x && r.py == y {
			count += 1
		}
	}
	return count
}

func findRobot(x, y int) int {
	for i, r := range robots {
		if r.px == x && r.py == y {
			return i
		}
	}
	return -1
}

func moveRobots() {
	for i := range robots {
		robots[i].px = (robots[i].px + robots[i].vx)
		robots[i].py = (robots[i].py + robots[i].vy)

		if robots[i].px >= maxX {
			robots[i].px = robots[i].px % (maxX)
		}

		if robots[i].py >= maxY {
			robots[i].py = robots[i].py % (maxY)
		}

		if robots[i].px < 0 {
			robots[i].px = robots[i].px + maxX
		}

		if robots[i].py < 0 {
			robots[i].py = robots[i].py + maxY
		}

	}
}

func countRobots(qx, qy int) int {
	var startX, startY, stopX, stopY int

	if qx == 0 {
		startX = 0
		stopX = maxX / 2
	} else if qx == 1 {
		startX = maxX/2 + 1
		stopX = maxX
	} else {
		return -1
	}

	if qy == 0 {
		startY = 0
		stopY = maxY / 2
	} else if qy == 1 {
		startY = maxY/2 + 1
		stopY = maxY
	} else {
		return -1
	}

	count := 0

	for y := startY; y < stopY; y++ {
		for x := startX; x < stopX; x++ {
			count += findRobotCount(x, y)
		}
	}
	return count
}

func drawMap(i int) {
	fmt.Println("====== Iteration ", i, "====== Counts: ", countRobots(0, 0), countRobots(1, 0), countRobots(0, 1), countRobots(1, 1), "=======")
	for y := range maxY {
		// Map
		for x := range maxX {
			var c *color.Color
			var v string

			if x == (maxX/2) || x == (maxX/2+1) {
				fmt.Print(" ")

			}

			count := findRobotCount(x, y)

			if count > 0 {
				c = color.New(color.FgRed)
				v = fmt.Sprintf("%x", count)
			} else {
				c = color.New(color.FgCyan)
				v = "."
			}
			c.Print(v)
		}
		// fmt.Print(" || ")
		// for x := range maxX {
		// 	var c *color.Color
		// 	var v string

		// 	id := findRobot(x, y)

		// 	if id >= 0 {
		// 		c = color.New(color.FgRed)
		// 		v = fmt.Sprintf("%x", id)
		// 	} else {
		// 		c = color.New(color.FgCyan)
		// 		v = "."
		// 	}
		// 	c.Print(v)
		// }

		fmt.Println()
	}
}

func readInput(file string) {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Can't open file: " + err.Error())
	}

	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	reRobot := regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

	for s.Scan() {
		line := s.Text()
		res := reRobot.FindStringSubmatch(line)
		r := robot{
			px: Atoi(res[1]),
			py: Atoi(res[2]),
			vx: Atoi(res[3]),
			vy: Atoi(res[4]),
		}

		robots = append(robots, r)
		fmt.Println(r)
	}
}

func run(file string, x, y int) int {
	maxX = x
	maxY = y
	readInput(file)

	for i := range 1000000 {
		drawMap(i + 1)
		moveRobots()
	}

	total := countRobots(0, 0) * countRobots(1, 0) * countRobots(0, 1) * countRobots(1, 1)

	return total
}

func main() {
	total := run("input.txt", 101, 103)

	fmt.Println("Total: ", total)
}
