package day2

import (
	"AOC-2022-Go/aoc2022/GoLang/utils"
	"fmt"
)

func PartOne() {
	score, _ := scoring()
	fmt.Println(score)
}

func PartTwo() {
	_, score := scoring()
	fmt.Println(score)
}

func scoring() (int, int) {
	lines := utils.ReadLines("aoc2022/GoLang/resources/Day02.txt")

	scorePart1, scorePart2 := 0, 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		opponent := line[0] - 'A'
		myChoice := line[2] - 'X'

		if (myChoice+1)%3 == opponent {
			scorePart1 += int(myChoice) + 1
		} else if myChoice == opponent {
			scorePart1 += 3 + int(myChoice) + 1
		} else {
			scorePart1 += 6 + int(myChoice) + 1
		}

		if myChoice == 0 {
			scorePart2 += (int(opponent)+2)%3 + 1
		} else if myChoice == 1 {
			scorePart2 += 3 + int(opponent) + 1
		} else {
			scorePart2 += 6 + (int(opponent)+1)%3 + 1
		}
	}
	return scorePart1, scorePart2
}
