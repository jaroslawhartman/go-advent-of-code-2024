package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var ordering [][]int
var updates [][]int

func getMiddlePage(u []int) int {
	return u[(len(u))/2]
}

func checkBackward(idx int, u []int) bool {
	// fmt.Println("Checking bwd....: ", idx)
	if idx < 1 {
		return true
	}
	// fmt.Println("Checking bwd......: ", idx)

	for i := idx - 1; i >= 0; i-- {
		if slices.IndexFunc(ordering, func(update []int) bool {
			if update[0] == u[i] && update[1] == u[idx] {
				return true
			} else {
				return false
			}
		}) == -1 {
			return false
		}
	}
	return true
}

func checkForward(idx int, u []int) bool {
	if idx+1 > len(u) {
		return true
	}

	for i := idx + 1; i < len(u); i++ {
		// fmt.Println("Checking fwd: ", u[idx], u[i])
		if slices.IndexFunc(ordering, func(update []int) bool {
			if update[0] == u[idx] && update[1] == u[i] {
				return true
			} else {
				return false
			}
		}) == -1 {
			return false
		}
	}
	return true
}

func checkUpdate(u []int) int {
	for i, _ := range u {
		if !(checkForward(i, u) && checkBackward(i, u)) {
			return 0
		}
	}
	return getMiddlePage(u)
}

func Atoi(n string) int {
	if i, err := strconv.Atoi(n); err != nil {
		return 0
	} else {
		return i
	}
}

func addOrderingRule(t string) {
	var o []int

	split := strings.Split(t, "|")

	for _, v := range split {
		o = append(o, Atoi(v))
	}

	ordering = append(ordering, o)
}

func addUpdate(t string) {
	var u []int

	split := strings.Split(t, ",")

	for _, v := range split {
		u = append(u, Atoi(v))
	}

	updates = append(updates, u)
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

		if strings.Contains(line, "|") {
			// Ordering rules
			addOrderingRule(line)
		} else if strings.Contains(line, ",") {
			// Updates
			addUpdate(line)
		}
		// fmt.Println(input)
	}
}

func fixUpdate(u []int) []int {
	// var fixed []int

	slices.SortFunc(u, func(a, b int) int {

		cmp1 := slices.IndexFunc(ordering, func(o []int) bool {
			if o[0] == a && o[1] == b {
				return true
			} else {
				return false
			}
		})

		cmp2 := slices.IndexFunc(ordering, func(o []int) bool {
			if o[0] == b && o[1] == a {
				return true
			} else {
				return false
			}
		})

		if cmp1 >= 0 {
			return -1
		}

		if cmp2 >= 0 {
			return 1
		}
		return 0
	})

	return u
}

func run(file string) int {
	readInput(file)

	total := 0

	fmt.Println(ordering)
	fmt.Println(updates)

	for _, v := range updates {
		value := checkUpdate(v)
		if value == 0 {
			fmt.Println("Bad:               ", v, value)
			v = fixUpdate(v)
			value = checkUpdate(v)
			fmt.Println("Bad (but fixed):   ", v, value)
			if value == 0 {
				fmt.Println("WARN - wrong fix above!")
			}
			total += value
		} else {
			// fmt.Println("Good: ", v, value)
		}

	}

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
