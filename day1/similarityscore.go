package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PrintSimilarityScore() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.SplitAfter(string(input), "\n")

	var list1 []int
	hashmap := make(map[int]int)

	for i := 0; i < len(lines)-1; i++ {
		nums := strings.Fields(lines[i])
		// fmt.Print(nums)

		num1, _ := strconv.Atoi(nums[0])
		list1 = append(list1, num1)

		num2, _ := strconv.Atoi(nums[1])
		freq, exists := hashmap[num2]
		if exists {
			hashmap[num2] = freq + 1
		} else {
			hashmap[num2] = 1
		}
	}

	sum := 0

	for i := 0; i < len(list1); i++ {
		sum += list1[i] * hashmap[list1[i]]
	}

	fmt.Println(sum)

}
