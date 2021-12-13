package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func load_input(input string) []string {
	f, err := os.Open(input)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	f_content := []string{}

	for scanner.Scan() {
		f_content = append(f_content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return f_content

}

func bit_occurence(data []string) []int {
	// Adding ones on every position in every element
	counter := make([]int, len(data[0]))

	for _, values := range data {

		bits_slice := strings.Split(values, "")
		for index, bit := range bits_slice {
			if bit == "1" {
				counter[index] += 1
			}
		}
	}
	return counter
}

func count_ratio(arr []int, num_of_el int) (int64, int64) {
	gamma_str, epsilon_str := "", ""

	// If number is greater than number of elements then number is positive
	for _, value := range arr {
		if value > num_of_el-value {

			gamma_str += "1"
			epsilon_str += "0"
		} else {
			gamma_str += "0"
			epsilon_str += "1"
		}
	}
	fmt.Println(gamma_str, epsilon_str)

	gamma, err := strconv.ParseInt(gamma_str, 2, 0)
	check(err)
	epsilon, err := strconv.ParseInt(epsilon_str, 2, 0)
	check(err)

	return gamma, epsilon
}

func main() {
	path := "./input"
	data := load_input(path)
	num_of_el := len(data)
	num_of_bits := bit_occurence(data)
	fmt.Println(num_of_bits, num_of_el)

	gamma, epsilon := count_ratio(num_of_bits, num_of_el)
	fmt.Println(gamma * epsilon)
}
