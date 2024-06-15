package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func firstPart() (sum uint) {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		var firstDigit uint
		var lastDigit uint

		for _, char := range line {
			if unicode.IsDigit(char) {
				digit, _ := strconv.Atoi(string(rune(char)))

				if firstDigit == 0 {
					firstDigit = uint(digit)
				}
				lastDigit = uint(digit)
			}
		}

		sum += firstDigit*10 + lastDigit
	}

	return
}
