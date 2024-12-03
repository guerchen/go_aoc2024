package main

import (
	"fmt"
	"os"
	"regexp"
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

	r, _ := regexp.Compile("(mul\\([0-9]{1,3},[0-9]{1,3}\\)|do(n't)?\\(\\))")

	mul_ops := r.FindAllString(string(dat), -1)

	sum_multiplications := 0
	activated := true
	for _, op := range mul_ops {
		if op == "do()" {
			activated = true
		} else if op == "don't()" {
			activated = false
		} else {
			if activated {
				mul_op_split := strings.Split(op, ",")
				first_factor := mul_op_split[0][4:]
				second_factor := mul_op_split[1][:len(mul_op_split[1])-1]

				first_factor_int, err := strconv.Atoi(first_factor)
				check(err)

				second_factor_int, err := strconv.Atoi(second_factor)
				check(err)

				sum_multiplications += first_factor_int * second_factor_int
			}
		}

	}
	fmt.Println(sum_multiplications)

}
