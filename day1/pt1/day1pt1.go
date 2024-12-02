package main

import (
	"fmt"
	"math"
	"os"
	"slices"
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

	slices.Sort(left_slice)
	slices.Sort(right_slice)

	diffs_sum := 0
	for i := range location_id_rows {
		diffs_sum += int(math.Abs(float64(left_slice[i] - right_slice[i])))
	}

	fmt.Println(diffs_sum)

}
