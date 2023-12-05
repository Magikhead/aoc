package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTurn(t *testing.T) {
	assert.Equal(t, map[string]int{"blue": 2, "red": 3, "green": 6}, parseTurn("2 blue, 3 red, 6 green"))
	assert.Equal(t, map[string]int{"green": 2}, parseTurn("2 green"))
}

func TestCheckTurn(t *testing.T) {
	assert.Equal(t, true, checkTurn(map[string]int{"blue": 2, "red": 3, "green": 6}))
}

func TestCheckGame(t *testing.T) {
	assert.Equal(t, true, checkGame([]string{"2 green", "2 blue"}))
	assert.Equal(t, true, checkGame([]string{"6 red, 1 blue, 3 green", "2 blue, 1 red, 2 green"}))
}

func TestGetPower(t *testing.T) {
	assert.Equal(t, 48, getPower([]string{"3 blue, 4 red", "1 red, 2 green, 6 blue", "2 green"}))
}
