package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		panic("an error happened when reading the file")
	}
	data := string(file)

	// ======================================================================
	// ======================================================================
	// =============================== PART 1 ===============================
	// ======================================================================
	// ======================================================================

	regex, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	if err != nil {
		panic("wrong regex expression")
	}

	Matchs := regex.FindAllString(data, -1)

	regexNum, err := regexp.Compile(`\d+`)

	result := 0

	for _, match := range Matchs {
		num := regexNum.FindAllString(match, 2)
		first, err := strconv.Atoi(num[0])
		if err != nil {
			panic("an error happened when converting string to int")
		}
		second, err := strconv.Atoi(num[1])
		if err != nil {
			panic("an error happened when converting string to int")
		}

		result += first * second
	}
	fmt.Println(len(Matchs))
	fmt.Printf("Part 1 : %d\n", result)

	// ======================================================================
	// ======================================================================
	// =============================== PART 2 ===============================
	// ======================================================================
	// ======================================================================

	result = 0
	regex, err = regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	Matchs = regex.FindAllString(data, -1)
	skipCheck := false
	for _, match := range Matchs {
		if match == "don't()" {
			skipCheck = true
		} else if match == "do()" {
			skipCheck = false
		} else if !skipCheck {
			num := regexNum.FindAllString(match, 2)
			first, err := strconv.Atoi(num[0])
			if err != nil {
				panic("an error happened when converting string to int")
			}
			second, err := strconv.Atoi(num[1])
			if err != nil {
				panic("an error happened when converting string to int")
			}

			result += first * second
		}
	}
	fmt.Printf("Part 2 : %d\n", result)

}
