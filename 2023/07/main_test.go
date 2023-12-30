package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseHand(t *testing.T) {
	hand := parseHand("32T3K")
	assert.Equal(t, []rune{'3', '2', 'T', '3', 'K'}, hand.cards)
}

func TestFiveOfAKind(t *testing.T) {
	{
		hand := parseHand("32T3K")
		assert.Equal(t, false, hand.isFiveOfAKind())
	}
	{
		hand := parseHand("KKKKK")
		assert.Equal(t, true, hand.isFiveOfAKind())
		assert.Equal(t, FiveOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("KKKJJ")
		assert.Equal(t, true, hand.isFiveOfAKind())
		assert.Equal(t, FiveOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("KKJJJ")
		assert.Equal(t, true, hand.isFiveOfAKind())
		assert.Equal(t, FiveOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("KJJJJ")
		assert.Equal(t, true, hand.isFiveOfAKind())
		assert.Equal(t, FiveOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("JJJJJ")
		assert.Equal(t, true, hand.isFiveOfAKind())
		assert.Equal(t, FiveOfAKind, hand.handType)
	}
}

func TestFourOfAKind(t *testing.T) {
	{
		hand := parseHand("32T3K")
		assert.Equal(t, false, hand.isFourOfAKind())
		assert.Equal(t, OnePair, hand.handType)
	}
	{
		hand := parseHand("KKKKJ")
		assert.Equal(t, true, hand.isFourOfAKind())
		assert.Equal(t, FourOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("KKKKJ")
		assert.Equal(t, false, hand.isFourOfAKind())
		assert.Equal(t, FiveOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("JJJJK")
		assert.Equal(t, false, hand.isFourOfAKind())
		assert.Equal(t, FiveOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("JJKK2")
		assert.Equal(t, true, hand.isFourOfAKind())
		assert.Equal(t, FourOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("KTJJT")
		assert.Equal(t, true, hand.isFourOfAKind())
		assert.Equal(t, FourOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("QQQJA")
		assert.Equal(t, true, hand.isFourOfAKind())
		assert.Equal(t, FourOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("T55J5")
		assert.Equal(t, true, hand.isFourOfAKind())
		assert.Equal(t, FourOfAKind, hand.handType)
	}
}

func TestFullHouse(t *testing.T) {
	{
		hand := parseHand("33322")
		assert.Equal(t, true, hand.isFullHouse())
		assert.Equal(t, FullHouse, hand.handType)
	}
	{
		hand := parseHand("KKKKJ")
		assert.Equal(t, false, hand.isFullHouse())
	}
	{
		hand := parseHandJokers("KKK2J")
		assert.Equal(t, false, hand.isFullHouse())
		assert.Equal(t, FourOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("KK22J")
		assert.Equal(t, true, hand.isFullHouse())
		assert.Equal(t, FullHouse, hand.handType)
	}
	{
		hand := parseHandJokers("QQQJA")
		assert.Equal(t, false, hand.isFullHouse())
	}
}

func TestThreeOfAKind(t *testing.T) {
	{
		hand := parseHand("32T3K")
		assert.Equal(t, false, hand.isThreeOfAKind())
	}
	{
		hand := parseHand("KKKKJ")
		assert.Equal(t, false, hand.isThreeOfAKind())
	}
	{
		hand := parseHand("KKKJJ")
		assert.Equal(t, false, hand.isThreeOfAKind())
	}
	{
		hand := parseHand("KKKJT")
		assert.Equal(t, true, hand.isThreeOfAKind())
		assert.Equal(t, ThreeOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("234JJ")
		assert.Equal(t, true, hand.isThreeOfAKind())
		assert.Equal(t, ThreeOfAKind, hand.handType)
	}
	{
		hand := parseHandJokers("223JJ")
		assert.Equal(t, false, hand.isThreeOfAKind())
	}
	{
		hand := parseHandJokers("QQQJA")
		assert.Equal(t, false, hand.isThreeOfAKind())
		assert.Equal(t, FourOfAKind, hand.handType)
	}
}

func TestTwoPair(t *testing.T) {
	{
		hand := parseHandJokers("2345J")
		assert.Equal(t, false, hand.isTwoPair())
	}
	{
		hand := parseHandJokers("2245J")
		assert.Equal(t, false, hand.isTwoPair())
	}
	{
		hand := parseHandJokers("22455")
		assert.Equal(t, true, hand.isTwoPair())
		assert.Equal(t, TwoPair, hand.handType)
	}
	{
		hand := parseHandJokers("KK677")
		assert.Equal(t, true, hand.isTwoPair())
		assert.Equal(t, TwoPair, hand.handType)
	}
}

func TestOnePair(t *testing.T) {
	{
		hand := parseHandJokers("23456")
		assert.Equal(t, false, hand.isOnePair())
	}
	{
		hand := parseHandJokers("2345J")
		assert.Equal(t, true, hand.isOnePair())
		assert.Equal(t, OnePair, hand.handType)
	}
}

func TestHighCard(t *testing.T) {
	{
		hand := parseHandJokers("23456")
		assert.Equal(t, true, hand.isHighCard())
		assert.Equal(t, HighCard, hand.handType)
	}
	{
		hand := parseHandJokers("2345J")
		assert.Equal(t, false, hand.isHighCard())
	}
}
