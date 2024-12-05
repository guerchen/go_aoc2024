package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type GridIndex []int

func checkWordMatch(grid [][]string, fst, scd, thr, fth GridIndex, len_x, len_y int) bool {
	if fst[0] < 0 || scd[0] < 0 || thr[0] < 0 || fth[0] < 0 {
		return false
	}
	if fst[1] < 0 || scd[1] < 0 || thr[1] < 0 || fth[1] < 0 {
		return false
	}
	if fst[0] >= len_y || scd[0] >= len_y || thr[0] >= len_y || fth[0] >= len_y {
		return false
	}
	if fst[1] >= len_x || scd[1] >= len_x || thr[1] >= len_x || fth[1] >= len_x {
		return false
	}

	curr_string := strings.Join([]string{
		grid[fst[0]][fst[1]],
		grid[scd[0]][scd[1]],
		grid[thr[0]][thr[1]],
		grid[fth[0]][fth[1]],
	}, "")

	if curr_string == "XMAS" || curr_string == "SAMX" {
		return true
	} else {
		return false
	}
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	letter_rows := strings.Split(string(dat), "\n")

	letter_grid := make([][]string, len(letter_rows))
	for i := range letter_rows {
		letter_grid[i] = make([]string, len(letter_rows[0]))
	}

	for i, row := range letter_rows {
		row_split := strings.Split(row, "")
		copy(letter_grid[i], row_split)
	}

	count_word := 0
	// rows
	for i := range letter_grid {
		for j := range letter_grid[0] {
			if checkWordMatch(
				letter_grid,
				[]int{i, j},
				[]int{i, j + 1},
				[]int{i, j + 2},
				[]int{i, j + 3},
				len(letter_grid[0]),
				len(letter_grid)) {
				count_word++
			}
		}
	}

	//cols
	for i := range letter_grid {
		for j := range letter_grid[0] {
			if checkWordMatch(
				letter_grid,
				[]int{i, j},
				[]int{i + 1, j},
				[]int{i + 2, j},
				[]int{i + 3, j},
				len(letter_grid[0]),
				len(letter_grid)) {
				count_word++
			}
		}
	}

	//right diagonal
	for i := range letter_grid {
		for j := range letter_grid[0] {
			if checkWordMatch(
				letter_grid,
				[]int{i, j},
				[]int{i + 1, j + 1},
				[]int{i + 2, j + 2},
				[]int{i + 3, j + 3},
				len(letter_grid[0]),
				len(letter_grid)) {
				count_word++
			}
		}
	}

	//left diagonal
	for i := range letter_grid {
		for j := range letter_grid[0] {
			if checkWordMatch(
				letter_grid,
				[]int{i, j},
				[]int{i + 1, j - 1},
				[]int{i + 2, j - 2},
				[]int{i + 3, j - 3},
				len(letter_grid[0]),
				len(letter_grid)) {
				count_word++
			}
		}
	}

	fmt.Println(count_word)

}
