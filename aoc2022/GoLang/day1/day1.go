package day1

import (
	"AOC-2022-Go/aoc2022/GoLang/utils"
	"fmt"
	"sort"
	"strconv"
)

func PartOne() {
	elves := getCaloriesCarried()

	fmt.Println(max(elves))
}

func PartTwo() {
	elves := getCaloriesCarried()

	top := maxThree(elves)
	var total int
	for _, i := range top {
		total += i
	}

	fmt.Println(total)
}

func getCaloriesCarried() []int {
	input := utils.ReadLinesFromFile("resources/Day01.txt")
	var elves []int
	var current int

	for _, l := range input {
		if l == "" {
			elves = append(elves, current)
			current = 0
			continue
		}
		i, _ := strconv.Atoi(l)

		current += i
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
