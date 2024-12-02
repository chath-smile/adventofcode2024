package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	score := 0
	list1, list2 := loadData()
	sort.Ints(list1)
	sort.Ints(list2)

	for _, val1 := range list1 {
		multiplier := 0
		for _, val2 := range list2 {
			if val1 == val2 {
				multiplier++
			}
			if val2 > val1 {
				break
			}
		}
		score += val1 * multiplier
	}

	fmt.Println(score)
}

func loadData() ([]int, []int) {
	list1 := make([]int, 1000)
	list2 := make([]int, 1000)

	file, err := os.Open("input.txt")
	if nil != err {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if i, err := strconv.Atoi(line[0]); err == nil {
			list1 = append(list1, i)
		}

		if i, err := strconv.Atoi(line[1]); err == nil {
			list2 = append(list2, i)
		}
	}

	if err := scanner.Err(); nil != err {
		log.Fatal(err)
	}

	return list1, list2
}
