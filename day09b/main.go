package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var files []int

func Atoi(n string) int {
	if i, err := strconv.Atoi(n); err != nil {
		return 0
	} else {
		return i
	}
}

func displayMap() {
	for _, v := range files {
		if v == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println()
}

func readInput(file string) {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal("Can't open file: " + err.Error())
	}

	s := bufio.NewScanner(f)

	s.Split(bufio.ScanLines)

	index := 0

	for s.Scan() {
		line := s.Text()

		for i := 0; i < len(line); i += 2 {
			fileSize := string(line[i])

			for range Atoi(fileSize) {
				files = append(files, index)
			}
			index++

			if i+1 < len(line) {
				freeSize := string(line[i+1])
				for range Atoi(freeSize) {
					files = append(files, -1)
				}
			}
		}
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

func findFree(max int) int {
	for i := range files {
		// check if we're at max, i.e. position from where we want to relocate
		// It means no more free and most likely we're done
		if i == max {
			return -1
		}

		if files[i] == -1 {
			return i
		}
	}
	return -1
}

func optimize() {
	for i := len(files) - 1; i > -1; i-- {
		free := findFree(i)
		fmt.Printf("Pos: %d, Id: %d, Free: %d\n", i, files[i], free)

		if free == -1 {
			break
		} else {
			// Relocate ...
			files[free] = files[i]
			// .. then free
			files[i] = -1
		}
	}
}

func getCksum() int {
	var cksum int

	for i, v := range files {
		// check if we're at max, i.e. position from where we want to relocate
		// It means no more free and most likely we're done

		if files[i] != -1 {
			cksum += (i * v)
		}
	}
	return cksum
}

func run(file string) int {
	readInput(file)

	displayMap()

	optimize()
	displayMap()
	total := getCksum()

	return total
}

func main() {
	total := run("input.txt")

	fmt.Println("Total: ", total)
}
