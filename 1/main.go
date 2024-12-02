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
	diff := 0
	list1, list2 := loadData()
	sort.Ints(list1)
	sort.Ints(list2)

	for i, v := range list1 {
		diff += abs(v - list2[i])
	}

	fmt.Println(diff)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
