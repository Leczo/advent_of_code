package part2

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

func average(data []int) int {
	var sum int
	for _, value := range data {
		sum += value
	}
	return int(math.Round(float64(sum) / float64(len(data))))
}

func countFuelCost(optimalPosition int, data []int) int {
	var cost int
	for _, position := range data {
		changes := math.Abs(float64(position) - float64(optimalPosition))
		cost += ((1 + int(changes)) * int(changes)) / 2
		// log.Println(position, optimalPosition, changes, cost)
	}
	return cost

}

func Run() {
	data := loadInput("./input")
	sort.Ints(data)
	avg := average(data)
	avgs := []int{avg, avg - 1, avg + 1}
	var cheapestSolution float64 = math.Inf(1)
	for _, v := range avgs {
		solution := countFuelCost(v, data)
		if float64(solution) < cheapestSolution {
			cheapestSolution = float64(solution)
		}
	}
	log.Println(int(cheapestSolution))
}
