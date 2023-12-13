package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	number         int
	winningNumbers []int
	gameNumbers    []int
}

// sample input string
//
//	Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func NewCard() *Card {
	c := Card{}
	c.number = 0

	return &c
}

func (c *Card) parseString(input string) {
	colonIndex := strings.Index(input, ":")
	c.number, _ = strconv.Atoi(strings.TrimSpace(input[4:colonIndex]))

	input = input[colonIndex+1:]

	pipeIndex := strings.Index(input, "|")

	winningNumberFields := strings.Fields(input[:pipeIndex])
	for _, n := range winningNumberFields {
		val, _ := strconv.Atoi(n)
		c.winningNumbers = append(c.winningNumbers, val)
	}

	gameNumberFields := strings.Fields(input[pipeIndex+1:])
	for _, n := range gameNumberFields {
		val, _ := strconv.Atoi(n)
		c.gameNumbers = append(c.gameNumbers, val)
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (c *Card) countWinners() int {
	matches := 0
	for _, n := range c.winningNumbers {
		if contains(c.gameNumbers, n) {
			matches = matches + 1
		}
	}

	if matches == 0 {
		return 0
	}

	return matches
}
func (c *Card) calcPoints() int {
	return int(math.Pow(2, float64(c.countWinners()-1)))
}

func main() {
	input := "input.txt"

	cardMap := map[int]*Card{}

	readFile, _ := os.Open(input)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		card := NewCard()
		card.parseString(fileScanner.Text())
		cardMap[card.number] = card
		sum = sum + card.calcPoints()
	}

	fmt.Println("sum:", sum)

	cardCount := map[int]int{}

	// initialize the count to a single copy of each card
	for number, _ := range cardMap {
		cardCount[number] = 1
	}

	fmt.Println(cardCount)
	for number := 1; number <= len(cardMap); number++ {
		c := cardMap[number]
		numWinners := c.countWinners()

		count := cardCount[number]

		for i := 1; i <= numWinners; i++ {
			index := number + i
			if index > len(cardCount) {
				break
			}
			cardCount[index] = cardCount[index] + count
		}
	}

	total := 0
	for _, c := range cardCount {
		total += c
	}

	fmt.Println("total:", total)

}
