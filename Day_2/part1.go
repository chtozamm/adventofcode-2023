package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	RED   = 12
	GREEN = 13
	BLUE  = 14
)

type set struct {
	red, green, blue int
}

func firstPart() (sum int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		/*
			Example of a line:
			Game 7: 3 red, 8 blue; 1 green, 9 blue; 4 red, 1 green, 3 blue
		*/
		line := scanner.Text()
		colon := strings.Index(line, ":")
		gameId, _ := strconv.Atoi(line[len("Game "):colon])

		// Trim "Game {id}: "
		gameSets := strings.Split(line[colon+2:], "; ")

		var setsOfCubes []set

		for _, s := range gameSets {
			var red, green, blue int
			amountOfColors := strings.Split(s, ", ")

			for _, amountOfColor := range amountOfColors {
				valueColorPair := strings.Split(amountOfColor, " ")
				value, _ := strconv.Atoi(valueColorPair[0])
				color := valueColorPair[1]

				switch color {
				case "red":
					red += value
				case "green":
					green += value
				case "blue":
					blue += value
				}
			}
			setsOfCubes = append(setsOfCubes, set{red: red, green: green, blue: blue})
		}
		if isGamePossible(setsOfCubes) {
			sum += gameId
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

// isGamePossible returns true if each of the game sets is possible.
func isGamePossible(sets []set) bool {
	for _, s := range sets {
		if !isSetPossible(s) {
			return false
		}
	}

	return true
}

// isSetPossible returns true if the amounts of cubes in the given set don't exceed
// the amounts of cubes defined as constants in the beginning of the file.
func isSetPossible(s set) bool {
	if s.red <= RED && s.green <= GREEN && s.blue <= BLUE {
		return true
	}

	return false
}
