package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadLinesFromFile(file string) []string {
	content, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(content)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ReadLines(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	return lines
}

func GetCharValueFromAscii(i uint8) int {
	if i >= 65 && i <= 90 {
		return int(i%32 + 26)
	} else {
		return int(i % 32)
	}
}
