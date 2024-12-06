package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func PrintXmasMatches() {
	var num_of_match int

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := strings.Split(string(input), "\n")

	// -
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0])-3; col++ {
			horiz := string(([]byte{grid[row][col], grid[row][col+1], grid[row][col+2], grid[row][col+3]})[:])
			if horiz == "XMAS" || horiz == "SAMX" {
				num_of_match += 1
			}
		}
	}

	// |
	for row := 0; row < len(grid)-3; row++ {
		for col := 0; col < len(grid[0]); col++ {
			vert := string(([]byte{grid[row][col], grid[row+1][col], grid[row+2][col], grid[row+3][col]})[:])
			if vert == "XMAS" || vert == "SAMX" {
				num_of_match += 1
			}
		}
	}

	// \
	for row := 0; row < len(grid)-3; row++ {
		for col := 0; col < len(grid[0])-3; col++ {
			diag := string(([]byte{grid[row][col], grid[row+1][col+1], grid[row+2][col+2], grid[row+3][col+3]})[:])
			if diag == "XMAS" || diag == "SAMX" {
				num_of_match += 1
			}
		}
	}
	// /
	for row := 0; row < len(grid)-3; row++ {
		for col := 3; col < len(grid[0]); col++ {
			diag := string(([]byte{grid[row][col], grid[row+1][col-1], grid[row+2][col-2], grid[row+3][col-3]})[:])
			if diag == "XMAS" || diag == "SAMX" {
				num_of_match += 1
			}
		}
	}

	fmt.Print(num_of_match, "\n")
}
