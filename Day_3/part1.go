package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type coordinate struct {
	row, column int
}

func firstPart() (sum int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Keep track of coordinates containing special characters
	specialCharactersMap := map[coordinate]bool{}

	// Keep track of coordinates and values of numbers
	numbersMap := map[coordinate]int{}

	// Keep track of the current row
	var row int

	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		var numStrBuilder strings.Builder
		for column, char := range chars {
			// Populate map of special characters
			if !isDigit(char) && char != "." {
				specialCharactersMap[coordinate{row, column}] = true
			}
			// Populate map of numbers
			numStr := numStrBuilder.String()
			if !isDigit(char) && len(numStr) != 0 {
				value, _ := strconv.Atoi(numStr)
				numbersMap[coordinate{row, column - len(numStr)}] = value
				numStrBuilder.Reset()
			}
			if isDigit(char) {
				numStrBuilder.WriteString(char)
			}
		}

		// Edge case: if the line ends with a number
		if numStrBuilder.Len() > 0 {
			value, _ := strconv.Atoi(numStrBuilder.String())
			numbersMap[coordinate{row, len(chars) - numStrBuilder.Len()}] = value
		}

		// Increment current row value
		row++
	}

	for c, value := range numbersMap {
		hasSymbol := false
		valLen := len(strconv.Itoa(value))

		// Check all surrounding coordinates for special characters
		for i := c.column; i < c.column+valLen; i++ {
			for _, d := range []coordinate{
				{c.row - 1, i - 1}, {c.row - 1, i}, {c.row - 1, i + 1},
				{c.row, i - 1}, {c.row, i + 1},
				{c.row + 1, i - 1}, {c.row + 1, i}, {c.row + 1, i + 1},
			} {
				if specialCharactersMap[d] {
					hasSymbol = true
					break
				}
			}
			if hasSymbol {
				break
			}
		}

		if hasSymbol {
			sum += value
		}
	}

	return
}

// isDigit returns true if the given character is a digit.
func isDigit(char string) bool {
	digitRegexp := regexp.MustCompile("[0-9]")
	return digitRegexp.MatchString(char)
}
