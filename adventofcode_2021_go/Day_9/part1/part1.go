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

func loadInput(path string) [][]int {
	f, err := os.Open(path)
	errCheck(err)

	defer f.Close()
	scanner := bufio.NewScanner(f)
	board := [][]int{}
	row := []int{}

	for scanner.Scan() {

		row = sliceToInts(strings.Split(scanner.Text(), ""))
		board = append(board, row)
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

func adjacentCheck(data [][]int) []int {
	var bottom, left, right, top int
	lowPoints := []int{} //map[int]int{}

	for i, row := range data {
		for j, num := range row {
			top = i - 1
			bottom = i + 1
			left = j - 1
			right = j + 1

			if top > -1 && data[top][j] <= num {
				continue
			}
			if bottom < len(data) && data[i+1][j] <= num {
				continue
			}
			if left > -1 && data[i][left] <= num {
				continue
			}
			if right < len(row) && data[i][right] <= num {
				continue
			}

			lowPoints = append(lowPoints, num)
		}
	}
	return lowPoints
}

func riskLevel(arr []int) {
	var sum int
	for _, num := range arr {
		sum += num + 1
	}
	println(sum)
}

func Run() {
	data := loadInput("./input")

	lowPoints := adjacentCheck(data)
	log.Println(lowPoints)

	riskLevel(lowPoints)

}
