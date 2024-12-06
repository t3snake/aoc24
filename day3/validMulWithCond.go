package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func PrintValidMultiplicationWithCond() {
	var sum int

	input_b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := "do()" + string(input_b)

	regex := regexp.MustCompile(regexp.QuoteMeta("do()") + "|" + regexp.QuoteMeta("don't()"))
	doOccurence := regex.FindAllStringIndex(input, -1)

	for i := 0; i < len(doOccurence); i++ {
		cond_index := doOccurence[i]

		if input[cond_index[0]:cond_index[1]] == "do()" {
			start := cond_index[1]
			var end int
			if i == len(doOccurence)-1 {
				end = len(input)
			} else {
				end = doOccurence[i+1][0]
			}

			mulOccurence := strings.SplitAfter(input[start:end], "mul(")

			for j := 1; j < len(mulOccurence); j++ {
				sum += GetProductIfValidMul(mulOccurence[j])
			}
		}

		// sum += num1 * num2
	}

	fmt.Print(sum)
}
