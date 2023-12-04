package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const digitRegex = `\d`
const digitRegexExtended = `(\d)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)`

var digitMap = map[string]int{}

func init() {
	digitMap["one"] = 1
	digitMap["two"] = 2
	digitMap["three"] = 3
	digitMap["four"] = 4
	digitMap["five"] = 5
	digitMap["six"] = 6
	digitMap["seven"] = 7
	digitMap["eight"] = 8
	digitMap["nine"] = 9
	digitMap["0"] = 0
	digitMap["1"] = 1
	digitMap["2"] = 2
	digitMap["3"] = 3
	digitMap["4"] = 4
	digitMap["5"] = 5
	digitMap["6"] = 6
	digitMap["7"] = 7
	digitMap["8"] = 8
	digitMap["9"] = 9
}

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	sumCalibrations := 0
	sumCalibrationsExtended := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		sumCalibrations += getCalibrationValue(line)
		sumCalibrationsExtended += getCalibrationValueExtended(line)
	}

	readFile.Close()

	fmt.Println(sumCalibrations)
	fmt.Println(sumCalibrationsExtended)
}

func getFirstDigit(line string, regex string) int {
	re := regexp.MustCompile(regex)
	match := re.FindString(line)
	digit := digitMap[match]
	return digit
}

// The regexp package does not provide the ability to find all overlapping
// matches. This is problematic when finding the last match in string. Therefore
// this function implements a simple means to perform matching on a widening
// window at the end of the string.
func getLastDigit(line string, regex string) int {
	re := regexp.MustCompile(regex)
	for i := len(line); i >= 0; i-- {
		slice := line[i:len(line)]
		match := re.FindAllString(slice, -1)
		if match != nil {
			digit := digitMap[match[len(match)-1]]
			return digit
		}
	}

	return 0
}

func getCalibrationValue(line string) int {
	calibration := getFirstDigit(line, digitRegex)*10 + getLastDigit(line, digitRegex)
	return calibration
}

func getCalibrationValueExtended(line string) int {
	calibration := getFirstDigit(line, digitRegexExtended)*10 + getLastDigit(line, digitRegexExtended)
	return calibration
}
