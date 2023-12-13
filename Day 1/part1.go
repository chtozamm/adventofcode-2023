package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Part1() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	var sum uint = 0

	for _, line := range lines {
		var firstDigit uint = 00
		var lastDigit uint = 00

		for _, char := range line {
			if unicode.IsDigit(char) && firstDigit == 00 {
				digit, _ := strconv.Atoi(string(char))
				firstDigit = uint(digit)
			}

			if unicode.IsDigit(char) {
				digit, _ := strconv.Atoi(string(char))
				lastDigit = uint(digit)
			}
		}

		sum += firstDigit*10 + lastDigit
	}
	fmt.Printf("Part 1 > Sum of calibration values: %v\n", sum)
}
