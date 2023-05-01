package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	rowIdx int
	colIdx int
}

type LowerPoint struct {
	rowIdx      int
	colIdx      int
	basinPoints map[Point]int
	basinSize   int
}

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

		row, _ = sliceToInts(strings.Split(scanner.Text(), ""))
		board = append(board, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return board
}

func sliceToInts(s []string) ([]int, error) {
	result := make([]int, len(s))
	for i, str := range s {
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}

func findLowPoints(data [][]int) []LowerPoint {
	lowPoints := []LowerPoint{}
	numRows := len(data)
	numCols := len(data[0])

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			if isLowPoint(data, i, j) {
				lowPoints = append(lowPoints, LowerPoint{
					rowIdx:      i,
					colIdx:      j,
					basinPoints: make(map[Point]int),
				})
			}
		}
	}

	return lowPoints
}

func isLowPoint(data [][]int, i, j int) bool {
	val := data[i][j]
	if i > 0 && data[i-1][j] < val { // check top
		return false
	}
	if i < len(data)-1 && data[i+1][j] < val { // check bottom
		return false
	}
	if j > 0 && data[i][j-1] < val { // check left
		return false
	}
	if j < len(data[i])-1 && data[i][j+1] < val { // check right
		return false
	}
	return true
}

func searchBasinPoints(data *[][]int, lowerPoint *LowerPoint, pointer *Point) {
	// Stop rescursion rules
	switch {
	case pointer.rowIdx > len(*data)-1:
		return
	case pointer.rowIdx < 0:
		return
	case pointer.colIdx > len((*data)[0])-1:
		return
	case pointer.colIdx < 0:
		return
	case (*data)[pointer.rowIdx][pointer.colIdx] == 9:
		return
	}

	if _, isPointAlreadyInMap := lowerPoint.basinPoints[*pointer]; isPointAlreadyInMap {
		return
	}

	// Add basin point to map
	value := (*data)[pointer.rowIdx][pointer.colIdx]
	lowerPoint.basinPoints[*pointer] = value

	// Create a list of neighboring cells
	neighbors := []Point{
		{colIdx: pointer.colIdx, rowIdx: pointer.rowIdx + 1}, // top
		{colIdx: pointer.colIdx, rowIdx: pointer.rowIdx - 1}, // bottom
		{colIdx: pointer.colIdx - 1, rowIdx: pointer.rowIdx}, // left
		{colIdx: pointer.colIdx + 1, rowIdx: pointer.rowIdx}, // right
	}

	for _, neighbor := range neighbors {
		searchBasinPoints(data, lowerPoint, &neighbor)
	}
}

func findBasinNeighbours(data [][]int, lPoints []LowerPoint) []LowerPoint {

	pointer := Point{}
	lPointsWithBasins := []LowerPoint{}

	for _, lpoint := range lPoints {

		pointer.colIdx = lpoint.colIdx
		pointer.rowIdx = lpoint.rowIdx

		searchBasinPoints(&data, &lpoint, &pointer)

		lpoint.basinSize = len(lpoint.basinPoints)
		lPointsWithBasins = append(lPointsWithBasins, lpoint)

	}
	return lPointsWithBasins
}

func largestBasinsTotalSize(lowerPoints []LowerPoint) {
	sort.Slice(lowerPoints, func(i, j int) bool {
		return lowerPoints[i].basinSize > lowerPoints[j].basinSize
	})

	total := lowerPoints[0].basinSize * lowerPoints[1].basinSize * lowerPoints[2].basinSize
	fmt.Println("Total: ", total)
}

func Run() {
	data := loadInput("./input")

	lowPoints := findLowPoints(data)
	lowPointsWithBasins := findBasinNeighbours(data, lowPoints)
	largestBasinsTotalSize(lowPointsWithBasins)

}
