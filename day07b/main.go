package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	value   int
	numbers []int
}

var equations []equation

func Atoi(n string) int {
	if i, err := strconv.Atoi(n); err != nil {
		return 0
	} else {
		return i
	}
}

func (e equation) String() string {
	return fmt.Sprintf("%d: %v\n", e.value, e.numbers)
}

func getOp(i int, o string) string {
	// fmt.Printf("[getOpt %b]", (o>>(i))&1)
	return string(o[i])
}

func GeneratePermutations(N int) []string {
	if N <= 0 {
		return []string{}
	}

	// chars := []rune{'+', '*'}
	chars := []rune{'+', '*', '|'}
	result := []string{}
	var backtrack func(current string)

	backtrack = func(current string) {
		if len(current) == N {
			result = append(result, current)
			return
		}
		for _, char := range chars {
			backtrack(current + string(char))
		}
	}

	backtrack("")
	return result
}

func printequation(e equation, o string) int {
	total := e.numbers[0]
	for i, v := range e.numbers {
		fmt.Print(v)

		if i < len(e.numbers)-1 {
			fmt.Printf(getOp(i, o))
		}
	}

	for i := 1; i < len(e.numbers); i++ {
		if getOp(i-1, o) == "+" {
			total += e.numbers[i]
			// fmt.Printf("[+%d = total %d]", e.numbers[i], total)
		}
		if getOp(i-1, o) == "*" {
			total = total * e.numbers[i]
			// fmt.Printf("[*%d = total %d]", e.numbers[i], total)
		}

	}

	fmt.Printf(" = %d    [%v]", total, e.value == total)

	fmt.Println()
	return total
}

func calculate(e equation) bool {

	operations := GeneratePermutations(len(e.numbers) - 1)

	for i, v := range operations {
		result := printequation(e, v)

		fmt.Println(i, v)

		if e.value == result {
			return true
		}
	}
	return false
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

		e := equation{}

		parts := strings.Split(line, ":")
		e.value = Atoi(parts[0])

		numbers := strings.Split(parts[1][1:], " ")

		for _, v := range numbers {
			e.numbers = append(e.numbers, Atoi(v))
		}
		equations = append(equations, e)
	}
}

func run(file string) int {
	readInput(file)

	total := 0

	fmt.Println(equations)

	for _, v := range equations {
		if calculate(v) {
			total += v.value
		}
	}

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
