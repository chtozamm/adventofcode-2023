package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func secondPart() (power int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var setsOfCubes []set

	for scanner.Scan() {
		/*
			Example of a line:
			Game 7: 3 red, 8 blue; 1 green, 9 blue; 4 red, 1 green, 3 blue
		*/
		line := scanner.Text()
		colon := strings.Index(line, ":")

		// Trim "Game {id}: "
		gameSets := strings.Split(line[colon+2:], "; ")

		var red, green, blue int

		for _, s := range gameSets {
			amountOfColors := strings.Split(s, ", ")

			for _, amountOfColor := range amountOfColors {
				valueColorPair := strings.Split(amountOfColor, " ")
				value, _ := strconv.Atoi(valueColorPair[0])
				color := valueColorPair[1]

				switch color {
				case "red":
					if value > red {
						red = value
					}
				case "green":
					if value > green {
						green = value
					}
				case "blue":
					if value > blue {
						blue = value
					}
				}
			}
		}
		setsOfCubes = append(setsOfCubes, set{red: red, green: green, blue: blue})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, set := range setsOfCubes {
		power += set.red * set.green * set.blue
	}

	return
}
