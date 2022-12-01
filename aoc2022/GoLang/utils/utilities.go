package utils

import (
	"bufio"
	"log"
	"os"
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
