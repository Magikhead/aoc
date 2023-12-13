package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCard(t *testing.T) {
	c := NewCard()
	c.parseString("Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53")
	assert.Equal(t, []int{41, 48, 83, 86, 17}, c.winningNumbers)
	assert.Equal(t, []int{83, 86, 6, 31, 17, 9, 48, 53}, c.gameNumbers)
	assert.Equal(t, 8, c.calcPoints())
}
