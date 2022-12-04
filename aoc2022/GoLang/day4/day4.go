package day4

import (
	"AOC-2022-Go/aoc2022/GoLang/utils"
	"fmt"
	"strconv"
	"strings"
)

func PartOne() {
	assignments := readFile()
	fullContains := 0
	for _, assignment := range assignments {
		elves := strings.Split(assignment, ",")
		elf1, elf2 := elves[0], elves[1]
		range1 := strings.Split(elf1, "-")
		range2 := strings.Split(elf2, "-")
		min1str, max1str := range1[0], range1[1]
		min2str, max2str := range2[0], range2[1]

		min1, _ := strconv.Atoi(min1str)
		max1, _ := strconv.Atoi(max1str)
		min2, _ := strconv.Atoi(min2str)
		max2, _ := strconv.Atoi(max2str)

		if min1 <= min2 && max1 >= max2 || min2 <= min1 && max2 >= max1 {
			fullContains++
		}
	}
	fmt.Println(fullContains)
}

func PartTwo() {
	assignments := readFile()
	allContains := 0
	for _, assignment := range assignments {
		elves := strings.Split(assignment, ",")
		elf1, elf2 := elves[0], elves[1]
		range1 := strings.Split(elf1, "-")
		range2 := strings.Split(elf2, "-")
		min1str, max1str := range1[0], range1[1]
		min2str, max2str := range2[0], range2[1]

		min1, _ := strconv.Atoi(min1str)
		max1, _ := strconv.Atoi(max1str)
		min2, _ := strconv.Atoi(min2str)
		max2, _ := strconv.Atoi(max2str)

		if (min1 >= min2 && min1 <= max2) ||
			(max1 >= min2 && max1 <= max2) ||
			(min2 >= min1 && min2 <= max1) ||
			(max2 >= min1 && max2 <= max1) {
			allContains++
		}
	}
	fmt.Println(allContains)
}

func readFile() []string {
	return utils.ReadLines("aoc2022/GoLang/resources/Day04.txt")
}
