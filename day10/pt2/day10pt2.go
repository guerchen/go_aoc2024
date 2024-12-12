package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func str_to_int(input string) int {
	input_int, err := strconv.Atoi(input)
	check(err)
	return input_int
}

type Coordinates struct {
	x int
	y int
}

func calculate_trailhead_rating(
	topographic_map [][]string,
	coords Coordinates,
) int {
	if topographic_map[coords.x][coords.y] == "9" {
		return 1
	}

	trailhead_rating := 0

	// UP
	if coords.x > 0 && topographic_map[coords.x-1][coords.y] != "." {
		if str_to_int(topographic_map[coords.x-1][coords.y])-str_to_int(topographic_map[coords.x][coords.y]) == 1 {
			trailhead_rating += calculate_trailhead_rating(topographic_map, Coordinates{coords.x - 1, coords.y})
		}
	}
	// DOWN
	if coords.x < len(topographic_map)-1 && topographic_map[coords.x+1][coords.y] != "." {
		if str_to_int(topographic_map[coords.x+1][coords.y])-str_to_int(topographic_map[coords.x][coords.y]) == 1 {
			trailhead_rating += calculate_trailhead_rating(topographic_map, Coordinates{coords.x + 1, coords.y})
		}
	}
	// LEFT
	if coords.y > 0 && topographic_map[coords.x][coords.y-1] != "." {
		if str_to_int(topographic_map[coords.x][coords.y-1])-str_to_int(topographic_map[coords.x][coords.y]) == 1 {
			trailhead_rating += calculate_trailhead_rating(topographic_map, Coordinates{coords.x, coords.y - 1})
		}
	}
	// RIGHT
	if coords.y < len(topographic_map[0])-1 && topographic_map[coords.x][coords.y+1] != "." {
		if str_to_int(topographic_map[coords.x][coords.y+1])-str_to_int(topographic_map[coords.x][coords.y]) == 1 {
			trailhead_rating += calculate_trailhead_rating(topographic_map, Coordinates{coords.x, coords.y + 1})
		}
	}

	return trailhead_rating
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	topographic_map := strings.Split(string(dat), "\n")

	// create topographic grid
	topographic_grid := make([][]string, len(topographic_map))
	for i, row := range topographic_map {
		row_split := strings.Split(row, "")
		topographic_grid[i] = make([]string, len(row))
		copy(topographic_grid[i], row_split)
	}

	// sum trailhead scores
	sum_ratings := 0
	for i, row := range topographic_grid {
		for j, elem := range row {
			if elem == "0" {
				rating := calculate_trailhead_rating(topographic_grid, Coordinates{i, j})
				sum_ratings += rating
			}
		}
	}

	fmt.Println(sum_ratings)
}
