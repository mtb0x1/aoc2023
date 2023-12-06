package main

import (
	_ "embed"
	"fmt"
	"strings"
)

///source of challenge is : https://adventofcode.com/2023/day/4

//go:embed input.txt
var input string

func solve_part1(data string) int {
	lines := strings.Split(string(data), "\n")
	lineCount := len(lines)

	cards := make([]ScratchCard, lineCount)

	var total int = 0

	for index, element := range lines {
		card := ProcessLine(element, index)
		cards = append(cards, card)
		total += card.Points
	}

	return total
}

func CheckNumber(search []string, nums []string) int {

	total := 0

	for i := 0; i < len(search); i++ {
		for j := 0; j < len(nums); j++ {
			if len(nums[j]) > 0 {
				if nums[j] == search[i] {
					if total == 0 {
						total = 1
					} else {
						total = total * 2
					}
					break
				}
			}
		}
	}

	return total
}

type ScratchCard struct {
	Numbers        []string
	WinningNumbers []string
	Points         int
	Cards          int
}

func isset(arr []ScratchCard, index int) bool {
	return (len(arr) > index)
}

func ProcessLine(line string, index int) ScratchCard {

	line = line[10:]

	sides := strings.Split(line, " | ")
	card := ScratchCard{
		Numbers:        strings.Split(sides[0], " "),
		WinningNumbers: strings.Split(sides[1], " "),
		Points:         0,
		Cards:          1,
	}

	card.Points = CheckNumber(card.Numbers, card.WinningNumbers)

	return card
}

func solve_part2(data string) int {
	lines := strings.Split(string(data), "\n")
	cards := make([]ScratchCard, len(lines))

	for index, element := range lines {
		element = element[10:]

		sides := strings.Split(element, " | ")
		cards[index].Numbers = strings.Split(sides[0], " ")
		cards[index].WinningNumbers = strings.Split(sides[1], " ")
		cards[index].Cards++

		for i := 0; i < len(cards[index].Numbers); i++ {
			for j := 0; j < len(cards[index].WinningNumbers); j++ {
				if len(cards[index].Numbers[i]) > 0 {
					if cards[index].Numbers[i] == cards[index].WinningNumbers[j] {
						cards[index].Points++
						break
					}
				}
			}
		}

		for i := 0; i < cards[index].Points; i++ {
			in := index + i + 1
			if isset(cards, in) {
				cards[in].Cards += cards[index].Cards
			}
		}
	}

	total := 0
	for _, element := range cards {
		total += element.Cards
	}

	return total
}

func main() {
	var p1, p2 int
	p1 = solve_part1(input)
	p2 = solve_part2(input)
	fmt.Println("part 1 = ", p1, "\npart 2 = ", p2)
}
