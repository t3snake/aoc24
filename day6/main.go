package main

import (
	"fmt"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("part1")()
	PrintTotalVisitedNodes()
	PrintPossibleCycles()
}
