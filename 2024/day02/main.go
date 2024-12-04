package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	list := [][]int{}
	for scanner.Scan() {
		data := readRow(scanner.Text())
		list = append(list, data)
	}

	// ======================================================================
	// ======================================================================
	// =============================== PART 1 ===============================
	// ======================================================================
	// ======================================================================

	nmbrOfSafeReport := 0
	for _, curr := range list {
		if isSafe := safeReport(curr); isSafe {
			nmbrOfSafeReport += 1
		}
	}
	fmt.Printf("Part 1 : %d\n", nmbrOfSafeReport)

	// ======================================================================
	// ======================================================================
	// =============================== PART 1 ===============================
	// ======================================================================
	// ======================================================================

	nmbrOfSafeReport = 0
	for i := range list {
		if isSafe := safeReport(list[i]); isSafe {
			nmbrOfSafeReport += 1
		} else {
			for j := range list[i] {
				newSlice := remove(list[i], j)
				if isSafe := safeReport(newSlice); isSafe {
					nmbrOfSafeReport += 1
					break
				}
			}
		}
	}
	fmt.Printf("Part 2 : %d\n", nmbrOfSafeReport)

}

func readRow(row string) []int {
	data := []int{}
	splitted := strings.Split(row, " ")
	for _, nmbr := range splitted {
		i, err := strconv.Atoi(nmbr)
		if err != nil {
			panic("can't converte a row to intiger")
		}
		data = append(data, i)
	}
	return data
}

const (
	VARIANT_INCREASING  = "increasing"
	VARIANT_DESCREASING = "descreasing"

	MIN_DIF = 1
	MAX_DIF = 3
)

func safeReport(report []int) bool {
	var variant string
	if report[1]-report[0] > 0 {
		variant = VARIANT_INCREASING
	} else if report[1]-report[0] < 0 {
		variant = VARIANT_DESCREASING
	} else {
		return false
	}
	for i := 0; i < (len(report) - 1); i++ {
		result := report[i+1] - report[i]
		if variant == VARIANT_INCREASING && result >= MIN_DIF && result <= MAX_DIF {
			continue
		} else if variant == VARIANT_DESCREASING && result < 0 && math.Abs(float64(result)) >= MIN_DIF && math.Abs(float64(result)) <= MAX_DIF {
			continue
		} else {
			return false
		}
	}
	return true
}
func remove(slice []int, s int) []int {
	destination := make([]int, len(slice))
	copy(destination, slice)
	return append(destination[:s], destination[s+1:]...)
}
