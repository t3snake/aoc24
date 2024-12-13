package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func isValidIndex(i, j, max_i, max_j int) bool {
	if i < 0 || j < 0 {
		return false
	}

	if i >= max_i || j >= max_j {
		return false
	}

	return true

}

func PrintUniqueAntinodes() {
	node_map := make(map[byte][][]int)

	input_b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := strings.Split(string(input_b), "\n")

	// map all node indexes
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != '.' {
				val, exists := node_map[grid[i][j]]
				if exists {
					node_map[grid[i][j]] = append(val, []int{i, j})
				} else {
					node_map[grid[i][j]] = [][]int{{i, j}}
				}
			}
		}
	}

	antinode_index_map := make(map[string]bool)

	for _, indexes := range node_map {
		for i := 0; i < len(indexes); i++ {
			ref_index := indexes[i]
			for j := i + 1; j < len(indexes); j++ {
				cmp_index := indexes[j]

				index_a := []int{2*ref_index[0] - cmp_index[0], 2*ref_index[1] - cmp_index[1]}

				if isValidIndex(index_a[0], index_a[1], len(grid), len(grid[0])) {
					antinode_index_map[fmt.Sprint(index_a[0], " ", index_a[1])] = true
				}

				index_b := []int{2*cmp_index[0] - ref_index[0], 2*cmp_index[1] - ref_index[1]}

				if isValidIndex(index_b[0], index_b[1], len(grid), len(grid[0])) {
					antinode_index_map[fmt.Sprint(index_b[0], " ", index_b[1])] = true
				}
			}
		}
	}

	fmt.Println(len(antinode_index_map))
}
