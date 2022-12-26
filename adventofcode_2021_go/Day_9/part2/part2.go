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

type Signal struct {
	leftSide  []map[string]struct{}
	rightSide []map[string]struct{}
}

func loadInput(path string) []Signal {
	f, err := os.Open(path)
	errCheck(err)

	defer f.Close()
	scanner := bufio.NewScanner(f)
	outputVals := []Signal{}

	for scanner.Scan() {
		signals := strings.Split(scanner.Text(), "|")
		rightSide := signals[1]
		leftSide := signals[0]
		signal := Signal{}
		for _, value := range strings.Split(rightSide, " ") {
			if value != " " && value != "  " && value != "" {
				charsMap := toMap(strings.Split(value, ""))
				signal.rightSide = append(signal.rightSide, charsMap)

			}
		}
		for _, value := range strings.Split(leftSide, " ") {
			if value != " " && value != "  " && value != "" {
				charsMap := toMap(strings.Split(value, ""))
				signal.leftSide = append(signal.leftSide, charsMap)
			}
		}
		outputVals = append(outputVals, signal)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return outputVals
}

func toMap(s []string) map[string]struct{} {
	// Just storing keys so skipping values with struct
	m := map[string]struct{}{}
	for _, char := range s {
		m[char] = struct{}{}
	}
	return m
}

func extractNumberParts(s []map[string]struct{}) []map[string]struct{} {
	sort.Slice(s, func(i, j int) bool {
		return len(s[i]) < len(s[j])
	})
	one := s[0]
	seven := s[1]
	four := s[2]
	eight := s[len(s)-1]

	three := map[string]struct{}{}
	nine := map[string]struct{}{}
	six := map[string]struct{}{}
	zero := map[string]struct{}{}
	five := map[string]struct{}{}
	two := map[string]struct{}{}

	fiveParts := []map[string]struct{}{s[3], s[4], s[5]}
	for _, number := range fiveParts {
		if ok, _ := isContained(seven, number); ok {
			three = number
		} else if _, occurence := isContained(four, number); occurence == 3 {
			five = number
		} else {
			two = number
		}
	}

	sixParts := []map[string]struct{}{s[6], s[7], s[8]}
	for _, number := range sixParts {

		if ok, _ := isContained(three, number); ok {
			nine = number
		} else if ok, _ := isContained(one, number); !ok {
			six = number
		} else {
			zero = number
		}
	}
	return []map[string]struct{}{zero, one, two, three, four, five, six, seven, eight, nine}
}

func isContained(s1, s2 map[string]struct{}) (bool, int) {
	c := true
	occurences := 0
	for char := range s1 {
		_, ok := s2[char]
		if !ok {
			c = false
		} else {
			occurences += 1
		}
	}
	return c, occurences
}

func computeSignal(pattern, cipher []map[string]struct{}) int {
	totalBuilder := ""
	for _, encodedDigit := range cipher {
		// pattern is ordered ascending from 0 to 9
		for number, decodedDigit := range pattern {
			if ok, occurences := isContained(decodedDigit, encodedDigit); ok && occurences == len(encodedDigit) {
				totalBuilder += fmt.Sprintf("%v", number)
				break
			}
		}
	}
	total, err := strconv.Atoi(totalBuilder)
	if err != nil {
		return 0
	}
	log.Println(total)
	return total
}

func Run() {
	data := loadInput("./input")
	var sum int
	for _, signal := range data {
		decodedNums := extractNumberParts(signal.leftSide)
		output := computeSignal(decodedNums, signal.rightSide)
		sum += output
	}
	log.Println(sum)
}
