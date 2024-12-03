package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func PrintDiffSum() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.SplitAfter(string(input), "\n")

	var list1 []int
	var list2 []int

	for i := 0; i < len(lines)-1; i++ {
		nums := strings.Fields(lines[i])
		// fmt.Print(nums)

		num1, _ := strconv.Atoi(nums[0])
		list1 = append(list1, num1)

		num2, _ := strconv.Atoi(nums[1])
		list2 = append(list2, num2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0

	for i := 0; i < len(list1); i++ {
		if list1[i] > list2[i] {
			sum += list1[i] - list2[i]
		} else {
			sum += list2[i] - list1[i]
		}
	}

	fmt.Println(sum)

}
