package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func loadInput(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	board := [][]string{}
	line := []string{}

	for scanner.Scan() {
		line = strings.Split(scanner.Text(), "")
		board = append(board, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return board
}

type Stack []string

func (s *Stack) Push(elem string) {
	*s = append(*s, elem)
}

func (s *Stack) Pop() (string, bool) {
	stackSize := len(*s)
	if stackSize == 0 {
		return "", false
	}

	top := (*s)[stackSize-1]
	*s = (*s)[:stackSize-1]
	return top, true
}

func isOpeningChar(char string) bool {
	if char != "{" && char != "[" && char != "(" && char != "<" {
		return false
	}
	return true
}

func isOpeningMatchingEnding(opening, ending string) bool {
	switch {
	case opening == "{" && ending != "}":
		return false
	case opening == "[" && ending != "]":
		return false
	case opening == "<" && ending != ">":
		return false
	case opening == "(" && ending != ")":
		return false
	}

	return true
}

func findIllegalChars(line []string, illegalChars *[]string) {
	var stack Stack = Stack{}
	var openingBracket string

	for _, char := range line {
		if isOpeningChar(char) {
			stack.Push(char)
		} else {
			openingBracket, _ = stack.Pop()
			if !isOpeningMatchingEnding(openingBracket, char) {
				*illegalChars = append(*illegalChars, char)
			}
		}
	}

}

func countErrorScore(illegalChars []string) {
	total := 0
	for _, char := range illegalChars {
		switch char {
		case ")":
			total += 3
		case "}":
			total += 1197
		case "]":
			total += 57
		case ">":
			total += 25137
		}
	}
	fmt.Println("Total errors score:", total)
}

func Run() {
	data := loadInput("./input")
	var illegalChars []string

	for _, line := range data {
		findIllegalChars(line, &illegalChars)
	}
	fmt.Println(illegalChars)
	countErrorScore(illegalChars)

}
