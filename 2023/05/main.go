package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeEntry struct {
	source int
	dest   int
	length int
}

func NewRangeEntry(source int, dest int, length int) *RangeEntry {
	entry := RangeEntry{source, dest, length}
	return &entry
}

// check if the input falls within the source range
func (entry *RangeEntry) isInRange(input int) bool {
	end := entry.source + entry.length
	if input < entry.source {
		return false
	} else if input < end {
		return true
	}

	return false
}

func (entry *RangeEntry) lookup(input int) (int, error) {
	if false == entry.isInRange(input) {
		return 0, errors.New("out of bounds")
	}

	position := input - entry.source
	mapping := entry.dest + position
	return mapping, nil
}

func searchMap(input int, searchMap []*RangeEntry) (int, error) {
	for _, entry := range searchMap {
		val, err := entry.lookup(input)
		if err == nil {
			return val, nil
		}
	}

	message := fmt.Sprintf("value [%d] not found", input)
	return input, errors.New(message)

}

func parseHeader(input string) (string, string) {
	words := strings.Fields(input)
	names := strings.Split(words[0], "-to-")

	source := names[0]
	dest := names[1]

	return source, dest
}

func main() {
	input := "input.txt"

	maps := make(map[string][]*RangeEntry)

	readFile, _ := os.Open(input)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()

	seedStr := strings.Fields(fileScanner.Text()[7:])
	seeds := []int{}

	for _, s := range seedStr {
		seedInt, _ := strconv.Atoi(s)
		seeds = append(seeds, seedInt)
	}

	// scan blank line
	fileScanner.Scan()

	for {
		header, mapping := readMap(fileScanner)
		if len(header) == 0 {
			break
		}
		maps[header] = mapping
	}

	{
		values := []int{}

		for _, seed := range seeds {
			val, err := searchMap(seed, maps["seed-to-soil"])
			if err != nil {
				//fmt.Println("seed-to-soil",err)
			}

			val, err = searchMap(val, maps["soil-to-fertilizer"])
			if err != nil {
				//fmt.Println("soil-to-fertilizer",err)
			}
			val, err = searchMap(val, maps["fertilizer-to-water"])
			if err != nil {
				//fmt.Println("fertilizer-to-water",err)
			}
			val, err = searchMap(val, maps["water-to-light"])
			if err != nil {
				//fmt.Println("water-to-light", err)
			}
			val, err = searchMap(val, maps["light-to-temperature"])
			if err != nil {
				//fmt.Println("light-to-temperature",err)
			}
			val, err = searchMap(val, maps["temperature-to-humidity"])
			if err != nil {
				//fmt.Println("termperature-to-humidity",err)
			}

			val, err = searchMap(val, maps["humidity-to-location"])
			if err != nil {
				//fmt.Println("humidity-to-location",err)
			}

			values = append(values, val)
		}

		min := values[0]

		for _, val := range values {
			if val < min {
				min = val
			}
		}

		fmt.Println("min:", min)
	}

	// build seedRange list of all seeds
	seedRange := []int{}
	for i := 0; i < len(seeds); i += 2 {
		begin := seeds[i]
		length := seeds[i+1]

		for s := 0; s < length; s++ {
			seedRange = append(seedRange, begin+s)
		}
	}

	fmt.Println("working...")
	{

		min := 99999999
		for _, seed := range seedRange {
			val, err := searchMap(seed, maps["seed-to-soil"])
			if err != nil {
				//continue
			}

			val, err = searchMap(val, maps["soil-to-fertilizer"])
			if err != nil {
				//continue
			}
			val, err = searchMap(val, maps["fertilizer-to-water"])
			if err != nil {
				//continue
			}
			val, err = searchMap(val, maps["water-to-light"])
			if err != nil {
				//continue
			}
			val, err = searchMap(val, maps["light-to-temperature"])
			if err != nil {
				//continue
			}
			val, err = searchMap(val, maps["temperature-to-humidity"])
			if err != nil {
				//continue
			}

			val, err = searchMap(val, maps["humidity-to-location"])
			if err != nil {
				//continue
			}

			if val < min {
				min = val
				fmt.Println("min:", min)
			}
		}

		fmt.Println("min:", min)
	}
}

func readMap(fileScanner *bufio.Scanner) (string, []*RangeEntry) {
	// read the header text
	if !fileScanner.Scan() {
		return "", nil
	}

	header := strings.Fields(fileScanner.Text())[0]
	//source, dest := parseHeader(fileScanner.Text())
	//associations[source] = dest

	var mapping []*RangeEntry

	// read map entries
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			break
		}

		fields := strings.Fields(line)
		source, _ := strconv.Atoi(fields[1])
		dest, _ := strconv.Atoi(fields[0])
		r, _ := strconv.Atoi(fields[2])

		entry := NewRangeEntry(source, dest, r)
		mapping = append(mapping, entry)
	}

	return header, mapping
}

func findMinimumLocation([]string) int {
	return 0
}

//func readMap([]string)
