package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	ONE   = "one"
	TWO   = "two"
	THREE = "three"
	FOUR  = "four"
	FIVE  = "five"
	SIX   = "six"
	SEVEN = "seven"
	EIGHT = "eight"
	NINE  = "nine"
)

func secondPart() (sum uint) {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		var firstDigit uint
		var lastDigit uint

		for idx, char := range line {

			if unicode.IsDigit(char) {
				digit, _ := strconv.Atoi(string(rune(char)))

				if firstDigit == 0 {
					firstDigit = uint(digit)
				}
				lastDigit = uint(digit)
				continue
			}

			// Get the line's tail from the current index
			s := line[idx:]

			if hasDigit, digitLength := startsWithSpelledDigit(s); hasDigit {
				spelledDigit := s[:digitLength] // Cut the tail to get pure spelled digit
				digit := spelledDigitToNum(spelledDigit)

				if firstDigit == 0 {
					firstDigit = digit
				}
				lastDigit = digit
			}
		}

		// Add two-digit number to the sum
		sum += firstDigit*10 + lastDigit
	}

	return
}

// startsWithSpelledDigit returns true if the given string starts with a spelled digit (from 1 to 9).
// Returns the length of the spelled digit string as the second value.
func startsWithSpelledDigit(s string) (response bool, spelledDigitLength int) {
	digits := []string{ONE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE}

	for _, digit := range digits {
		if strings.HasPrefix(s, digit) {
			return true, len(digit)
		}
	}

	return false, 0
}

// spelledDigitToNum returns numerical representation of a spelled digit.
// Returns 0 if the given string is not a digit.
func spelledDigitToNum(s string) uint {
	switch strings.ToLower(s) {
	case ONE:
		return 1
	case TWO:
		return 2
	case THREE:
		return 3
	case FOUR:
		return 4
	case FIVE:
		return 5
	case SIX:
		return 6
	case SEVEN:
		return 7
	case EIGHT:
		return 8
	case NINE:
		return 9
	default:
		return 0
	}
}
