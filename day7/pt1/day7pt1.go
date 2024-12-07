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

func check_operation_valid(
	expected_result int, operand_slice []int, current_value int,
) bool {

	if len(operand_slice) == 0 {
		return current_value == expected_result
	}

	// addition
	if check_operation_valid(
		expected_result, operand_slice[1:], current_value+operand_slice[0],
	) {
		return true
	}

	// multiplication
	if check_operation_valid(
		expected_result, operand_slice[1:], current_value*operand_slice[0],
	) {
		return true
	}

	return false

}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	test_cases := strings.Split(string(dat), "\n")

	total_calibration_result := 0
	for _, test_case := range test_cases {
		result_operands := strings.Split(test_case, ": ")
		result, err := strconv.Atoi(result_operands[0])
		check(err)
		operand_slice := strings.Split(result_operands[1], " ")
		operand_slice_int := make([]int, len(operand_slice))

		for i, operand := range operand_slice {
			operand_int, err := strconv.Atoi(operand)
			check(err)
			operand_slice_int[i] = operand_int
		}

		if check_operation_valid(result, operand_slice_int[1:], operand_slice_int[0]) {
			total_calibration_result += result
		}
	}

	fmt.Println(total_calibration_result)
}
