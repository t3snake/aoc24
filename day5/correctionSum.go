package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func makeValid(update []int, rulemap map[int][]int) int {
	set := make(map[int]bool)
	for i := 0; i < len(update); i++ {
		afters := rulemap[update[i]]

		for j := 0; j < len(afters); j++ {
			if set[afters[j]] {
				idx := slices.Index(update, afters[j])
				update[idx], update[i] = update[i], update[idx]

				if IsValid(update, rulemap) {
					return update[(len(update)-1)/2]
				} else {
					i = 0
					set = make(map[int]bool)
					continue
				}
			}
		}

		set[update[i]] = true
	}
	return update[(len(update)-1)/2]
}

func PrintWrongUpdateSum() {
	var update_sum int

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(input), "\n\n")

	rules := strings.Split(inputs[0], "\n")

	updates := strings.Split(inputs[1], "\n")

	rulemap := make(map[int][]int)

	for i := 0; i < len(rules); i++ {
		rule := strings.Split(rules[i], "|")
		before, _ := strconv.Atoi(rule[0])
		after, _ := strconv.Atoi(rule[1])
		rulemap[before] = append(rulemap[before], after)
	}

	// fmt.Print(rulemap)

	for i := 0; i < len(updates); i++ {
		pages_str := strings.Split(updates[i], ",")
		pages := make([]int, len(pages_str))

		for j := 0; j < len(pages_str); j++ {
			pages[j], _ = strconv.Atoi(pages_str[j])
		}

		if !IsValid(pages, rulemap) {
			update_sum += makeValid(pages, rulemap)
		}
	}

	fmt.Print(update_sum, "\n")
}
