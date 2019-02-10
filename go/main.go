package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func convertStringsToNumbers(strings []string) []int {
	var ints = []int{}
	for _, s := range strings {
		n, _ := strconv.Atoi(s)
		ints = append(ints, n)
	}
	return ints
}

func add(input string) int {
	if input == "" {
		return 0
	}

	splitRegex := regexp.MustCompile(",|\n")

	numbersStrings := splitRegex.Split(input, -1)
	numbers := convertStringsToNumbers(numbersStrings)

	var sum = 0
	for _, element := range numbers {
		sum += element
	}

	return sum
}

func main() {
	fmt.Println("Run tests instead")
}
