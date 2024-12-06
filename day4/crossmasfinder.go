package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func PrintCrossMasMatches() {
	var num_of_match int

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := strings.Split(string(input), "\n")

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[0])-1; col++ {
			if grid[row][col] == 'A' {
				var forw bool
				var back bool

				// \
				if grid[row-1][col-1] == 'M' && grid[row+1][col+1] == 'S' {
					forw = true
				} else if grid[row-1][col-1] == 'S' && grid[row+1][col+1] == 'M' {
					forw = true
				}

				// /
				if grid[row-1][col+1] == 'M' && grid[row+1][col-1] == 'S' {
					back = true
				} else if grid[row-1][col+1] == 'S' && grid[row+1][col-1] == 'M' {
					back = true
				}

				if forw && back {
					num_of_match += 1
				}
			}
		}
	}

	fmt.Print(num_of_match)
}
