package day3

import (
	"AOC-2022-Go/aoc2022/GoLang/utils"
	"fmt"
	"github.com/ghetzel/go-stockutil/sliceutil"
	"strings"
)

func PartOne() {
	backpacks := readFile()
	sum := 0
	for _, backpack := range backpacks {
		if backpack != "" {
			commonChar := sliceutil.Intersect(backpack[:len(backpack)/2], backpack[len(backpack)/2:])
			sum += utils.GetCharValueFromAscii(commonChar[0].(uint8))
		}
	}
	fmt.Println(sum)

}

func PartTwo() {
	backpacks := readFile()
	sum := 0
	for i := 0; i < len(backpacks); i += 3 {
		for _, c := range backpacks[i] {
			if strings.Contains(backpacks[i+1], string(c)) && strings.Contains(backpacks[i+2], string(c)) {
				sum += utils.GetCharValueFromAscii(uint8(c))
				break
			}
		}
	}
	fmt.Println(sum)

}

func readFile() []string {
	return utils.ReadLines("aoc2022/GoLang/resources/Day03.txt")
}
