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

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	location_id_rows := strings.Split(string(dat), "\n")

	left_slice := []int{}
	right_slice := []int{}

	for _, row := range location_id_rows {
		row_split := strings.Split(row, "   ")
		left_value, err := strconv.Atoi(row_split[0])
		check(err)
		left_slice = append(left_slice, left_value)
		right_value, err := strconv.Atoi(row_split[1])
		check(err)
		right_slice = append(right_slice, right_value)
	}

	left_splice_count := make(map[int]int)
	for _, left_id := range left_slice {
		left_splice_count[left_id] = 0
	}

	for _, right_id := range right_slice {
		if _, ok := left_splice_count[right_id]; ok {
			left_splice_count[right_id] += 1
		}
	}

	similarity_score := 0
	for id, count := range left_splice_count {
		similarity_score += id * count
	}

	fmt.Println(similarity_score)

}
