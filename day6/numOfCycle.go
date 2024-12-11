package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func uniqueKey(position, direction Vector) string {
	return fmt.Sprintf("%d %d %d %d", position.x, position.y, direction.x, direction.y)
}

func simulateRun(init_pos, init_direc, wall_pos Vector, grid []string) bool {

	if !WithinBounds(grid, wall_pos) {
		return false
	}

	visited := make(map[string]bool) // key: position (x,y) + direction (x,y)

	position := Vector{
		x: init_pos.x,
		y: init_pos.y,
	}
	direction := Vector{
		x: init_direc.x,
		y: init_direc.y,
	}

	for WithinBounds(grid, position) {

		next_pos := GetNextPosition(position, direction)

		if WithinBounds(grid, next_pos) &&
			(grid[next_pos.y][next_pos.x] == '#' ||
				(next_pos.x == wall_pos.x && next_pos.y == wall_pos.y)) {
			direction = Turn90Degrees(direction)
			continue
		} else {
			if visited[uniqueKey(position, direction)] {
				return true
			}
			visited[uniqueKey(position, direction)] = true
			position = next_pos
		}
	}

	return false

}

func PrintPossibleCycles() {
	var sum int
	solved := make(map[Vector]bool)

	input_b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	grid := strings.Split(string(input_b), "\n")

	var init_direction Vector
	var init_position Vector

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid); col++ {
			if grid[row][col] != '.' && grid[row][col] != '#' {
				init_direction = GetInitialDirection(grid[row][col])
				init_position = Vector{
					x: col,
					y: row,
				}
			}
		}
	}

	direction := Vector{
		x: init_direction.x,
		y: init_direction.y,
	}
	position := Vector{
		x: init_position.x,
		y: init_position.y,
	}

	for WithinBounds(grid, position) {
		next_pos := GetNextPosition(position, direction)

		if WithinBounds(grid, next_pos) && grid[next_pos.y][next_pos.x] == '#' {
			direction = Turn90Degrees(direction)
			continue
		} else {
			position = next_pos
			if !solved[position] && simulateRun(init_position, init_direction, position, grid) {
				solved[position] = true
				sum += 1
			}
		}
	}
	// fmt.Print(visited)
	fmt.Print(sum, "\n")
}
