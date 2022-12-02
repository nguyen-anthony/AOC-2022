package day1

import (
	"AOC-2022-Go/aoc2022/GoLang/utils"
	"fmt"
	"sort"
	"strconv"
)

func PartOne() {
	elves := getCaloriesSlice()

	fmt.Println(max(elves))
}

func PartTwo() {
	elves := getCaloriesSlice()

	top := maxThree(elves)
	var total int
	for _, i := range top {
		total += i
	}

	fmt.Println(total)
}

func getCaloriesSlice() []int {
	input := utils.ReadLinesFromFile("aoc2022/GoLang/resources/Day01.txt")
	var elves []int
	var sum int

	for _, cal := range input {
		if cal == "" {
			elves = append(elves, sum)
			sum = 0
			continue
		}
		calInt, _ := strconv.Atoi(cal)
		sum += calInt
	}
	return elves
}

func max(input []int) int {
	sort.Ints(input)
	return input[len(input)-1]
}

func maxThree(input []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(input[:])))
	return input[:3]
}
