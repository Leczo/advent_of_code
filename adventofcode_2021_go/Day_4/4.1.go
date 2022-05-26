package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type element struct {
	location  byte // assuming that board size is 5x5
	occurence bool
	row       int
	column    int
}

type Board struct {
	content          []string
	mappedContent    map[string]element
	selectedNumbers  [25]bool // 5x5 board
	rowOccurences    [5]byte
	columnOccurences [5]byte
	bingoNumber      string
}

func NewBoard(arr []string) *Board {
	b := new(Board)
	b.content = arr
	b.mappedContent = boardArrayToMap(arr)
	return b
}

func main() {
	drawnNumbers := loadNumbers("./input2")
	boards := loadBingoBoards("./input1")

out:
	for _, number := range drawnNumbers {
		for _, boardPointer := range boards {
			isWinner := checkBingoBoard(boardPointer, number)
			if isWinner {
				boardPointer.bingoNumber = number
				calculateResult(*boardPointer)
				break out
			}
		}
	}
}

func loadNumbers(input string) []string {

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	f_content := []string{}

	for scanner.Scan() {
		f_content = strings.Split(scanner.Text(), ",")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return f_content
}

func loadBingoBoards(input string) []*Board {

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	boards := []*Board{}
	boardBuffer := []string{}
	i := 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		if len(line) == 0 {
			i++
			boards = append(boards, NewBoard(boardBuffer))
			boardBuffer = []string{}
			continue
		}

		// currentBoard := boards[i].content
		boardBuffer = append(boardBuffer, line...)
		// boards[i].content = currentBoard
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return boards
}

func boardArrayToMap(array []string) map[string]element {
	elementsMap := map[string]element{}

	for index, item := range array {
		row := index / 5
		col := index % 5
		elementsMap[item] = element{location: uint8(index), occurence: true, row: row, column: col}
	}
	return elementsMap
}

func checkBingoBoard(board *Board, drawnElement string) bool { // omitting casting drawnNumber to int because it is unnecessary

	boardElem := board.mappedContent[drawnElement]

	if boardElem.occurence {
		board.selectedNumbers[boardElem.location] = true
		board.rowOccurences[boardElem.row] += 1
		board.columnOccurences[boardElem.column] += 1

		if board.rowOccurences[boardElem.row] == 5 {
			return true
		}
		if board.columnOccurences[boardElem.column] == 5 {
			return true
		}

	}
	return false
}

func calculateResult(b Board) {
	lastNumber, err := strconv.Atoi(b.bingoNumber)
	if err != nil {
		log.Fatal(err)
	}

	var counter int
	for index, value := range b.content {
		if !b.selectedNumbers[index] {
			i, err := strconv.Atoi(value)

			if err != nil {
				log.Fatal(err)
			}
			counter += i
		}
	}
	fmt.Println(counter * lastNumber)
}
