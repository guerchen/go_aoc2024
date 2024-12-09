package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coordinates struct {
	x int
	y int
}

func vector_modulus(vector Coordinates) float64 {
	return math.Sqrt(float64(vector.x*vector.x) + float64(vector.y*vector.y))
}

func is_collinear(first_vector, second_vector Coordinates) bool {
	dot_product := float64(first_vector.x*second_vector.x+first_vector.y*second_vector.y) /
		(vector_modulus(first_vector) * vector_modulus(second_vector))

	if math.Abs(dot_product) > 0.999999999 {
		return true
	} else {
		return false
	}

}

func is_distance_ratio_two(first_vector, second_vector Coordinates) bool {
	first_vector_modulus := vector_modulus(first_vector)
	second_vector_modulus := vector_modulus(second_vector)

	ratio := math.Abs(first_vector_modulus / second_vector_modulus)

	if ratio == 2 || ratio == 0.5 {
		return true
	} else {
		return false
	}
}

func is_antinode(
	point_coordenates Coordinates,
	antenna_positions map[string][]Coordinates,
) bool {
	for _, coords := range antenna_positions {
		for k := range coords {
			for l := range coords {
				if k < l {
					first_vector := Coordinates{
						coords[k].x - point_coordenates.x,
						coords[k].y - point_coordenates.y,
					}
					second_vector := Coordinates{
						coords[l].x - point_coordenates.x,
						coords[l].y - point_coordenates.y,
					}

					if is_collinear(first_vector, second_vector) &&
						is_distance_ratio_two(first_vector, second_vector) {
						return true
					}
				}
			}
		}

	}

	return false
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	map_rows := strings.Split(string(dat), "\n")

	// create map matrix and antenna positions
	map_grid := make([][]string, len(map_rows))
	antenna_positions := make(map[string][]Coordinates)
	for i, row := range map_rows {
		map_grid[i] = make([]string, len(map_rows[0]))
		row_split := strings.Split(row, "")
		for j, pos := range row_split {
			map_grid[i][j] = pos
			if pos != "." {
				antenna_positions[pos] = append(antenna_positions[pos], Coordinates{i, j})
			}
		}
	}

	count_antinodes := 0
	// find antinodes
	for i, row := range map_grid {
		for j := range row {
			if is_antinode(Coordinates{i, j}, antenna_positions) {
				count_antinodes++
			}
		}
	}

	fmt.Println(count_antinodes)
}
