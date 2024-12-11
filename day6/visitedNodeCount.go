package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
)

type Vector struct {
	x int
	y int
}

func GetInitialDirection(direction byte) Vector {
	if direction == '^' {
		return Vector{
			x: 0,
			y: -1,
		}
	} else if direction == '>' {
		return Vector{
			x: 1,
			y: 0,
		}
	} else if direction == '<' {
		return Vector{
			x: -1,
			y: 0,
		}
	} else if direction == 'v' {
		return Vector{
			x: 0,
			y: 1,
		}
	} else {
		return Vector{
			x: 0,
			y: 0,
		}
	}
}

func Turn90Degrees(direction Vector) Vector {
	if direction.x == 1 && direction.y == 0 {
		return GetInitialDirection('v')
	} else if direction.x == 0 && direction.y == 1 {
		return GetInitialDirection('<')
	} else if direction.x == -1 && direction.y == 0 {
		return GetInitialDirection('^')
	} else if direction.x == 0 && direction.y == -1 {
		return GetInitialDirection('>')
	} else {
		// (0,0)
		return GetInitialDirection('a')
	}
}

func WithinBounds(grid []string, position Vector) bool {
	if position.x < 0 || position.y < 0 {
		return false
	}

	if position.x >= len(grid[0]) || position.y >= len(grid) {
		return false
	}

	return true
}

func GetNextPosition(position, direction Vector) Vector {
	return Vector{
		x: position.x + direction.x,
		y: position.y + direction.y,
	}
}

func PrintTotalVisitedNodes() {
	visited := make(map[Vector]bool)

	input_b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := strings.Split(string(input_b), "\n")

	var direction Vector
	var position Vector

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid); col++ {
			if grid[row][col] != '.' && grid[row][col] != '#' {
				direction = GetInitialDirection(grid[row][col])
				position = Vector{
					x: col,
					y: row,
				}
			}
		}
	}

	visited[position] = true

	for WithinBounds(grid, position) {
		next_pos := GetNextPosition(position, direction)

		if WithinBounds(grid, next_pos) && grid[next_pos.y][next_pos.x] == '#' {
			direction = Turn90Degrees(direction)
			continue
		} else {
			position = next_pos
			visited[position] = true
		}
	}
	// fmt.Print(visited)
	fmt.Print(len(slices.Collect(maps.Keys(visited)))-1, "\n")
}
