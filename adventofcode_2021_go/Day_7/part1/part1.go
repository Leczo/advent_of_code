package part1

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
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
	positions := []int{}

	for scanner.Scan() {
		for _, position := range strings.Split(scanner.Text(), ",") {
			position, err := strconv.Atoi(position)
			errCheck(err)
			positions = append(positions, position)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return positions
}

func median(data []int) []int {
	if len(data)%2 != 0 {
		return []int{data[(len(data)+1)/2]}
	}
	// # WARINNG
	return []int{data[len(data)/2], data[len(data)/2+1]}

}

func countFuelCost(optimalPosition int, data []int) float64 {
	var cost float64
	for _, position := range data {
		cost += math.Abs(float64(position) - float64(optimalPosition))
	}
	return cost

}

func Run() {
	data := loadInput("./input")
	sort.Ints(data)
	middle := median(data)
	var cheapestSolution float64 = math.Inf(1)
	for _, v := range middle {
		solution := countFuelCost(v, data)
		if solution < cheapestSolution {
			cheapestSolution = solution
		}
	}
	log.Println(cheapestSolution)
}
