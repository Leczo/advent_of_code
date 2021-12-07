package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func load_input(input string) []int {
	f, err := os.Open(input)
	err_check(err)

	defer f.Close()
	scanner := bufio.NewScanner(f)
	resultArray := make([]int, 0)

	for scanner.Scan() {

		int_value, err := strconv.Atoi(scanner.Text())
		err_check(err)

		resultArray = append(resultArray, int_value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return resultArray

}

func err_check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	path := "./input"
	data := load_input(path)

	inc_counter := 0
	for index := range data {
		if index == 0 || index == 1 || index == 2 {
			continue
		}

		f_block := data[index-3] + data[index-2] + data[index-1]
		s_block := data[index-2] + data[index-1] + data[index]

		if f_block < s_block {
			inc_counter += 1
		}

	}
	fmt.Println(inc_counter)
}
