package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		return "", true
	}

	top := (*s)[stackSize-1]
	*s = (*s)[:stackSize-1]
	return top, false
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

func findIllegalChars(line []string, illegalChars *[]string) (Stack, bool) {
	var stack Stack = Stack{}
	var openingBracket string
	var hasIllegal bool

	for _, char := range line {
		if isOpeningChar(char) {
			stack.Push(char)
		} else {
			openingBracket, _ = stack.Pop()
			if !isOpeningMatchingEnding(openingBracket, char) {
				hasIllegal = true
				*illegalChars = append(*illegalChars, char)
			}
		}
	}
	return stack, hasIllegal
}

func getBracketScore(char string) int {
	switch char {
	case "{": // therefore } to complete
		return 3
	case "(": // therefore 0 to complete
		return 1
	case "[": // therefore ] to complete
		return 2
	case "<": // therefore > to complete
		return 4
	}
	return 0
}

func autocompleteScore(stack Stack) int {
	bracket := ""
	empty := false
	total := 0

	for {
		bracket, empty = stack.Pop()
		if empty {
			break
		}
		total = (total * 5) + getBracketScore(bracket)
	}
	return total
}

func Run() {
	data := loadInput("./input")
	var illegalChars []string
	stacks := []Stack{}

	for _, line := range data {
		stack, invalidChars := findIllegalChars(line, &illegalChars)
		if invalidChars {
			continue
		}
		stacks = append(stacks, stack)

	}

	scores := []int{}
	for _, stack := range stacks {
		fmt.Println(stack)
		fmt.Println()
		fmt.Println()
		score := autocompleteScore(stack)
		scores = append(scores, score)
	}

	sort.Ints(scores)

	fmt.Println(scores)
	fmt.Println("Middle autocompletion score:", scores[len(scores)/2])
}
