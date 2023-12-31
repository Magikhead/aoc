package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := "input.txt"
	readFile, _ := os.Open(input)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	turnsStr := fileScanner.Text()

	fmt.Println(turnsStr)

	// blank line
	fileScanner.Scan()

	leftMap := map[string]string{}
	rightMap := map[string]string{}

	for fileScanner.Scan() {
		entry := fileScanner.Text()
		key := entry[0:3]
		left := entry[7:10]
		right := entry[12:15]

		leftMap[key] = left
		rightMap[key] = right
	}

	// find starting locations
	locations := map[string]int{}

	for k, _ := range leftMap {
		if k[2] == 'A' {
			locations[k] = 0
		}
	}

	fmt.Println(locations)
	for start, _ := range locations {
		numTurns := findEnd(turnsStr, leftMap, rightMap, start)
		locations[start] = numTurns
		fmt.Println(start, numTurns)
	}

	// answer is LCM of all the above
}

func findEnd(turnsStr string, leftMap map[string]string, rightMap map[string]string, start string) int {

	index := 0
	location := start

	for {
		turn := turnsStr[index % len(turnsStr)]
		index += 1
		if turn == 'L' {
			location = leftMap[location]
		} else if turn == 'R' {
			location = rightMap[location]
		}

		if location[2] == 'Z' {
			return index
		}
	}
}
