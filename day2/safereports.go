package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PrintSafeReports() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reports := strings.SplitAfter(string(input), "\n")

	sum := 0

	for i := 0; i < len(reports); i++ {
		var incr bool
		var decr bool
		var fail bool

		levels := strings.Fields(reports[i])
		// fmt.Print(levels)

		for j := 1; j < len(levels); j++ {
			prev, _ := strconv.Atoi(levels[j-1])
			cur, _ := strconv.Atoi(levels[j])

			if prev < cur {
				incr = true
				if cur-prev < 1 || cur-prev > 3 {
					fail = true
					break
				}

			} else if cur < prev {
				decr = true
				if prev-cur < 1 || prev-cur > 3 {
					fail = true
					break
				}
			} else {
				fail = true
				break
			}
		}

		if !fail && !(incr && decr) {
			sum += 1
		}
	}

	fmt.Print(sum, "\n")

}
