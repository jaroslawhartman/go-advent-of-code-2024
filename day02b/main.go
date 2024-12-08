package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type report []int

var reports []report

func Abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func Normalize(i int) int {
	if i >= 0 {
		return 1
	} else {
		return -1
	}
}

func remove(slice []int, s int) []int {
	var result report
	result = append(result, slice[:s]...)
	result = append(result, slice[s+1:]...)
	return result
}

func safe(r report, removed int) bool {
	prevDelta := 0
	var rep report

	if removed >= 0 {
		rep = remove(r, removed)
	} else {
		rep = r
	}

	// fmt.Println(removed, ">>>", rep)

	for i := range rep {
		if i == 0 {
			continue
		}

		delta := rep[i] - rep[i-1]

		if prevDelta != 0 {
			if Normalize(prevDelta*delta) == -1 {
				return false
			}
		}

		prevDelta = delta

		if Abs(delta) > 3 || delta == 0 {
			return false
		}
	}
	return true
}

func readInput(file string) {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Can't open file: " + err.Error())
	}

	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	for s.Scan() {
		levelsStr := strings.Split(s.Text(), " ")

		var report report

		for _, v := range levelsStr {
			i, _ := strconv.Atoi(v)
			report = append(report, i)
		}

		reports = append(reports, report)
	}
}

func run(file string) int {
	readInput(file)
	total := 0

	for i, v := range reports {
		isSafe := safe(v, -1)

		if !isSafe {
			for removed := 0; removed < len(v); removed++ {
				isSafe = safe(v, removed)

				if isSafe {
					break
				}
			}
		}

		fmt.Println(i, v, isSafe)

		if isSafe {
			total += 1
		}
	}

	return total
}

func main() {

	total := run("input.txt")

	fmt.Println("Total: ", total)
}
