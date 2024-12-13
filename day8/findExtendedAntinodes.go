package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func findExtendedNodes(index1, index2 []int, max_i, max_j int) [][]int {
	result := make([][]int, 0)
	for i := 0; ; i++ {
		new_index := []int{index1[0] - i*(index2[0]-index1[0]), index1[1] - i*(index2[1]-index1[1])}

		if isValidIndex(new_index[0], new_index[1], max_i, max_j) {
			result = append(result, new_index)
		} else {
			break
		}
	}
	return result
}

func PrintExtendedAntinodes() {
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

				result_left := findExtendedNodes(ref_index, cmp_index, len(grid), len(grid[0]))

				for _, value := range result_left {
					antinode_index_map[fmt.Sprint(value[0], " ", value[1])] = true
				}

				result_right := findExtendedNodes(cmp_index, ref_index, len(grid), len(grid[0]))

				for _, value := range result_right {
					antinode_index_map[fmt.Sprint(value[0], " ", value[1])] = true
				}

			}
		}
	}

	fmt.Println(len(antinode_index_map))
}
