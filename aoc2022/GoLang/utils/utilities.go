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