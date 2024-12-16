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

type BlinkInput struct {
	stone_value int
	num_blinks  int
}

var memo = make(map[BlinkInput]int)

func blink_stone(stone_value, num_blinks int) int {
	num_stones := 0

	if num_blinks == 0 {
		return 1
	} else {
		if val, ok := memo[BlinkInput{stone_value, num_blinks}]; ok {
			return val
		}
		if stone_value == 0 {
			num_stones += blink_stone(1, num_blinks-1)
		} else if len(fmt.Sprint(stone_value))%2 == 0 {
			value_str := fmt.Sprintln(stone_value)

			first_half := strings.Replace(value_str[:len(value_str)/2], "\n", "", -1)
			second_half := strings.Replace(value_str[len(value_str)/2:], "\n", "", -1)
			first_half_int, err := strconv.Atoi(first_half)
			check(err)
			second_half_int, err := strconv.Atoi(second_half)
			check(err)

			num_stones += blink_stone(first_half_int, num_blinks-1)
			num_stones += blink_stone(second_half_int, num_blinks-1)
		} else {
			num_stones += blink_stone(stone_value*2024, num_blinks-1)
		}
	}

	memo[BlinkInput{stone_value, num_blinks}] = num_stones
	return num_stones
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

	final_stone_count := 0
	num_blinks := 75
	for _, stone_value := range stone_queue {
		final_stone_count += blink_stone(stone_value, num_blinks)
	}

	fmt.Println(final_stone_count)

}
