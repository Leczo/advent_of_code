package part1

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func loadInput(path string) []string {
	f, err := os.Open(path)
	errCheck(err)

	defer f.Close()
	scanner := bufio.NewScanner(f)
	outputVals := []string{}

	for scanner.Scan() {
		signals := strings.Split(scanner.Text(), "|")
		rightSide := signals[1]
		outputVals = append(outputVals, strings.Split(rightSide, " ")...)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return outputVals
}

func Run() {
	data := loadInput("./input")

	sum := 0
	for _, value := range data {
		switch {
		case len(value) == 2:
			sum += 1
		case len(value) == 4:
			sum += 1
		case len(value) == 3:
			sum += 1
		case len(value) == 7:
			sum += 1
		}
	}
	log.Println(sum)
}
