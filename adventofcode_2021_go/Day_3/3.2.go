package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func get_rating_value(priority string, data_map []string) string {
	subset := data_map
	position := 0
	for {
		counter := make(map[string][]string)
		for _, bit_arr := range subset {
			if string(bit_arr[position]) == "1" {
				counter["ones"] = append(counter["ones"], bit_arr)
			} else if string(bit_arr[position]) == "0" {
				counter["zeros"] = append(counter["zeros"], bit_arr)
			}
		}
		// Rule check
		switch {
		case len(counter["ones"]) == len(counter["zeros"]):
			subset = counter[priority]
		case len(counter["ones"]) > len(counter["zeros"]) && priority == "ones":
			subset = counter["ones"]
		case len(counter["ones"]) < len(counter["zeros"]) && priority == "ones":
			subset = counter["zeros"]
		case len(counter["ones"]) > len(counter["zeros"]) && priority == "zeros":
			subset = counter["zeros"]
		case len(counter["ones"]) < len(counter["zeros"]) && priority == "zeros":
			subset = counter["ones"]
		}

		// Returning last array
		if len(subset) == 1 {
			return subset[0]
		}

		position += 1 // moving through positions of bit array
		if position == len(subset[0]) {
			break
		}
	}
	return "Error"
}

func to_decimal(data string) int64 {
	dec, err := strconv.ParseInt(data, 2, 0)
	check(err)
	return dec
}
func main() {
	path := "./input"
	data := load_input(path)
	oxygen_ratio := get_rating_value("ones", data)
	co2_ratio := get_rating_value("zeros", data)
	fmt.Println(oxygen_ratio, co2_ratio)
	fmt.Println(to_decimal(oxygen_ratio) * to_decimal(co2_ratio))

}
