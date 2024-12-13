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

func update_stone_queue(stone_queue []int) []int {
	new_stone_queue := make([]int, 0)
	for i := range stone_queue {
		if stone_queue[i] == 0 {
			new_stone_queue = append(new_stone_queue, 1)
		} else if len(fmt.Sprint(stone_queue[i]))%2 == 0 {
			value_str := fmt.Sprintln(stone_queue[i])

			first_half := strings.Replace(value_str[:len(value_str)/2], "\n", "", -1)
			second_half := strings.Replace(value_str[len(value_str)/2:], "\n", "", -1)
			first_half_int, err := strconv.Atoi(first_half)
			check(err)
			second_half_int, err := strconv.Atoi(second_half)
			check(err)

			new_stone_queue = append(new_stone_queue, first_half_int)
			new_stone_queue = append(new_stone_queue, second_half_int)
		} else {
			new_stone_queue = append(new_stone_queue, stone_queue[i]*2024)
		}
	}

	return new_stone_queue
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	stone_queue_str := strings.Split(string(dat), " ")

	stone_queue := make([]int, len(stone_queue_str))
	for i, stone := range stone_queue_str {
		stone_value, err := strconv.Atoi(stone)
		check(err)
		stone_queue[i] = stone_value
	}

	for count_blinks := 25; count_blinks > 0; count_blinks-- {
		stone_queue = update_stone_queue(stone_queue)
	}

	fmt.Println(len(stone_queue))

}
