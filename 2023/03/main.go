package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func countDigits(i int) int {
	return len(strconv.Itoa(i))
}

type Engine struct {
	schematic  [][]rune
	partNumber int
	length     int
	posX       int
	posY       int
	width      int
	height     int
}

func NewEngine(input string) *Engine {
	e := Engine{}
	e.posX = -1
	e.posY = 0
	e.partNumber = 0

	readFile, _ := os.Open(input)
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		runes := []rune(fileScanner.Text())
		e.schematic = append(e.schematic, runes)
	}

	e.height = len(e.schematic)
	e.width = len(e.schematic[0])

	return &e
}

func (e *Engine) print() {
	for r, _ := range e.schematic {
		for c, _ := range e.schematic[r] {
			fmt.Printf(string(e.schematic[r][c]))
		}
		fmt.Printf("\n")
	}
}

func (e *Engine) get(x int, y int) (rune, error) {
	if (x < 0) || (x >= e.width) || (y < 0) || (y >= e.height) {
		return '_', errors.New("out of bounds")
	}
	return e.schematic[y][x], nil
}

func (e *Engine) getPos() rune {
	val, _ := e.get(e.posX, e.posY)
	return val
}

func (e *Engine) readPartNumber(x int, y int) {
	var sb strings.Builder
	for ; x < len(e.schematic[y]); x++ {
		val, _ := e.get(x, y)
		if unicode.IsDigit(val) {
			sb.WriteRune(val)
		} else {
			break
		}

	}
	e.length = len(sb.String())
	e.partNumber, _ = strconv.Atoi(sb.String())
}

func isSymbol(r rune) bool {
	if unicode.IsDigit(r) {
		return false
	} else if r == '.' {
		return false
	}
	return true
}

func (e *Engine) checkNeighbor(x int, y int) bool {
	val, err := e.get(x, y)
	if (err == nil) && isSymbol(val) {
		return true
	}

	return false
}

func (e *Engine) addNeighbor(x int, y int, neighbors *map[int]bool) {
	number, err := e.readNumber(x, y)
	if err == nil {
		(*neighbors)[number] = true
	}
}

// get a list of numbers that neigbor a coordinate
func (e *Engine) getNeighbors(x int, y int) []int {
	neighbors := map[int]bool{}

	e.addNeighbor(x-1, y-1, &neighbors)
	e.addNeighbor(x, y-1, &neighbors)
	e.addNeighbor(x+1, y-1, &neighbors)
	e.addNeighbor(x+1, y, &neighbors)
	e.addNeighbor(x+1, y+1, &neighbors)
	e.addNeighbor(x, y+1, &neighbors)
	e.addNeighbor(x-1, y+1, &neighbors)
	e.addNeighbor(x-1, y, &neighbors)

	numbers := []int{}
	for k, _ := range neighbors {
		numbers = append(numbers, k)
	}

	return numbers
}

func (e *Engine) checkNeighbors(x int, y int) bool {
	if e.checkNeighbor(x-1, y-1) {
		return true
	} else if e.checkNeighbor(x, y-1) {
		return true
	} else if e.checkNeighbor(x+1, y-1) {
		return true
	} else if e.checkNeighbor(x+1, y) {
		return true
	} else if e.checkNeighbor(x+1, y+1) {
		return true
	} else if e.checkNeighbor(x, y+1) {
		return true
	} else if e.checkNeighbor(x-1, y+1) {
		return true
	} else if e.checkNeighbor(x-1, y) {
		return true
	}

	return false
}

func (e *Engine) isPartNumber() bool {
	x := e.posX
	y := e.posY

	for i := 0; i < e.length; i++ {
		if e.checkNeighbors(x+i, y) {
			return true
		}
	}

	return false
}

func (e *Engine) readNumber(x int, y int) (int, error) {

	start, _ := e.get(x, y)
	if !unicode.IsDigit(start) {
		return -1, errors.New("not a number")
	}

	// move to beginning of number
	atBeginning := false
	for !atBeginning {
		left, err := e.get(x-1, y)
		if err != nil {
			atBeginning = true
		} else if !unicode.IsDigit(left) {
			atBeginning = true
		} else {
			x -= 1
		}
	}

	// once cursor is at the beginning read the full number
	e.readPartNumber(x, y)
	return e.partNumber, nil
}

func (e *Engine) findNextNumber() (int, int, error) {
	// start by advancing the pointer to the end of the current part number
	e.posX += countDigits(e.partNumber)

	for ; e.posY < len(e.schematic); e.posY++ {
		for ; e.posX < len(e.schematic[e.posY]); e.posX++ {
			val := e.getPos()

			if unicode.IsDigit(val) {
				e.readPartNumber(e.posX, e.posY)
				return e.posX, e.posY, nil
			}
		}

		// end of line, reset to beggining
		e.posX = 0
	}
	return 0, -1, errors.New("not found")
}

func (e *Engine) findNextGear() (int, int, error) {
	// start by advancing the pointer past the current gear
	e.posX += 1

	for ; e.posY < len(e.schematic); e.posY++ {
		for ; e.posX < len(e.schematic[e.posY]); e.posX++ {
			val := e.getPos()

			if val == '*' {
				return e.posX, e.posY, nil
			}
		}

		// end of line, reset to beggining
		e.posX = 0
	}
	return 0, -1, errors.New("not found")
}

func findSum(e *Engine) int {
	sum := 0
	for {
		_, _, err := e.findNextNumber()
		if err != nil {
			break
		}
		if e.isPartNumber() {
			sum += e.partNumber
		}

	}

	return sum
}

func (e *Engine) reset() {
	e.posX = -1
	e.posY = 0
}

func findGearSum(e *Engine) int {
	sum := 0
	for {
		x, y, err := e.findNextGear()

		// no more gears found
		if err != nil {
			break
		}

		neighbors := e.getNeighbors(x, y)
		if len(neighbors) == 2 {
			gearRatio := neighbors[0] * neighbors[1]
			sum += gearRatio
		}

	}

	return sum
}

func main() {
	e := NewEngine("input.txt")
	sum := findSum(e)
	fmt.Println("sum:", sum)

	e.reset()
	gearSum := findGearSum(e)
	fmt.Println("gear sum:", gearSum)
}
