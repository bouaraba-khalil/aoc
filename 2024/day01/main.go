package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic("an error happened when reading the file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	firstList := []int{}
	secondList := []int{}

	for scanner.Scan() {
		data := readRow(scanner.Text())
		firstList = append(firstList, data[0])
		secondList = append(secondList, data[1])
	}

	// ======================================================================
	// ======================================================================
	// =============================== PART 1 ===============================
	// ======================================================================
	// ======================================================================

	sort.Ints(firstList)
	sort.Ints(secondList)

	result := 0
	for i := 0; i < len(firstList); i++ {
		diff := int(math.Abs(float64(firstList[i]) - float64(secondList[i])))
		result += diff
	}

	fmt.Printf("Part 1 : %d\n", result)

	// ======================================================================
	// ======================================================================
	// =============================== PART 2 ===============================
	// ======================================================================
	// ======================================================================

	result = calculateDistance(firstList, secondList)

	fmt.Printf("Part 2 : %d\n", result)

}

func readRow(row string) [2]int {
	data := [2]int{}
	splitted := strings.Split(row, "   ")
	for index, nmbr := range splitted {
		i, err := strconv.Atoi(nmbr)
		if err != nil {
			panic("can't converte a row to intiger")
		}
		data[index] = i
	}
	return data
}

func calculateDistance(firstList, secondList []int) int {
	alreadyCalculated := map[int]int{}
	result := 0
	for i := 0; i < len(firstList); i++ {
		if appearance, err := alreadyCalculated[firstList[i]]; err {
			result += firstList[i] * appearance
			continue
		}
		appearance := 0
		for j := 0; j < len(secondList); j++ {
			if firstList[i] == secondList[j] {
				appearance += 1
			}
		}
		alreadyCalculated[firstList[i]] = appearance
		result += firstList[i] * appearance
	}
	return result
}
