package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type HandType int64

const (
	Unknown      HandType = -1
	HighCard     HandType = 0
	OnePair      HandType = 1
	TwoPair      HandType = 2
	ThreeOfAKind HandType = 3
	FullHouse    HandType = 4
	FourOfAKind  HandType = 5
	FiveOfAKind  HandType = 6
)

type Hand struct {
	cards     []rune
	handType  HandType
	cardMap   map[rune]int
	bid       int
	jokers    bool
	numJokers int
}

func NewHand(cards []rune) *Hand {
	hand := Hand{cards, Unknown, make(map[rune]int), 0, false, 0}
	return &hand
}

func parseHand(input string) *Hand {
	return parseHandImpl(input, false)
}

func parseHandJokers(input string) *Hand {
	return parseHandImpl(input, true)
}

func parseHandImpl(input string, jokers bool) *Hand {
	cards := []rune(input)

	hand := Hand{cards, Unknown, make(map[rune]int), 0, jokers, 0}

	for _, val := range hand.cards {
		if jokers {
			if val == 'J' {
				hand.numJokers += 1
			}
		}
		count, ok := hand.cardMap[val]
		if ok {
			hand.cardMap[val] = count + 1
		} else {
			hand.cardMap[val] = 1
		}
	}

	if hand.isHighCard() {
		hand.handType = HighCard
	} else if hand.isOnePair() {
		hand.handType = OnePair
	} else if hand.isTwoPair() {
		hand.handType = TwoPair
	} else if hand.isThreeOfAKind() {
		hand.handType = ThreeOfAKind
	} else if hand.isFullHouse() {
		hand.handType = FullHouse
	} else if hand.isFourOfAKind() {
		hand.handType = FourOfAKind
	} else if hand.isFiveOfAKind() {
		hand.handType = FiveOfAKind
	}

	return &hand
}
func (hand *Hand) setBid(bidStr string) {
	bid, _ := strconv.Atoi(bidStr)
	hand.bid = bid
}

func (hand *Hand) toString() string {
	s := string(hand.cards)
	return s
}

func compare(lhs rune, rhs rune, jokers bool) bool {

	cardValue := map[rune]int{}
	cardValue['2'] = 2
	cardValue['3'] = 3
	cardValue['4'] = 4
	cardValue['5'] = 5
	cardValue['6'] = 6
	cardValue['7'] = 7
	cardValue['8'] = 8
	cardValue['9'] = 9
	cardValue['T'] = 10
	if jokers {
		cardValue['J'] = 1
	} else {
		cardValue['J'] = 10
	}
	cardValue['Q'] = 12
	cardValue['K'] = 13
	cardValue['A'] = 14

	return cardValue[lhs] < cardValue[rhs]

	return false
}

func (lhs *Hand) lessThan(rhs *Hand) bool {
	if lhs.handType < rhs.handType {
		return true
	} else if lhs.handType == rhs.handType {
		for i := 0; i < 5; i++ {
			if lhs.cards[i] != rhs.cards[i] {
				return compare(lhs.cards[i], rhs.cards[i], lhs.jokers)
			}
		}
	}

	return false
}

func (hand *Hand) isFiveOfAKind() bool {
	five := false
	four := false
	three := false
	pairs := 0
	for _, count := range hand.cardMap {
		if count == 5 {
			five = true
		} else if count == 4 {
			four = true
		} else if count == 3 {
			three = true
		} else if count == 2 {
			pairs += 1
		}
	}

	// cases that result in five of a kind:
	//   - five of a kind (including jokers)
	//   - four of a kind and a joker
	//   - three of a kind and two jokers
	//   - two of a kind and three jokers (form of full house)
	//   - high card and four jokers
	if hand.jokers {
		if five {
			return true
		}
		if four && hand.numJokers == 1 {
			return true
		} else if three && hand.numJokers == 2 {
			return true
		} else if pairs == 1 && hand.numJokers == 3 {
			return true
		} else if hand.numJokers == 4 {
			return true
		} else {
			return false
		}
	}

	// no jokers
	return five
}

func (hand *Hand) isFourOfAKind() bool {
	four := false
	three := false
	pairs := 0
	for _, count := range hand.cardMap {
		if count == 4 {
			four = true
		} else if count == 3 {
			three = true
		} else if count == 2 {
			pairs += 1
		}
	}

	// cases that result in four of a kind:
	//   - four of a kind (no jokers)
	//   - three of a kind and a joker
	//   - two of a kind and two jokers (form of two pair)
	//   - high card and three jokers
	if hand.jokers {
		if four && hand.numJokers == 0 {
			return true
		} else if three && hand.numJokers == 1 {
			return true
		} else if pairs == 2 && hand.numJokers == 2 {
			return true
		} else if pairs == 0 && hand.numJokers == 3 {
			return true
		} else {
			return false
		}
	}

	// no jokers
	return four
}

func (hand *Hand) isThreeOfAKind() bool {
	three := false
	pairs := 0
	maxCount := 0

	for _, count := range hand.cardMap {
		if count > maxCount {
			maxCount = count
		}
		if count == 3 {
			three = true
		}
		if count == 2 {
			pairs += 1
		}
	}

	// cases that result in three of a kind:
	//  - three of a kind (no other pair and no jokers)
	//  - one pair and a joker
	//  - high card and two jokers
	if hand.jokers {
		if three && pairs == 0 && hand.numJokers == 0 {
			return true
		} else if pairs == 1 && hand.numJokers == 1 {
			return true
		} else if !three && pairs == 1 && hand.numJokers == 2 {
			return true
		} else {
			return false
		}
	}

	// no jokers
	return (three && pairs == 0)
}

func (hand *Hand) isFullHouse() bool {
	three := false
	pairs := 0

	for _, count := range hand.cardMap {
		if count == 3 {
			three = true
		}
		if count == 2 {
			pairs += 1
		}
	}

	// cases that result in a full house:
	//  - three of a kind & a pair (no jokers)
	//  - two pair and a joker
	if hand.jokers {
		if pairs == 1 && three && hand.numJokers == 0 {
			return true
		} else if pairs == 2 && hand.numJokers == 1 {
			return true
		} else {
			return false
		}
	}

	// no jokers
	return (three && pairs == 1)
}

func (hand *Hand) isTwoPair() bool {
	pairs := 0

	for _, count := range hand.cardMap {
		if count == 2 {
			pairs += 1
		}
	}

	// cases that result in two pair:
	//  - two pair and no jokers
	if hand.jokers {
		if pairs == 2 && hand.numJokers == 0 {
			return true
		} else {
			return false
		}
	}

	// no jokers
	return pairs == 2
}

func (hand *Hand) isOnePair() bool {
	pairs := 0
	three := false
	four := false

	for _, count := range hand.cardMap {
		if count == 2 {
			pairs += 1
		} else if count == 3 {
			three = true
		} else if count == 4 {
			four = true
		}
	}

	// cases that result in one pair:
	//   - one pair and no jokers
	//   - high card and one joker
	if hand.jokers {
		if pairs == 1 && !three && hand.numJokers == 0 {
			return true
		} else if !three && !four && pairs == 0 && hand.numJokers == 1 {
			return true
		} else {
			return false
		}

	}

	// no jokers
	return (pairs == 1 && !three)
}

func (hand *Hand) isHighCard() bool {
	maxCount := 0

	for _, count := range hand.cardMap {
		if count > maxCount {
			maxCount = count
		}
	}

	// jokers increase the number of matches
	if hand.jokers {
		maxCount += hand.numJokers
	}

	return maxCount == 1
}

func main() {
	input := "input.txt"

	{
		readFile, _ := os.Open(input)
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		hands := []*Hand{}

		for fileScanner.Scan() {
			handStr := fileScanner.Text()[0:5]
			bidStr := fileScanner.Text()[6:]
			hand := parseHand(handStr)
			hand.setBid(bidStr)
			hands = append(hands, hand)
		}

		sort.Slice(hands[:], func(i, j int) bool {
			return hands[i].lessThan(hands[j])
		})

		winnings := 0
		for rank, h := range hands {
			fmt.Println((rank + 1), h.toString(), h.bid)
			winnings += ((rank + 1) * h.bid)
		}

		fmt.Println("winnings:", winnings)
	}

	{
		readFile, _ := os.Open(input)
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		hands := []*Hand{}

		for fileScanner.Scan() {
			handStr := fileScanner.Text()[0:5]
			bidStr := fileScanner.Text()[6:]
			hand := parseHandJokers(handStr)
			hand.setBid(bidStr)
			hands = append(hands, hand)
		}

		sort.Slice(hands[:], func(i, j int) bool {
			return hands[i].lessThan(hands[j])
		})

		winnings := 0
		for rank, h := range hands {
			fmt.Println((rank + 1), h.toString(), h.bid)
			winnings += ((rank + 1) * h.bid)
		}

		fmt.Println("winnings:", winnings)
	}
}

// guess: 249863993
// guess: 250432965
// guess: 251652067
// guess: 251464789
// guess: 251528448
// guess: 251660760
