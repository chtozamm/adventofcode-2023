package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Part2() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	var sum uint = 0

	for _, line := range lines {
		var firstDigit uint = 00
		var lastDigit uint = 00

		for idx, char := range line {
			if unicode.IsDigit(char) && firstDigit == 00 {
				digit, _ := strconv.Atoi(string(char))
				firstDigit = uint(digit)
			}

			if unicode.IsDigit(char) {
				digit, _ := strconv.Atoi(string(char))
				lastDigit = uint(digit)
			}

			// Check for spelled digits
			s := line[idx:]

			if strings.HasPrefix(s, "one") {
				if firstDigit == 00 {
					firstDigit = 1
				}
				lastDigit = 1
				continue
			}

			if strings.HasPrefix(s, "two") {
				if firstDigit == 00 {
					firstDigit = 2
				}
				lastDigit = 2
				continue
			}

			if strings.HasPrefix(s, "three") {
				if firstDigit == 00 {
					firstDigit = 3
				}
				lastDigit = 3
				continue
			}

			if strings.HasPrefix(s, "four") {
				if firstDigit == 00 {
					firstDigit = 4
				}
				lastDigit = 4
				continue
			}

			if strings.HasPrefix(s, "five") {
				if firstDigit == 00 {
					firstDigit = 5
				}
				lastDigit = 5
				continue
			}

			if strings.HasPrefix(s, "six") {
				if firstDigit == 00 {
					firstDigit = 6
				}
				lastDigit = 6
				continue
			}

			if strings.HasPrefix(s, "seven") {
				if firstDigit == 00 {
					firstDigit = 7
				}
				lastDigit = 7
				continue
			}

			if strings.HasPrefix(s, "eight") {
				if firstDigit == 00 {
					firstDigit = 8
				}
				lastDigit = 8
				continue
			}

			if strings.HasPrefix(s, "nine") {
				if firstDigit == 00 {
					firstDigit = 9
				}
				lastDigit = 9
				continue
			}
		}

		sum += firstDigit*10 + lastDigit
	}

	fmt.Printf("Part 2 > Sum of calibration values: %v\n", sum)
}
