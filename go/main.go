package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func processDelimiter(input string) (string, rune) {
	hasDelim, _ := regexp.MatchString("//.*", input)
	if !hasDelim {
		return input, ','
	}
	splitRegex := regexp.MustCompile("^//|\n")
	split := splitRegex.Split(input, -1)
	delim := split[1]
	return split[2], []rune(delim)[0]
}

func convertStringsToNumbers(strings []string) ([]int, error) {
	var ints = []int{}
	for _, s := range strings {
		n, _ := strconv.Atoi(s)
		if n < 0 {
			return nil, fmt.Errorf("negatives not allowed: %d", n)
		}
		ints = append(ints, n)
	}
	return ints, nil
}

func add(input string) (int, error) {
	if input == "" {
		return 0, nil
	}
	cleanInput, delim := processDelimiter(input)
	splitRegex := regexp.MustCompile(fmt.Sprintf("%c|\n", delim))

	numbersStrings := splitRegex.Split(cleanInput, -1)
	numbers, err := convertStringsToNumbers(numbersStrings)
	if err != nil {
		return 0, err
	}

	var sum = 0
	for _, element := range numbers {
		sum += element
	}

	return sum, nil
}

func main() {
	fmt.Println("Run tests instead")
}
