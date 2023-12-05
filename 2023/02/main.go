package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var bag = map[string]int{}

func init() {
	bag["red"] = 12
	bag["green"] = 13
	bag["blue"] = 14
}

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	gameIdSum := 0
	powerSum := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		index := strings.IndexByte(line, ':')
		gameId, _ := strconv.Atoi(line[5:index])
		game := line[index+2:]

		turns := strings.Split(game, "; ")
		isValid := checkGame(turns)
		if isValid {
			gameIdSum += gameId
		}

		powerSum += getPower(turns)
	}
	fmt.Println("part 1:", gameIdSum)
	fmt.Println("part 2:", powerSum)
}

func parseTurn(line string) map[string]int {
	game := map[string]int{}

	cubes := strings.Split(line, ", ")

	for _, cube := range cubes {
		split := strings.Fields(cube)
		val, _ := strconv.Atoi(split[0])
		color := split[1]
		game[color] = val
	}

	return game
}

func checkTurn(game map[string]int) bool {
	for color, num := range game {
		bagNum, found := bag[color]
		if !found {
			return false
		}
		if bagNum < num {
			return false
		}
	}
	return true
}

func checkGame(turns []string) bool {
	for _, rawTurn := range turns {
		turn := parseTurn(rawTurn)
		isValid := checkTurn(turn)
		if !isValid {
			return false
		}
	}
	return true
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func getPower(game []string) int {
	var maximums = map[string]int{}
	maximums["red"] = 0
	maximums["green"] = 0
	maximums["blue"] = 0

	for _, rawTurn := range game {
		turn := parseTurn(rawTurn)

		for color, num := range turn {
			maximums[color] = max(num, maximums[color])
		}
	}

	power := maximums["red"] * maximums["green"] * maximums["blue"]
	return power

}
