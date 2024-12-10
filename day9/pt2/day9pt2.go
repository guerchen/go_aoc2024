package main

import (
	"fmt"
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
	disk_map := strings.Split(string(dat), "")

	// create sparse representation
	sparse_format_slice := make([][]string, 0)
	for i, val := range disk_map {
		val_int, err := strconv.Atoi(val)
		check(err)
		if val_int > 0 {
			var id string
			if i%2 == 0 {
				id = fmt.Sprint(i / 2)
			} else {
				id = "."
			}
			file := make([]string, val_int)
			for j := 0; j < val_int; j++ {
				file[j] = id
			}
			sparse_format_slice = append(sparse_format_slice, file)
		}
	}

	// creating compact representation
	for j := 0; j < len(sparse_format_slice); j++ {
		for i := len(sparse_format_slice) - 1; i > j; i-- {
			if sparse_format_slice[i][0] != "." &&
				sparse_format_slice[j][0] == "." &&
				len(sparse_format_slice[i]) <= len(sparse_format_slice[j]) {
				size_diff := len(sparse_format_slice[j]) - len(sparse_format_slice[i])
				sparse_format_slice[i], sparse_format_slice[j] = sparse_format_slice[j], sparse_format_slice[i]
				if size_diff > 0 {
					new_blank := make([]string, size_diff)
					for k := 0; k < size_diff; k++ {
						new_blank[k] = "."
					}
					sparse_format_slice = slices.Insert(sparse_format_slice, j+1, new_blank)
					sparse_format_slice[i+1] = sparse_format_slice[i+1][:len(sparse_format_slice[j])]
				}

			}
		}
	}

	// calculate checksum
	checksum := 0
	flat_slice := make([]string, 0)
	for _, file := range sparse_format_slice {
		flat_slice = append(flat_slice, file...)
	}

	for pos, id := range flat_slice {
		if id != "." {
			id_int, err := strconv.Atoi(id)
			check(err)
			checksum += pos * id_int
		}

	}

	fmt.Println(checksum)
}
