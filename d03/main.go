package main

import (
	_ "embed"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

///source of challenge is : https://adventofcode.com/2023/day/3

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

type numberLocation struct {
	lineNumber int
	startPos   int
	endPos     int
	value      int
}

func findAllNumbers(s []string) (l []numberLocation) {
	n := -1

	for i, line := range s {
		firstFound := false
		for j, char := range line {
			_, err := strconv.Atoi(string(char))

			if err != nil {
				firstFound = false

				continue
			}

			if !firstFound {
				l = append(l, numberLocation{})
				n++
				l[n].lineNumber = i
				l[n].startPos = j
				l[n].endPos = j
				firstFound = true
			} else {
				l[n].endPos = j
			}
		}
	}

	for i, number := range l {
		var value int
		var err error

		if number.endPos+1 > len(s[number.lineNumber]) {
			value, err = strconv.Atoi(s[number.lineNumber][number.startPos:])
		} else {
			value, err = strconv.Atoi(s[number.lineNumber][number.startPos : number.endPos+1])
		}

		if err != nil {
			panic("can't parse int ...")
		}

		l[i].value = value
	}

	return
}

func isSymbol(s rune) bool {
	_, err := strconv.Atoi(string(s))
	if err == nil {
		return false
	}

	if s == '.' {
		return false
	}

	return true
}

func isEnginePart(s []string, l numberLocation) bool {

	var lineNumberUpper, lineNumberLower int

	if l.lineNumber > 0 {
		lineNumberUpper = l.lineNumber - 1
	} else {
		lineNumberUpper = l.lineNumber
	}

	if l.lineNumber < len(s) {
		lineNumberLower = l.lineNumber + 2
	} else {
		lineNumberLower = l.lineNumber + 1
	}

	if lineNumberLower > len(s) {
		for _, line := range s[lineNumberUpper:] {
			var startPosFirst, endPostLast int

			if l.startPos > 0 {
				startPosFirst = l.startPos - 1
			} else {
				startPosFirst = l.startPos
			}

			if l.endPos < len(line) {
				endPostLast = l.endPos + 2
			} else {
				endPostLast = l.endPos + 1
			}

			if endPostLast > len(line) {
				for _, char := range line[startPosFirst:] {
					if isSymbol(char) {
						return true
					}
				}
			} else {
				for _, char := range line[startPosFirst:endPostLast] {
					if isSymbol(char) {
						return true
					}
				}
			}
		}
	} else {
		for _, line := range s[lineNumberUpper:lineNumberLower] {
			var startPosFirst, endPostLast int

			if l.startPos > 0 {
				startPosFirst = l.startPos - 1
			} else {
				startPosFirst = l.startPos
			}

			if l.endPos < len(line) {
				endPostLast = l.endPos + 2
			} else {
				endPostLast = l.endPos + 1
			}

			if endPostLast > len(line) {
				for _, char := range line[startPosFirst:] {
					if isSymbol(char) {
						return true
					}
				}
			} else {
				for _, char := range line[startPosFirst:endPostLast] {
					if isSymbol(char) {
						return true
					}
				}
			}
		}
	}

	return false
}

func solve_part1(input string) int {
	parsed := parseInput(input)
	possibleNumbers := findAllNumbers(parsed)
	sum := 0

	for _, p := range possibleNumbers {

		if isEnginePart(parsed, p) {
			sum += p.value
		}
	}

	return sum
}

type gearLocation struct {
	lineNum int
	pos     int
}

func findAllStars(s []string) (g []gearLocation) {
	for i, line := range s {
		for j, char := range line {
			if char == '*' {
				g = append(g, gearLocation{lineNum: i, pos: j})
			}
		}
	}

	return
}

type numberForGears struct {
	firstDigitLine int
	firstDigitPos  int
}

func figureOutNumber(s []string, line int, pos int) (l numberForGears, value int) {
	l.firstDigitLine = line
	l.firstDigitPos = pos
	for l.firstDigitPos > 0 {
		_, err := strconv.Atoi(string(s[line][l.firstDigitPos-1]))
		if err != nil {
			break
		}
		l.firstDigitPos--
	}

	var numberString string
	for x := l.firstDigitPos; x < len(s[line]); x++ {
		_, err := strconv.Atoi(string(s[line][x]))
		if err != nil {
			x--
			break
		}
		numberString += string(s[line][x])
	}

	value, _ = strconv.Atoi(numberString)

	return
}
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gearRatio(s []string, g gearLocation) (r int, e error) {
	adjacentNumbers := make(map[numberForGears]int)

	lineStart := maxInt(g.lineNum-1, 0)
	posStart := maxInt(g.pos-1, 0)
	lineEnd := minInt(g.lineNum+1, len(s))
	posEnd := minInt(g.pos+1, len(s[lineEnd]))

	for x := lineStart; x <= lineEnd; x++ {
		for y := posStart; y <= posEnd; y++ {
			_, err := strconv.Atoi(string(s[x][y]))
			if err != nil {
				continue
			}
			l, value := figureOutNumber(s, x, y)
			adjacentNumbers[l] = value
		}
	}

	if len(adjacentNumbers) != 2 {
		e = errors.New("This isn't a gear")
		return
	}

	r = 1
	for _, v := range adjacentNumbers {
		r *= v
	}

	return
}

func solve_part2(input string) int {
	parsed := parseInput(input)
	possibleGears := findAllStars(parsed)
	sum := 0

	for _, p := range possibleGears {
		r, err := gearRatio(parsed, p)
		if err != nil {
			continue
		}

		sum += r
	}

	return sum
}

func parseInput(input string) (parsedInput []string) {
	for _, line := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, line)
	}
	return parsedInput
}

func main() {
	var p1, p2 int
	p1 = solve_part1(input)
	p2 = solve_part2(input)
	fmt.Println("part 1 = ", p1, "\npart 2 = ", p2)
}
