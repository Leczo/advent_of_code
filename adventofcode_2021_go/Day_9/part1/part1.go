package part1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)


func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadInput(path string) []int {
	f, err := os.Open(path)
	errCheck(err)

	defer f.Close()
	scanner := bufio.NewScanner(f)
	board := []int{}

	for scanner.Scan() {

		cells := sliceToInts(strings.Split(scanner.Text(), ""))
		board = append(board, cells...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return board
}

func sliceToInts(s []string) []int {

	intSlice := []int{}
	for _, elem := range s {
		n, err := strconv.Atoi(elem)
		if err != nil {
			log.Println(err)
		} else {
			intSlice = append(intSlice, n)
		}

	}
	return intSlice
}

type Neighbour {
	position string
	value int 

}

type Cell struct {
	// Top, Bottom, Left, Right bool
	Neighbours  []Neighbour
}

func calcNeighbours(board []int) []Cell {
	cells := []Cell{}
	for index, cell := range board {
		//left neighbour
		cells = append(cells, board[index-1])	
	}

}

func Run() {
	data := loadInput("./input")
	log.Println(data)
	bSize := len(data)
	eSize := int(math.Sqrt(bSize))
	log.Println(bSize,eSize)

}
