package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coordinate struct {
	x int
	y int
}

func simulate_guard_movement_loops(
	area_grid [][]string,
	guard_coords Coordinate,
	guard_direction string,
) bool {
	guard_visited := make(map[Coordinate]int)
	for {
		guard_visited[guard_coords]++
		var next_coordinates Coordinate
		if guard_direction == "UP" {
			next_coordinates = Coordinate{guard_coords.x - 1, guard_coords.y}
		} else if guard_direction == "RIGHT" {
			next_coordinates = Coordinate{guard_coords.x, guard_coords.y + 1}
		} else if guard_direction == "DOWN" {
			next_coordinates = Coordinate{guard_coords.x + 1, guard_coords.y}
		} else if guard_direction == "LEFT" {
			next_coordinates = Coordinate{guard_coords.x, guard_coords.y - 1}
		} else {
			panic("Unrecognized direction!")
		}

		if next_coordinates.x < 0 ||
			next_coordinates.x >= len(area_grid) ||
			next_coordinates.y < 0 ||
			next_coordinates.y >= len(area_grid[0]) {
			return false
		}

		if guard_visited[next_coordinates] > 10 {
			return true
		}

		if area_grid[next_coordinates.x][next_coordinates.y] == "#" {
			if guard_direction == "UP" {
				guard_direction = "RIGHT"
			} else if guard_direction == "RIGHT" {
				guard_direction = "DOWN"
			} else if guard_direction == "DOWN" {
				guard_direction = "LEFT"
			} else if guard_direction == "LEFT" {
				guard_direction = "UP"
			} else {
				panic("Unrecognized direction!")
			}
		} else {
			guard_coords = next_coordinates
		}
	}
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	mapped_area := strings.Split(string(dat), "\n")

	var guard_coords Coordinate

	// build grid
	area_grid := make([][]string, len(mapped_area))
	for i, row := range mapped_area {
		area_grid[i] = make([]string, len(mapped_area[0]))
		row_split := strings.Split(row, "")
		for j, pos := range row_split {
			if pos == "^" {
				guard_coords = Coordinate{i, j}
			}
			area_grid[i][j] = pos
		}
	}

	number_of_loops := 0
	bar := progressbar.Default(100)
	for i, row := range area_grid {
		for j := range row {
			bar.Add(1)
			if area_grid[i][j] == "." {
				area_grid[i][j] = "#"
				if simulate_guard_movement_loops(area_grid, guard_coords, "UP") {
					number_of_loops++
				}
				area_grid[i][j] = "."
			}
		}
	}

	fmt.Println(number_of_loops)
}
