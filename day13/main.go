package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var totalCost int

func Atoi(n string) int {
	if i, err := strconv.Atoi(n); err != nil {
		return 0
	} else {
		return i
	}
}

func checkClawMachine(Xa, Ya, Xb, Yb, Px, Py float64) int {
	var A, B float64

	B = (Xa*Py - Ya*Px) / (Yb*Xa - Xb*Ya)
	A = (Px - B*Xb) / Xa

	fmt.Println("A,B", A, B)

	if A == math.Trunc(A) && B == math.Trunc(B) {
		return int(A*3.0 + B*1.0)
	}

	return 0
}

func readInput(file string) {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Can't open file: " + err.Error())
	}

	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	reButton := regexp.MustCompile(`Button [A|B]\: X\+(\d+), Y\+(\d+)`)
	rePrize := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	var Xa, Ya, Xb, Yb, Px, Py int

	for s.Scan() {
		line := s.Text()
		if strings.Contains(line, "Button A") {
			res := reButton.FindStringSubmatch(line)
			fmt.Println("A", res)
			Xa = Atoi(res[1])
			Ya = Atoi(res[2])
		}

		if strings.Contains(line, "Button B") {
			res := reButton.FindStringSubmatch(line)
			fmt.Println("B", res)
			Xb = Atoi(res[1])
			Yb = Atoi(res[2])
		}

		if strings.Contains(line, "Prize") {
			res := rePrize.FindStringSubmatch(line)
			fmt.Println("P", res)
			Px = Atoi(res[1])
			Py = Atoi(res[2])

			cost := checkClawMachine(float64(Xa), float64(Ya), float64(Xb), float64(Yb), float64(Px), float64(Py))
			totalCost += cost

			fmt.Printf("Cost: %d\nXa: %d, Ya: %d\nXb: %d, Yb: %d\nPx: %d, Py: %d\n", cost, Xa, Ya, Xb, Yb, Px, Py)

			Xa = 0
			Ya = 0
			Xb = 0
			Yb = 0
			Px = 0
			Py = 0
		}

	}
}

func run(file string) int {
	readInput(file)

	total := totalCost

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
