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

func safe(r report) bool {
	prevDelta := 0

	for i := range r {
		if i == 0 {
			continue
		}

		delta := r[i] - r[i-1]

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
		isSafe := safe(v)
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
