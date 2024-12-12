package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isSumPossible(nums []int, target, current, index int) bool {
	if index == 0 {
		return isSumPossible(nums, target, nums[0], 1)
	}

	if index == len(nums)-1 {
		return (current+nums[index] == target) || (current*nums[index] == target)
	}

	return isSumPossible(nums, target, current+nums[index], index+1) ||
		isSumPossible(nums, target, current*nums[index], index+1)
}

func PrintSumValidResult() {
	var valid_sum int

	input_b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(input_b), "\n")

	for i := 0; i < len(inputs); i++ {
		input := strings.Split(inputs[i], ": ")

		target, _ := strconv.Atoi(input[0])

		nums_str := strings.Split(input[1], " ")

		nums := make([]int, len(nums_str))

		for index, num_str := range nums_str {
			num, _ := strconv.Atoi(num_str)
			nums[index] = num
		}

		if isSumPossible(nums, target, 0, 0) {
			valid_sum += target
		}

	}

	fmt.Println(valid_sum)
}
