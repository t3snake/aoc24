package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func PrintChecksum() {

	input_b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(input_b)

	unpacked := make([]int, 0)
	filled := make([]int, 0)

	index := 0

	for i := 0; i < len(input); i++ {
		multiplier, _ := strconv.Atoi(string(input[i]))
		if i%2 == 0 {
			for range multiplier {
				unpacked = append(unpacked, index)
				filled = append(filled, index)
			}
			index++
		} else {
			for range multiplier {
				unpacked = append(unpacked, -1)
			}
		}
	}

	dot_count := 0
	result := make([]int, len(filled))
	for idx := range len(filled) {
		if unpacked[idx] == -1 {
			dot_count++
			result[idx] = filled[len(filled)-dot_count]
		} else {
			result[idx] = unpacked[idx]
		}
	}

	checksum := 0

	for idx, val := range result {
		if val == -1 {
			continue
		}
		checksum += idx * val
	}

	fmt.Println(checksum)
}
