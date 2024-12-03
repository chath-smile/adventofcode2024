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
	list := loadData("input2.txt")

	safe := 0
	for _, v := range list {
		safe += checkReport(v)
	}

	fmt.Println(safe)
}

func checkReport(report []int) int {
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
			return 0
		}

		if !validDirection(report[prev], v, direction) {
			return 0
		}

		prev = i
	}

	return 1
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
