package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func isIncreasing(levels []int) bool {
	levels_ascending := slices.Clone(levels)
	sort.Ints(levels_ascending)

	return slices.Equal(levels, levels_ascending)
}

func isValidIncr(levels []int) bool {
	for j := 1; j < len(levels); j++ {
		if levels[j]-levels[j-1] < 1 || levels[j]-levels[j-1] > 3 {
			return false
		}
	}
	return true
}

func isDecreasing(levels []int) bool {
	levels_descending := slices.Clone(levels)
	sort.Ints(levels_descending)
	slices.Reverse(levels_descending)

	return slices.Equal(levels, levels_descending)
}

func isValidDecr(levels []int) bool {
	for j := 1; j < len(levels); j++ {
		if levels[j-1]-levels[j] < 1 || levels[j-1]-levels[j] > 3 {
			return false
		}
	}
	return true
}

func isValidLevel(levels []int) bool {
	if isDecreasing(levels) && isValidDecr(levels) {
		return true
	}

	if isIncreasing(levels) && isValidIncr(levels) {
		return true
	}

	return false
}

func levelWithoutIndex(slice []int, index int) []int {
	newSlice := make([]int, 0)
	newSlice = append(newSlice, slice[:index]...)
	if index != len(slice)-1 {
		newSlice = append(newSlice, slice[index+1:]...)
	}
	return newSlice
}

func isValidWithProblemDampener(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		truncLevel := levelWithoutIndex(levels, i)
		if isValidLevel(truncLevel) {
			return true
		}
	}
	return false
}

func PrintSafeReportsWithProblemDampener() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reports := strings.SplitAfter(string(input), "\n")

	sum := 0

	for i := 0; i < len(reports); i++ {

		levels_str := strings.Fields(reports[i])
		levels := make([]int, len(levels_str))
		// fmt.Print(levels)

		for j := 0; j < len(levels_str); j++ {
			num, _ := strconv.Atoi(levels_str[j])
			levels[j] = num
		}

		if isValidLevel(levels) {
			sum += 1
		} else if isValidWithProblemDampener(levels) {
			sum += 1
		}
	}

	fmt.Print(sum)

}
