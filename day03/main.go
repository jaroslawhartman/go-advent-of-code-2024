package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func getMulResult(m string) int {
	var a, b int

	fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
	return a * b
}

func scanMuls(data []byte, atEOF bool) (advance int, token []byte, err error) {
	i := 0
	advance = 0
	token = nil
	err = nil

	if atEOF {
		return
	}
	fmt.Println(">> START >>>", string(data[:]))
	for {

		if i > len(data)-4 {
			// err = bufio.ErrFinalToken
			return
		}

		if data[i+0] == 'm' && data[i+1] == 'u' && data[i+2] == 'l' && data[i+3] == '(' {
			advance = i + 4
			token = []byte("mul(")

			fmt.Println(">> Found >>>>>>>>>", i, advance, string(data[advance]), string(token[:]))

			for data[advance] >= '0' && data[advance] <= '9' {
				token = append(token, data[advance])
				advance++
				fmt.Println(">> Found digit >>>>>>>>>", i, advance, string(data[advance]), string(token[:]))
			}

			if data[advance] == ',' {
				token = append(token, data[advance])
				advance++
				fmt.Println(">> Found comma (,) >>>>>>>>>", i, advance, string(data[advance]), string(token[:]))
			} //else {

			// 	token = nil
			// 	// fmt.Println(">> Breaking >>>>>>>>>", i, advance, string(data[advance]), string(token[:]))
			// 	break
			// }

			for data[advance] >= '0' && data[advance] <= '9' {
				token = append(token, data[advance])
				fmt.Println(">> Found digit >>>>>>>>>", i, advance, string(data[advance]), string(token[:]))
				advance++
			}

			if data[advance] == ')' {
				token = append(token, data[advance])
				fmt.Println(">> Found Final >>>>>>>>>", i, advance, string(data[advance]), string(token[:]))
				return
			}

		}
		fmt.Println(">> Looping >>>>>>", i, string(data[i]))
		i++
	}
}

func readInput(file string) int {
	total := 0
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Can't open file: " + err.Error())
	}

	s := bufio.NewScanner(f)
	const maxCapacity = 100000
	buf := make([]byte, maxCapacity)
	s.Buffer(buf, maxCapacity)
	s.Split(scanMuls)

	for s.Scan() {
		ch := s.Text()
		fmt.Println(ch, getMulResult(ch))
		total += getMulResult(ch)
	}
	return total
}

func run(file string) int {
	return readInput(file)
}

func main() {

	total := run("input.txt")

	fmt.Println("Total: ", total)
}
