package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int // rows / lines
	y int // cols
}

var (
	posibilities = []Point{{x: -1, y: -1}, {x: -1, y: 0}, {x: -1, y: 1}, {x: 0, y: -1}, {x: 0, y: 1}, {x: 1, y: -1}, {x: 1, y: 0}, {x: 1, y: 1}}
	letters      = []rune{'X', 'M', 'A', 'S'}
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic("an error happened when opening the file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	data := []string{}
	xIndexes := []Point{}
	for scanner.Scan() {
		line := scanner.Text()
		data = append(data, line)
		x := len(data) - 1
		for y, charac := range line {
			if charac == letters[0] {
				xIndexes = append(xIndexes, Point{
					x: x,
					y: y,
				})
			}
		}
	}

	result := 0
	for _, xs := range xIndexes {
		result += checkPoint(xs, data)
	}

	fmt.Printf("Part 1 : %d\n", result)

	aIndexes := []Point{}
	for x, rows := range data {
		for y, char := range rows {
			if char == 'A' {
				aIndexes = append(aIndexes, Point{x: x, y: y})
			}
		}
	}
	result = 0
	for _, point := range aIndexes {
		exist := 0
		if (point.x-1 < 0) || (point.x+1 >= len(data)) || (point.y-1 < 0) || (point.y+1 >= len(data[0])) {
			continue
		}
		if (data[point.x-1][point.y-1] == 'M' && data[point.x+1][point.y+1] == 'S') || (data[point.x-1][point.y-1] == 'S' && data[point.x+1][point.y+1] == 'M') {
			exist += 1
		}
		if (data[point.x+1][point.y-1] == 'M' && data[point.x-1][point.y+1] == 'S') || (data[point.x+1][point.y-1] == 'S' && data[point.x-1][point.y+1] == 'M') {
			exist += 1
		}
		if exist == 2 {
			result += 1
		}
	}
	fmt.Printf("Part 2 : %d\n", result)

}

func checkPoint(point Point, data []string) int {
	nbmrOfSuccess := 0
	for _, direction := range posibilities {
		if exist := checkPosibility(point, direction, data); exist {
			nbmrOfSuccess += 1
		}
	}
	return nbmrOfSuccess
}

func checkPosibility(point Point, direction Point, data []string) bool {
	for i, rune := range letters {
		if (point.x+(direction.x*i) < 0) || (point.x+(direction.x*i) >= len(data)) || (point.y+(direction.y*i) < 0) || (point.y+(direction.y*i) >= len(data[0])) {
			return false
		}
		if string(rune) != string(data[point.x+direction.x*i][point.y+(direction.y*i)]) {
			return false
		}
	}
	return true
}
