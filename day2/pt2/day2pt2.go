package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func check_report_safety(report []int) bool {
	report_safe := true

	for i := range report {
		if i > 0 && (math.Abs(float64(report[i]-report[i-1])) > 3 || report[i]-report[i-1] == 0) {
			report_safe = false
			break
		}
		if i > 1 && (report[i]-report[i-1])*(report[i-1]-report[i-2]) < 0 {
			report_safe = false
			break
		}
	}

	return report_safe
}

func main() {
	dat, err := os.ReadFile("../input.txt")
	check(err)
	reports := strings.Split(string(dat), "\n")

	reports_treated := make([][]int, len(reports))
	for r := range reports {
		reports_treated[r] = []int{}
	}

	for r := range reports {
		levels := strings.Split(reports[r], " ")
		for l := range levels {
			value, err := strconv.Atoi(levels[l])
			check(err)
			reports_treated[r] = append(reports_treated[r], value)
		}
	}

	safe_reports_count := 0
	for _, report := range reports_treated {
		report_safe := check_report_safety(report)

		if !report_safe {
			for i := range report {
				report_skip := make([]int, len(report))
				copy(report_skip, report)

				report_skip = append(report_skip[:i], report_skip[i+1:]...)
				if check_report_safety(report_skip) {
					report_safe = true
					break
				}
			}
		}

		if report_safe {
			safe_reports_count++
		}
	}

	fmt.Println(safe_reports_count)

}
