package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSum(t *testing.T) {
	e := NewEngine("sample.txt")
	assert.Equal(t, 4361, findSum(e))
}

func TestGet(t *testing.T) {
	e := NewEngine("sample.txt")

	number, err := e.get(1, 0)
	assert.Equal(t, nil, err)
	assert.Equal(t, '6', number)
}

func TestReadNumber(t *testing.T) {
	e := NewEngine("sample.txt")

	number, err := e.readNumber(1, 0)
	assert.Equal(t, nil, err)
	assert.Equal(t, 467, number)
}
