package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFirstDigit(t *testing.T) {
	assert.Equal(t, 1, getFirstDigit("1abc2", digitRegex))
	assert.Equal(t, 3, getFirstDigit("pqr3stu8vwx", digitRegex))
	assert.Equal(t, 1, getFirstDigit("a1b2c3d4e5f", digitRegex))
	assert.Equal(t, 7, getFirstDigit("treb7uchet", digitRegex))
	assert.Equal(t, 1, getFirstDigit("abcone2threexyz", digitRegexExtended))
	assert.Equal(t, 2, getFirstDigit("two1nine", digitRegexExtended))
	assert.Equal(t, 3, getFirstDigit("asdfthreewo1nine", digitRegexExtended))
	assert.Equal(t, 8, getFirstDigit("eightwothree", digitRegexExtended))
}

func TestGetLastDigit(t *testing.T) {
	assert.Equal(t, 2, getLastDigit("1abc2", digitRegex))
	assert.Equal(t, 8, getLastDigit("pqr3stu8vwx", digitRegex))
	assert.Equal(t, 5, getLastDigit("a1b2c3d4e5f", digitRegex))
	assert.Equal(t, 7, getLastDigit("treb7uchet", digitRegex))
	assert.Equal(t, 9, getLastDigit("two1nine", digitRegexExtended))
	assert.Equal(t, 1, getLastDigit("6twones", digitRegexExtended))
	assert.Equal(t, 1, getLastDigit("6twone", digitRegexExtended))
	assert.Equal(t, 1, getLastDigit("one", digitRegexExtended))
}

func TestGetCalibrationValue(t *testing.T) {
	assert.Equal(t, 12, getCalibrationValue("1abc2"))
	assert.Equal(t, 38, getCalibrationValue("pqr3stu8vwx"))
	assert.Equal(t, 15, getCalibrationValue("a1b2c3d4e5f"))
	assert.Equal(t, 77, getCalibrationValue("treb7uchet"))
}

func TestGetCalibrationValueExtended(t *testing.T) {
	assert.Equal(t, 29, getCalibrationValueExtended("two1nine"))
	assert.Equal(t, 83, getCalibrationValueExtended("eightwothree"))
	assert.Equal(t, 13, getCalibrationValueExtended("abcone2threexyz"))
	assert.Equal(t, 24, getCalibrationValueExtended("xtwone3four"))
	assert.Equal(t, 42, getCalibrationValueExtended("4nineeightseven2"))
	assert.Equal(t, 14, getCalibrationValueExtended("zoneight234"))
	assert.Equal(t, 76, getCalibrationValueExtended("7pqrstsixteen"))
	assert.Equal(t, 86, getCalibrationValueExtended("nldeightwoshgnsjnzmbkbxcxltsqtstrgdmvqvxbfour6six"))
	assert.Equal(t, 43, getCalibrationValueExtended("4sevenfpnmsqlhllrxrhjonesmgfhzmhvstwo3"))
	assert.Equal(t, 98, getCalibrationValueExtended("nine426six8zk"))
	assert.Equal(t, 67, getCalibrationValueExtended("rdfqcdrxdc6twotwo8fourthreeftrlzseven"))
	assert.Equal(t, 38, getCalibrationValueExtended("379eight"))
	assert.Equal(t, 61, getCalibrationValueExtended("6twones"))
}
