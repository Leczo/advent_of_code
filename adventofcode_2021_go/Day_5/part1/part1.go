package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinates struct {
	x1 uint16
	x2 uint16
	y1 uint16
	y2 uint16
}

func loadCoordinates(input string) []Coordinates {

	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	coordinatesList := []Coordinates{}

	for scanner.Scan() {
		coordinatesList = append(coordinatesList, parseCoords(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)

	}
	return coordinatesList
}

func parseCoords(coords string) Coordinates {
	coordFields := strings.Fields(coords) // exp. "123,421 -> 552,641"
	rightSide := strings.Split(coordFields[0], ",")
	leftSide := strings.Split(coordFields[2], ",")
	x1 := strToUint16(rightSide[0])
	x2 := strToUint16(leftSide[0])
	y1 := strToUint16(rightSide[1])
	y2 := strToUint16(leftSide[1])

	return Coordinates{x1, x2, y1, y2}
}

func strToUint16(s string) uint16 {
	sInt, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}
	return uint16(sInt)
}

func validateCoords(c []Coordinates) []Coordinates {
	// accepting only horizontal and vertical lines x1 == x2 or y1 == y2
	complientCoords := []Coordinates{}
	for _, coord := range c {
		if coord.x1 == coord.x2 || coord.y2 == coord.y1 {
			complientCoords = append(complientCoords, coord)
		}
	}
	return complientCoords
}

func markOnMap(coords []Coordinates) [999][999]byte {
	area := [999][999]byte{}

	for _, c := range coords {
		xDist := countDistance(c.x1, c.x2)
		yDist := countDistance(c.y1, c.y2)
		xDist, yDist = adjustLength(xDist, yDist)
		for i := 0; i < len(xDist); i++ {
			area[yDist[i]][xDist[i]] += 1
		}
	}
	return area
}

func countDistance(v1, v2 uint16) []uint16 {
	distance := []uint16{}

	if v1 >= v2 {
		for ; v1 >= v2; v1-- {
			distance = append(distance, v1)
		}
	} else {
		for ; v1 <= v2; v1++ {
			distance = append(distance, v1)
		}
	}
	return distance
}

func adjustLength(v1, v2 []uint16) ([]uint16, []uint16) {
	// When lines are horizontal or vertical distance between them is 0
	// therefore extending slice to match longer one
	for len(v1) > len(v2) {
		v2 = append(v2, v2[0])
	}

	for len(v2) > len(v1) {
		v1 = append(v1, v1[0])
	}
	return v1, v2

}

func countOverlap(a [999][999]byte) int {
	var overlap int
	for _, row := range a {
		for _, elem := range row {
			if elem > 1 {
				overlap += 1
			}
		}
	}
	return overlap
}

func Run() {
	data := loadCoordinates("./input")
	properCoordinates := validateCoords(data)
	area := markOnMap(properCoordinates)
	overlapping := countOverlap(area)
	fmt.Println(overlapping)

}
