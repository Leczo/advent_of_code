package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	direction string
	value     int
}

func err_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func load_input(input string) []Pair {
	f, err := os.Open(input)
	err_check(err)

	defer f.Close()
	scanner := bufio.NewScanner(f)
	resultArray := make([]Pair, 0)

	for scanner.Scan() {

		split := strings.Split(scanner.Text(), " ")

		d := split[0]
		v, err := strconv.Atoi(split[1])
		err_check(err)

		pair := Pair{
			direction: d,
			value:     v,
		}

		resultArray = append(resultArray, pair)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return resultArray

}

func count_position(coordinates []Pair) (int, int) {
	horizontal_p, depth_p, aim := 0, 0, 0

	for _, coord := range coordinates {

		switch coord.direction {
		case "forward":
			horizontal_p += coord.value
			depth_p += aim * coord.value
		case "down":
			aim += coord.value
		case "up":
			aim -= coord.value

		}
	}
	return horizontal_p, depth_p
}

func main() {
	path := "./input"
	data := load_input(path)
	h_pos, dep_pos := count_position(data)
	product := h_pos * dep_pos
	fmt.Println(product)
}
