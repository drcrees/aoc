package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var hands []Hand

type Hand struct {
	cards      string
	bid        int
	handType   int
	cardValues []int
}

var faceCards = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 1, // weakling
	"T": 10,
}

func main() {
	fmt.Println("--- 7-2 ---")

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parse(scanner.Text())
	}

	for i := range hands {
		hands[i].handType = getHandType(&hands[i])
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		}

		for k, left := range hands[i].cardValues {
			right := hands[j].cardValues[k]
			if left == right {
				continue
			}
			return left < right
		}

		return hands[i].handType < hands[j].handType
	})

	var result = 0
	for i := len(hands) - 1; i >= 0; i-- {
		result += (hands[i].bid * (i + 1))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result: %v", result)
}

func getHandType(hand *Hand) int {
	var cardCounts = make(map[string]int)
	var numJokers = 0

	for _, card := range strings.Split(hand.cards, "") {
		cardCounts[card]++

		if card == "J" {
			numJokers++
		}

		if val, ok := faceCards[card]; ok {
			hand.cardValues = append(hand.cardValues, val)
		} else {
			val, _ := strconv.Atoi(card)
			hand.cardValues = append(hand.cardValues, val)
		}
	}

	var highest = ""
	for card, val := range cardCounts {
		if val > cardCounts[highest] && card != "J" {
			highest = card
		}
	}

	cardCounts[highest] += numJokers

	var threeOfKind, pair = false, false
	for _, count := range cardCounts {
		if count == 5 {
			return 7
		}
		if count == 4 {
			return 6
		}
		if count == 3 {
			threeOfKind = true
			if numJokers == 2 {
				return 4
			}
		}
		if count == 2 {
			if pair && numJokers == 2 {
				return 3
			}
			if pair {
				return 3
			}
			pair = true
		}
	}

	if threeOfKind && pair {
		return 5
	}

	if threeOfKind {
		return 4
	}

	if pair {
		return 2
	}

	return 1
}

func parse(str string) {
	hand := strings.Split(str, " ")
	bid, _ := strconv.Atoi(hand[1])
	hands = append(hands, Hand{cards: hand[0], bid: bid})
}
