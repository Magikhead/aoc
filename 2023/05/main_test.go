package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsInRange(t *testing.T) {
	{
		entry := NewRangeEntry(0, 0, 10)
		assert.Equal(t, true, entry.isInRange(5))
		assert.Equal(t, false, entry.isInRange(-1))
		assert.Equal(t, false, entry.isInRange(10))
		assert.Equal(t, true, entry.isInRange(9))
	}

	{
		entry := NewRangeEntry(50, 98, 2)
		assert.Equal(t, true, entry.isInRange(50))
	}
}

func TestLookup(t *testing.T) {
	entry := NewRangeEntry(98, 50, 2)

	{
		mapping, err := entry.lookup(98)
		assert.Equal(t, nil, err)
		assert.Equal(t, 50, mapping)
	}

	{
		mapping, err := entry.lookup(99)
		assert.Equal(t, nil, err)
		assert.Equal(t, 51, mapping)
	}

	{
		_, err := entry.lookup(100)
		assert.NotEmpty(t, err)
	}
}

func TestParseHeader(t *testing.T) {
	source, dest := parseHeader("seed-to-soil map:")
	assert.Equal(t, "seed", source)
	assert.Equal(t, "soil", dest)
}

func TestSearchMap(t *testing.T) {
	var testMap []*RangeEntry
	entry := NewRangeEntry(50, 98, 2)
	testMap = append(testMap, entry)

	val, err := searchMap(50, testMap)
	assert.Equal(t, 98, val)
	assert.Equal(t, nil, err)
}
