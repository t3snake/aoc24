package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// after mul( is there num1, num2 and )
func GetProductIfValidMul(input string) int {
	index_comma := strings.Index(input, ",")
	index_bracket := strings.Index(input, ")")

	if index_bracket == -1 || index_comma == -1 || index_bracket < index_comma {
		return 0
	}

	// fmt.Print(cur_str[:index_comma], " ", cur_str[index_comma+1:index_bracket], "\n")

	num1, err := strconv.Atoi(input[:index_comma])
	if err != nil {
		return 0
	}

	num2, err := strconv.Atoi(input[index_comma+1 : index_bracket])
	if err != nil {
		return 0
	}

	return num1 * num2
}

func PrintSumOfValidMultiplication() {
	var sum int

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	mulOccurence := strings.SplitAfter(string(input), "mul(")

	for i := 1; i < len(mulOccurence); i++ {
		sum += GetProductIfValidMul(mulOccurence[i])
	}

	fmt.Print(sum, "\n")
}
