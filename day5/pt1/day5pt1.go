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
	rules_updates := strings.Split(string(dat), "\n\n")

	//build rulebook
	rules := strings.Split(rules_updates[0], "\n")
	rule_book := make(map[string]map[string]bool)
	for _, rule := range rules {
		key_value := strings.Split(rule, "|")
		if _, ok := rule_book[key_value[0]]; !ok {
			rule_book[key_value[0]] = make(map[string]bool)
		}
		rule_book[key_value[0]][key_value[1]] = true
	}

	//check updates
	mid_term_sum := 0
	updates := strings.Split(rules_updates[1], "\n")
	for _, update := range updates {
		update_slice := strings.Split(update, ",")
		valid_slice := true
		mid_term_index := (len(update_slice) - 1) / 2

		for i, key := range update_slice {
			for _, value := range update_slice[:i] {
				if _, ok := rule_book[key][value]; ok {
					valid_slice = false
					break
				}
			}
			if !valid_slice {
				break
			}
		}

		if valid_slice {
			mid_term_value, err := strconv.Atoi(update_slice[mid_term_index])
			check(err)
			mid_term_sum += mid_term_value
		}
	}

	fmt.Println(mid_term_sum)
}
