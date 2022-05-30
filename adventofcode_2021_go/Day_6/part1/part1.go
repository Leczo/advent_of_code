package part1

import (
	"bufio"
	"fmt"
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

func loadInput(path string) []Fish {
	f, err := os.Open(path)
	errCheck(err)

	defer f.Close()
	scanner := bufio.NewScanner(f)
	fishes := []Fish{}

	for scanner.Scan() {
		for _, fish := range strings.Split(scanner.Text(), ",") {
			state, err := strconv.Atoi(fish)
			errCheck(err)
			fishes = append(fishes, Fish{state: state})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return fishes
}

type Fish struct {
	state int
}

func (f *Fish) resetState() {
	f.state = 7
}

func (f *Fish) decreaseState() Fish {
	f.state = f.state - 1
	return *f
}

func nextDay(f []Fish) []Fish {
	oldGeneration := []Fish{}
	newGeneration := []Fish{}

	for _, fish := range f {
		if fish.state == 0 {
			newGeneration = append(newGeneration, Fish{state: 8})
			fish.resetState()
		}
		oldGeneration = append(oldGeneration, fish.decreaseState())

	}
	oldGeneration = append(oldGeneration, newGeneration...)
	return oldGeneration
}

func Run() {
	fishes := loadInput("./input")
	for i := 0; i < 80; i++ {
		fishes = nextDay(fishes)
	}
	fmt.Println(len(fishes))
}
