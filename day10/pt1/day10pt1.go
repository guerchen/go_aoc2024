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

func merge_sets(set0, set1 map[Coordinates]bool) map[Coordinates]bool {
	new_set := make(map[Coordinates]bool)
	for key := range set0 {
		if _, ok := new_set[key]; !ok {
			new_set[key] = true
		}
	}
	for key := range set1 {
		if _, ok := new_set[key]; !ok {
			new_set[key] = true
		}
	}
	return new_set
}

type Coordinates struct {
	x int
	y int
}

func reacheable_9_postitions(
	topographic_map [][]string,
	coords Coordinates,
	reachable_9s map[Coordinates]bool,
) map[Coordinates]bool {
	if topographic_map[coords.x][coords.y] == "9" {
		position_set := make(map[Coordinates]bool)
		position_set[coords] = true
		return merge_sets(reachable_9s, position_set)
	}

	// UP
	if coords.x > 0 && topographic_map[coords.x-1][coords.y] != "." {
		if str_to_int(topographic_map[coords.x-1][coords.y])-str_to_int(topographic_map[coords.x][coords.y]) == 1 {
			reachable_9s = merge_sets(
				reachable_9s,
				reacheable_9_postitions(topographic_map, Coordinates{coords.x - 1, coords.y}, reachable_9s),
			)
		}
	}
	// DOWN
	if coords.x < len(topographic_map)-1 && topographic_map[coords.x+1][coords.y] != "." {
		if str_to_int(topographic_map[coords.x+1][coords.y])-str_to_int(topographic_map[coords.x][coords.y]) == 1 {
			reachable_9s = merge_sets(
				reachable_9s,
				reacheable_9_postitions(topographic_map, Coordinates{coords.x + 1, coords.y}, reachable_9s),
			)
		}
	}
	// LEFT
	if coords.y > 0 && topographic_map[coords.x][coords.y-1] != "." {
		if str_to_int(topographic_map[coords.x][coords.y-1])-str_to_int(topographic_map[coords.x][coords.y]) == 1 {
			reachable_9s = merge_sets(
				reachable_9s,
				reacheable_9_postitions(topographic_map, Coordinates{coords.x, coords.y - 1}, reachable_9s),
			)
		}
	}
	// RIGHT
	if coords.y < len(topographic_map[0])-1 && topographic_map[coords.x][coords.y+1] != "." {
		if str_to_int(topographic_map[coords.x][coords.y+1])-str_to_int(topographic_map[coords.x][coords.y]) == 1 {
			reachable_9s = merge_sets(
				reachable_9s,
				reacheable_9_postitions(topographic_map, Coordinates{coords.x, coords.y + 1}, reachable_9s),
			)
		}
	}

	return reachable_9s
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
	sum_scores := 0
	for i, row := range topographic_grid {
		for j, elem := range row {
			if elem == "0" {
				reachable_9s := make(map[Coordinates]bool)
				score_map := reacheable_9_postitions(topographic_grid, Coordinates{i, j}, reachable_9s)
				sum_scores += len(score_map)
			}
		}
	}

	fmt.Println(sum_scores)
}
