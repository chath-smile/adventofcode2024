package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	safe := 0
	list := loadData("input2.txt")

	for _, v := range list {
		safe += isSafe(v, true)
	}

	fmt.Println(safe)
}

func isSafe(report []int, useProblemDampener bool) int {
	if len(report) < 2 {
		return 1
	}

	prev := 0
	direction := report[0] < report[1]
	for i, v := range report {
		if i == 0 {
			continue
		}

		if unsafeDiff(report[prev], v) {
			if useProblemDampener {
				return problemDampener(report)
			}
			return 0
		}

		if !validDirection(report[prev], v, direction) {
			if useProblemDampener {
				return problemDampener(report)
			}
			return 0
		}

		prev = i
	}

	return 1
}

func problemDampener(report []int) int {
	for ignore := range report {
		dampened := remove(report, ignore)
		if isSafe(dampened, false) == 1 {
			return 1
		}
	}

	return 0
}

func remove(s []int, index int) []int {
	ret := make([]int, 0, len(s)-1)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

// Check if diff between 2 numbers is too high or too low.
func unsafeDiff(a int, b int) bool {
	return a-b > 3 || b-a > 3 || a-b == 0
}

func validDirection(a int, b int, increase bool) bool {
	if increase {
		return a < b
	} else {
		return a > b
	}
}

func loadData(filename string) [][]int {
	list := make([][]int, 0, 1000)

	file, err := os.Open(filename)
	if nil != err {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		// Convert line to slice of ints.
		report := make([]int, 0, len(line))
		for _, v := range line {
			if i, err := strconv.Atoi(v); err == nil {
				report = append(report, i)
			}
		}

		list = append(list, report)
	}

	if err := scanner.Err(); nil != err {
		log.Fatal(err)
	}

	return list
}
