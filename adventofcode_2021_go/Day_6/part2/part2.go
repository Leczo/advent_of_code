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
	fishes := []int{}

	for scanner.Scan() {
		for _, population := range strings.Split(scanner.Text(), ",") {
			fish, err := strconv.Atoi(population)
			errCheck(err)
			fishes = append(fishes, fish)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return fishes
}

func nextDay(p map[int]int) map[int]int {
	nextGen := map[int]int{}
	for age, size := range p {
		if age-1 == -1 {
			nextGen[6] += size
			nextGen[8] += size
		} else {
			nextGen[age-1] += size
		}

	}
	return nextGen
}

func sum(p map[int]int) int {
	var x int
	for _, size := range p {
		x += size
	}
	return x
}

func Run() {
	fishes := loadInput("./input")
	sort.Ints(fishes)
	population := map[int]int{}
	for _, fish := range fishes {
		population[fish] += 1
	}

	var entire int
	for i := 1; i <= 256; i++ {
		population = nextDay(population)
		entire = sum(population)

		fmt.Printf("Day: %v, Population size : %v \n", i, entire)
	}

}
